package goemu6502

import (
	"fmt"
	"sync"
)

type (
	Registers struct {
		a, x, y, p, sp uint8
		pc             uint16
	}

	InternalRegisters struct {
		fetched       uint8  // Fetched value
		temp          uint16 // Temporary value
		addr_absolute uint16 // All used memory addresses end up here
		addr_relative uint16 // Absolute address following a branch instruction
		addr_mode     AddressingMode
		opcode        uint8 // Current opcode
	}

	InternalStatus struct {
		Cycles                   uint8
		currentInstruction       InstructionInfo
		currentInstructionString string
	}

	CPU struct {
		mutex  sync.Mutex
		r      Registers
		i      InternalRegisters
		status InternalStatus
		bus    Bus
	}

	StatusFlag uint8
)

const (
	_                StatusFlag = iota
	Carry                       = 1 << 0
	Zero                        = 1 << 1
	InterruptDisable            = 1 << 2
	Decimal                     = 1 << 3
	Break                       = 1 << 4
	Unused                      = 1 << 5
	Overflow                    = 1 << 6
	Negative                    = 1 << 7
)

func NewCPU(bus Bus) *CPU {
	return &CPU{
		r:   Registers{},
		i:   InternalRegisters{},
		bus: bus,
	}
}

func (c *CPU) setFlag(flag StatusFlag, value bool) {
	if value {
		c.r.p |= uint8(flag)
	} else {
		c.r.p &= ^uint8(flag)
	}
}

func (c *CPU) getFlag(flag StatusFlag) bool {
	return (c.r.p & uint8(flag)) != 0
}

func (c *CPU) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.r.a = 0x00
	c.r.x = 0x00
	c.r.y = 0x00
	// P == 0x00 | U | I
	c.r.p = 0x00 | uint8(Unused) | uint8(InterruptDisable)
	// PC == read from 0xFFFC and 0xFFFD
	c.r.pc = uint16(c.bus.Read(0xFFFC)) | uint16(c.bus.Read(0xFFFD))<<8
}

func (c *CPU) interrupt() {
	// Push the program counter to the stack
	c.pushWord(c.r.pc)

	// Set the break flag to 0
	c.setFlag(Break, false)

	// Push the processor status to the stack
	c.setFlag(Unused, true)
	c.setFlag(Break, true)
	c.setFlag(InterruptDisable, true)
	c.pushByte(c.r.p)
	c.setFlag(InterruptDisable, false)

	// Set the program counter to the interrupt vector
	c.r.pc = uint16(c.bus.Read(0xFFFE)) | uint16(c.bus.Read(0xFFFF))<<8

	// Set cycles to 7
	c.status.Cycles = 7
}

func Irq(c *CPU) {
	// If the InterruptDisable flag is not set, push the pc and p to the stack
	if !c.getFlag(InterruptDisable) {
		c.interrupt()
	}
}

func Nmi(c *CPU) {
	// This is a non-maskable interrupt
	c.interrupt()
}

func (c *CPU) Complete() bool {
	return c.status.Cycles == 0
}

func (c *CPU) Tick() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.status.Cycles == 0 {
		c.status.currentInstructionString = c.DisassembleAt(c.r.pc)

		// Fetch the next instruction
		c.status.currentInstruction = Instructions[c.bus.Read(c.r.pc)]
		c.i.opcode = c.status.currentInstruction.Opcode
		c.r.pc++

		// Get the number of cycles for the instruction
		c.status.Cycles = c.status.currentInstruction.Cycles

		// Get the addressing mode
		c.i.addr_mode = c.status.currentInstruction.Mode

		// Get the address of the data that the instruction will operate on
		var extraCycles = c.executeAddressingMode(c.i.addr_mode)

		// Now execute the instruction
		extraCycles += c.status.currentInstruction.Execute(c)

		// Add any extra cycles
		c.status.Cycles += extraCycles
	}

	// Decrement the number of cycles
	c.status.Cycles--
}

func (c *CPU) executeAddressingMode(mode AddressingMode) uint8 {
	c.i.addr_mode = mode

	// Get the address of the data that the instruction will operate on.
	switch mode {
	case Accumulator:
		return c.accumulator()
	case Implied:
		return c.implied()
	case Immediate:
		return c.immediate()
	case ZeroPage:
		return c.zeroPage()
	case ZeroPageX:
		return c.zeroPageX()
	case ZeroPageY:
		return c.zeroPageY()
	case Absolute:
		return c.absolute()
	case AbsoluteX:
		return c.absoluteX()
	case AbsoluteY:
		return c.absoluteY()
	case Indirect:
		return c.indirect()
	case IndexedIndirect:
		return c.indexedIndirect()
	case IndirectIndexed:
		return c.indirectIndexed()
	case Relative:
		return c.relative()
	default:
		return 0
	}
}

func (c *CPU) getOperandString(mode AddressingMode, address uint16) string {
	// Switch on the addressing mode
	switch mode {
	case Implied:
		return ""
	case Immediate:
		return fmt.Sprintf("#$%02X", c.bus.Read(address))
	case ZeroPage:
		return fmt.Sprintf("$%02X", c.bus.Read(address))
	case ZeroPageX:
		return fmt.Sprintf("$%02X,X", c.bus.Read(address))
	case ZeroPageY:
		return fmt.Sprintf("$%02X,Y", c.bus.Read(address))
	case Absolute:
		var temp uint16 = uint16(c.bus.Read(address)) | (uint16(c.bus.Read(address+1)) << 8)
		return fmt.Sprintf("$%04X", temp)
	case AbsoluteX:
		// We want the next two bytes
		var temp uint16 = uint16(c.bus.Read(address)) | (uint16(c.bus.Read(address+1)) << 8)
		return fmt.Sprintf("$%04X,X", temp)
	case AbsoluteY:
		var temp uint16 = uint16(c.bus.Read(address)) | (uint16(c.bus.Read(address+1)) << 8)
		return fmt.Sprintf("$%04X,Y", temp)
	case Indirect:
		var temp uint16 = uint16(c.bus.Read(address)) | (uint16(c.bus.Read(address+1)) << 8)
		return fmt.Sprintf("($%04X)", temp)
	case IndexedIndirect:
		return fmt.Sprintf("($%02X,X)", c.bus.Read(address))
	case IndirectIndexed:
		return fmt.Sprintf("($%02X),Y", c.bus.Read(address))
	case Relative:
		var temp uint16 = uint16(c.bus.Read(address)) | (uint16(c.bus.Read(address+1)) << 8)
		return fmt.Sprintf("$%04X", temp)
	default:
		return ""
	}
}

func (c *CPU) DisassembleAt(addr uint16) string {
	var opcode uint8 = c.bus.Read(addr)
	var instruction InstructionInfo = Instructions[opcode]
	var addrMode AddressingMode = instruction.Mode
	var addrString string = c.getOperandString(addrMode, addr+1)
	var insn = InstructionNames[instruction.Instruction]
	return fmt.Sprintf("%s %s", insn, addrString)
}

func (c *CPU) fetchByte() uint8 {
	// If the addressing mode is not implied or accumulator, read the data
	if c.i.addr_mode != Implied && c.i.addr_mode != Accumulator {
		c.i.fetched = c.bus.Read(c.i.addr_absolute)
	}
	return c.i.fetched
}

func (c *CPU) pushByte(data uint8) {
	c.bus.Write(0x100+uint16(c.r.sp), data)
	c.r.sp--
}

func (c *CPU) pushWord(data uint16) {
	c.pushByte(uint8(data >> 8))
	c.pushByte(uint8(data))
}

func (c *CPU) popByte() uint8 {
	c.r.sp++
	return c.bus.Read(0x100 + uint16(c.r.sp))
}

func (c *CPU) popWord() uint16 {
	var low uint16 = uint16(c.popByte())
	var high uint16 = uint16(c.popByte())
	return high<<8 | low
}

func (c *CPU) String() string {
	return fmt.Sprintf("Current Instruction: %s\nA: %02X X: %02X Y: %02X P: %02X SP: %02X PC: %04X\n",
		c.status.currentInstructionString, c.r.a, c.r.x, c.r.y, c.r.p, c.r.sp, c.r.pc)
}
