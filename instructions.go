package goemu6502

// Instruction is an enum for the different instructions
type Instruction uint8

// Legal instructions
const (
	_ Instruction = iota
	adc
	and
	asl
	bcc
	bcs
	beq
	bit
	bmi
	bne
	bpl
	brk
	bvc
	bvs
	clc
	cld
	cli
	clv
	cmp
	cpx
	cpy
	dec
	dex
	dey
	eor
	inc
	inx
	iny
	jmp
	jsr
	lda
	ldx
	ldy
	lsr
	nop
	ora
	PHA
	php
	pla
	plp
	rol
	ror
	rti
	rts
	sbc
	sec
	sed
	sei
	sta
	stx
	sty
	tax
	tay
	tsx
	txa
	txs
	tya
)

// InstructionNames is a map of instruction names
var InstructionNames = map[Instruction]string{
	adc: "adc",
	and: "and",
	asl: "asl",
	bcc: "bcc",
	bcs: "bcs",
	beq: "beq",
	bit: "bit",
	bmi: "bmi",
	bne: "bne",
	bpl: "bpl",
	brk: "brk",
	bvc: "bvc",
	bvs: "bvs",
	clc: "clc",
	cld: "cld",
	cli: "cli",
	clv: "clv",
	cmp: "cmp",
	cpx: "cpx",
	cpy: "cpy",
	dec: "dec",
	dex: "dex",
	dey: "dey",
	eor: "eor",
	inc: "inc",
	inx: "inx",
	iny: "iny",
	jmp: "jmp",
	jsr: "jsr",
	lda: "lda",
	ldx: "ldx",
	ldy: "ldy",
	lsr: "lsr",
	nop: "nop",
	ora: "ora",
	PHA: "PHA",
	php: "php",
	pla: "pla",
	plp: "plp",
	rol: "rol",
	ror: "ror",
	rti: "rti",
	rts: "rts",
	sbc: "sbc",
	sec: "sec",
	sed: "sed",
	sei: "sei",
	sta: "sta",
	stx: "stx",
	sty: "sty",
	tax: "tax",
	tay: "tay",
	tsx: "tsx",
	txa: "txa",
	txs: "txs",
	tya: "tya",
}

// InstructionInfo contains information about an instruction
type InstructionInfo struct {
	Instruction Instruction
	Opcode      uint8
	Mode        AddressingMode
	Cycles      uint8
	Execute     func(*CPU) uint8
}

// Instructions is a map of instruction infos
var Instructions = map[uint8]InstructionInfo{
	0x69: {adc, 0x69, Immediate, 2, (*CPU).adc},
	0x65: {adc, 0x65, ZeroPage, 3, (*CPU).adc},
	0x75: {adc, 0x75, ZeroPageX, 4, (*CPU).adc},
	0x6D: {adc, 0x6D, Absolute, 4, (*CPU).adc},
	0x7D: {adc, 0x7D, AbsoluteX, 4, (*CPU).adc},
	0x79: {adc, 0x79, AbsoluteY, 4, (*CPU).adc},
	0x61: {adc, 0x61, IndexedIndirect, 6, (*CPU).adc},
	0x71: {adc, 0x71, IndirectIndexed, 5, (*CPU).adc},
	0x29: {and, 0x29, Immediate, 2, (*CPU).and},
	0x25: {and, 0x25, ZeroPage, 3, (*CPU).and},
	0x35: {and, 0x35, ZeroPageX, 4, (*CPU).and},
	0x2D: {and, 0x2D, Absolute, 4, (*CPU).and},
	0x3D: {and, 0x3D, AbsoluteX, 4, (*CPU).and},
	0x39: {and, 0x39, AbsoluteY, 4, (*CPU).and},
	0x21: {and, 0x21, IndexedIndirect, 6, (*CPU).and},
	0x31: {and, 0x31, IndirectIndexed, 5, (*CPU).and},
	0x0A: {asl, 0x0A, Accumulator, 2, (*CPU).asl},
	0x06: {asl, 0x06, ZeroPage, 5, (*CPU).asl},
	0x16: {asl, 0x16, ZeroPageX, 6, (*CPU).asl},
	0x0E: {asl, 0x0E, Absolute, 6, (*CPU).asl},
	0x1E: {asl, 0x1E, AbsoluteX, 7, (*CPU).asl},
	0x90: {bcc, 0x90, Relative, 2, (*CPU).bcc},
	0xB0: {bcs, 0xB0, Relative, 2, (*CPU).bcs},
	0xF0: {beq, 0xF0, Relative, 2, (*CPU).beq},
	0x24: {bit, 0x24, ZeroPage, 3, (*CPU).bit},
	0x2C: {bit, 0x2C, Absolute, 4, (*CPU).bit},
	0x30: {bmi, 0x30, Relative, 2, (*CPU).bmi},
	0xD0: {bne, 0xD0, Relative, 2, (*CPU).bne},
	0x10: {bpl, 0x10, Relative, 2, (*CPU).bpl},
	0x00: {brk, 0x00, Implied, 7, (*CPU).brk},
	0x50: {bvc, 0x50, Relative, 2, (*CPU).bvc},
	0x70: {bvs, 0x70, Relative, 2, (*CPU).bvs},
	0x18: {clc, 0x18, Implied, 2, (*CPU).clc},
	0xD8: {cld, 0xD8, Implied, 2, (*CPU).cld},
	0x58: {cli, 0x58, Implied, 2, (*CPU).cli},
	0xB8: {clv, 0xB8, Implied, 2, (*CPU).clv},
	0xC9: {cmp, 0xC9, Immediate, 2, (*CPU).cmp},
	0xC5: {cmp, 0xC5, ZeroPage, 3, (*CPU).cmp},
	0xD5: {cmp, 0xD5, ZeroPageX, 4, (*CPU).cmp},
	0xCD: {cmp, 0xCD, Absolute, 4, (*CPU).cmp},
	0xDD: {cmp, 0xDD, AbsoluteX, 4, (*CPU).cmp},
	0xD9: {cmp, 0xD9, AbsoluteY, 4, (*CPU).cmp},
	0xC1: {cmp, 0xC1, IndexedIndirect, 6, (*CPU).cmp},
	0xD1: {cmp, 0xD1, IndirectIndexed, 5, (*CPU).cmp},
	0xE0: {cpx, 0xE0, Immediate, 2, (*CPU).cpx},
	0xE4: {cpx, 0xE4, ZeroPage, 3, (*CPU).cpx},
	0xEC: {cpx, 0xEC, Absolute, 4, (*CPU).cpx},
	0xC0: {cpy, 0xC0, Immediate, 2, (*CPU).cpy},
	0xC4: {cpy, 0xC4, ZeroPage, 3, (*CPU).cpy},
	0xCC: {cpy, 0xCC, Absolute, 4, (*CPU).cpy},
	0xC6: {dec, 0xC6, ZeroPage, 5, (*CPU).dec},
	0xD6: {dec, 0xD6, ZeroPageX, 6, (*CPU).dec},
	0xCE: {dec, 0xCE, Absolute, 6, (*CPU).dec},
	0xDE: {dec, 0xDE, AbsoluteX, 7, (*CPU).dec},
	0xCA: {dex, 0xCA, Implied, 2, (*CPU).dex},
	0x88: {dey, 0x88, Implied, 2, (*CPU).dey},
	0x49: {eor, 0x49, Immediate, 2, (*CPU).eor},
	0x45: {eor, 0x45, ZeroPage, 3, (*CPU).eor},
	0x55: {eor, 0x55, ZeroPageX, 4, (*CPU).eor},
	0x4D: {eor, 0x4D, Absolute, 4, (*CPU).eor},
	0x5D: {eor, 0x5D, AbsoluteX, 4, (*CPU).eor},
	0x59: {eor, 0x59, AbsoluteY, 4, (*CPU).eor},
	0x41: {eor, 0x41, IndexedIndirect, 6, (*CPU).eor},
	0x51: {eor, 0x51, IndirectIndexed, 5, (*CPU).eor},
	0xE6: {inc, 0xE6, ZeroPage, 5, (*CPU).inc},
	0xF6: {inc, 0xF6, ZeroPageX, 6, (*CPU).inc},
	0xEE: {inc, 0xEE, Absolute, 6, (*CPU).inc},
	0xFE: {inc, 0xFE, AbsoluteX, 7, (*CPU).inc},
	0xE8: {inx, 0xE8, Implied, 2, (*CPU).inx},
	0xC8: {iny, 0xC8, Implied, 2, (*CPU).iny},
	0x4C: {jmp, 0x4C, Absolute, 3, (*CPU).jmp},
	0x6C: {jmp, 0x6C, Indirect, 5, (*CPU).jmp},
	0x20: {jsr, 0x20, Absolute, 6, (*CPU).jsr},
	0xA9: {lda, 0xA9, Immediate, 2, (*CPU).lda},
	0xA5: {lda, 0xA5, ZeroPage, 3, (*CPU).lda},
	0xB5: {lda, 0xB5, ZeroPageX, 4, (*CPU).lda},
	0xAD: {lda, 0xAD, Absolute, 4, (*CPU).lda},
	0xBD: {lda, 0xBD, AbsoluteX, 4, (*CPU).lda},
	0xB9: {lda, 0xB9, AbsoluteY, 4, (*CPU).lda},
	0xA1: {lda, 0xA1, IndexedIndirect, 6, (*CPU).lda},
	0xB1: {lda, 0xB1, IndirectIndexed, 5, (*CPU).lda},
	0xA2: {ldx, 0xA2, Immediate, 2, (*CPU).ldx},
	0xA6: {ldx, 0xA6, ZeroPage, 3, (*CPU).ldx},
	0xB6: {ldx, 0xB6, ZeroPageY, 4, (*CPU).ldx},
	0xAE: {ldx, 0xAE, Absolute, 4, (*CPU).ldx},
	0xBE: {ldx, 0xBE, AbsoluteY, 4, (*CPU).ldx},
	0xA0: {ldy, 0xA0, Immediate, 2, (*CPU).ldy},
	0xA4: {ldy, 0xA4, ZeroPage, 3, (*CPU).ldy},
	0xB4: {ldy, 0xB4, ZeroPageX, 4, (*CPU).ldy},
	0xAC: {ldy, 0xAC, Absolute, 4, (*CPU).ldy},
	0xBC: {ldy, 0xBC, AbsoluteX, 4, (*CPU).ldy},
	0x4A: {lsr, 0x4A, Accumulator, 2, (*CPU).lsr},
	0x46: {lsr, 0x46, ZeroPage, 5, (*CPU).lsr},
	0x56: {lsr, 0x56, ZeroPageX, 6, (*CPU).lsr},
	0x4E: {lsr, 0x4E, Absolute, 6, (*CPU).lsr},
	0x5E: {lsr, 0x5E, AbsoluteX, 7, (*CPU).lsr},
	0xEA: {nop, 0xEA, Implied, 2, (*CPU).nop},
	0x09: {ora, 0x09, Immediate, 2, (*CPU).ora},
	0x05: {ora, 0x05, ZeroPage, 3, (*CPU).ora},
	0x15: {ora, 0x15, ZeroPageX, 4, (*CPU).ora},
	0x0D: {ora, 0x0D, Absolute, 4, (*CPU).ora},
	0x1D: {ora, 0x1D, AbsoluteX, 4, (*CPU).ora},
	0x19: {ora, 0x19, AbsoluteY, 4, (*CPU).ora},
	0x01: {ora, 0x01, IndexedIndirect, 6, (*CPU).ora},
	0x11: {ora, 0x11, IndirectIndexed, 5, (*CPU).ora},
	0x48: {PHA, 0x48, Implied, 3, (*CPU).PHA},
	0x08: {php, 0x08, Implied, 3, (*CPU).php},
	0x68: {pla, 0x68, Implied, 4, (*CPU).pla},
	0x28: {plp, 0x28, Implied, 4, (*CPU).plp},
	0x2A: {rol, 0x2A, Accumulator, 2, (*CPU).rol},
	0x26: {rol, 0x26, ZeroPage, 5, (*CPU).rol},
	0x36: {rol, 0x36, ZeroPageX, 6, (*CPU).rol},
	0x2E: {rol, 0x2E, Absolute, 6, (*CPU).rol},
	0x3E: {rol, 0x3E, AbsoluteX, 7, (*CPU).rol},
	0x6A: {ror, 0x6A, Accumulator, 2, (*CPU).ror},
	0x66: {ror, 0x66, ZeroPage, 5, (*CPU).ror},
	0x76: {ror, 0x76, ZeroPageX, 6, (*CPU).ror},
	0x6E: {ror, 0x6E, Absolute, 6, (*CPU).ror},
	0x7E: {ror, 0x7E, AbsoluteX, 7, (*CPU).ror},
	0x40: {rti, 0x40, Implied, 6, (*CPU).rti},
	0x60: {rts, 0x60, Implied, 6, (*CPU).rts},
	0xE9: {sbc, 0xE9, Immediate, 2, (*CPU).sbc},
	0xE5: {sbc, 0xE5, ZeroPage, 3, (*CPU).sbc},
	0xF5: {sbc, 0xF5, ZeroPageX, 4, (*CPU).sbc},
	0xED: {sbc, 0xED, Absolute, 4, (*CPU).sbc},
	0xFD: {sbc, 0xFD, AbsoluteX, 4, (*CPU).sbc},
	0xF9: {sbc, 0xF9, AbsoluteY, 4, (*CPU).sbc},
	0xE1: {sbc, 0xE1, IndexedIndirect, 6, (*CPU).sbc},
	0xF1: {sbc, 0xF1, IndirectIndexed, 5, (*CPU).sbc},
	0x38: {sec, 0x38, Implied, 2, (*CPU).sec},
	0xF8: {sed, 0xF8, Implied, 2, (*CPU).sed},
	0x78: {sei, 0x78, Implied, 2, (*CPU).sei},
	0x85: {sta, 0x85, ZeroPage, 3, (*CPU).sta},
	0x95: {sta, 0x95, ZeroPageX, 4, (*CPU).sta},
	0x8D: {sta, 0x8D, Absolute, 4, (*CPU).sta},
	0x9D: {sta, 0x9D, AbsoluteX, 5, (*CPU).sta},
	0x99: {sta, 0x99, AbsoluteY, 5, (*CPU).sta},
	0x81: {sta, 0x81, IndexedIndirect, 6, (*CPU).sta},
	0x91: {sta, 0x91, IndirectIndexed, 6, (*CPU).sta},
	0x86: {stx, 0x86, ZeroPage, 3, (*CPU).stx},
	0x96: {stx, 0x96, ZeroPageY, 4, (*CPU).stx},
	0x8E: {stx, 0x8E, Absolute, 4, (*CPU).stx},
	0x84: {sty, 0x84, ZeroPage, 3, (*CPU).sty},
	0x94: {sty, 0x94, ZeroPageX, 4, (*CPU).sty},
	0x8C: {sty, 0x8C, Absolute, 4, (*CPU).sty},
	0xAA: {tax, 0xAA, Implied, 2, (*CPU).tax},
	0xA8: {tay, 0xA8, Implied, 2, (*CPU).tay},
	0xBA: {tsx, 0xBA, Implied, 2, (*CPU).tsx},
	0x8A: {txa, 0x8A, Implied, 2, (*CPU).txa},
	0x9A: {txs, 0x9A, Implied, 2, (*CPU).txs},
	0x98: {tya, 0x98, Implied, 2, (*CPU).tya},
}

// placeholder for illegal instructions
func (c *CPU) XXX() uint8 {
	// Panic
	panic("Illegal or unsupported instruction")
}

// adc adds with carry
func (c *CPU) adc() uint8 {
	var extraCycle uint8 = 0

	// Fetch the next byte
	c.fetchByte()

	// Add the fetched byte to the accumulator
	c.i.temp = uint16(c.r.a) + uint16(c.i.fetched)

	// Add the carry flag
	if c.getFlag(Carry) {
		c.i.temp++
	}

	// Set the Z flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Are we in decimal mode?
	if c.getFlag(Decimal) {
		// If the result is greater than 99, add 96 to it
		if (c.r.a&0x0F)+(c.i.fetched&0x0F)+uint8(func() uint8 {
			if c.getFlag(Carry) {
				return 1
			}
			return 0
		}()) > 9 {
			c.i.temp += 96 + uint16(func() uint8 {
				if c.getFlag(Carry) {
					return 1
				}
				return 0
			}())
		}

		// Set the negative flag if the result is negative
		c.setFlag(Negative, c.i.temp&0x80 != 0)

		// Set the overflow flag if the result is greater than 127 or less than -128
		c.setFlag(Overflow, (c.r.a^uint8(c.i.temp))&(uint8(c.i.fetched)^uint8(c.i.temp))&0x80 != 0)

		// If the result is greater than 99, add 96 to it
		if c.i.temp > 99 {
			c.i.temp += 96
		}

		// Set the carry flag if the result is greater than 99
		c.setFlag(Carry, c.i.temp > 99)

		// We used an extra cycle
		extraCycle++
	}

	// Store the result in the accumulator
	c.r.a = uint8(c.i.temp & 0x00FF)

	// Return the number of extra cycles
	return extraCycle
}

// and ands with accumulator
func (c *CPU) and() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the bitwise and operation
	c.r.a = c.r.a & c.i.fetched

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.r.a == 0)
	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.r.a&0x80 != 0)

	return 1
}

// asl shifts left one bit
func (c *CPU) asl() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Shift the fetched byte left by 1 bit
	c.i.temp = uint16(c.i.fetched) << 1

	// Set the carry flag if the 9th bit is set
	c.setFlag(Carry, c.i.temp&0xFF00 != 0)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the 7th bit is set
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	// If we are in accumulator mode, store the result in the accumulator
	if c.status.currentInstruction.Mode == Accumulator {
		c.r.a = uint8(c.i.temp & 0x00FF)
	} else {
		// Otherwise, store the result in memory
		c.bus.Write(c.i.addr_absolute, uint8(c.i.temp&0x00FF))
	}

	return 0
}

// bcc branches if carry clear
func (c *CPU) bcc() uint8 {
	extraCycles := uint8(0)
	// If the carry flag is not set, branch
	if !c.getFlag(Carry) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// bcs branches if carry set
func (c *CPU) bcs() uint8 {
	extraCycles := uint8(0)

	// If the carry flag is set, branch
	if c.getFlag(Carry) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// beq branches if equal
func (c *CPU) beq() uint8 {
	extraCycles := uint8(0)

	// If the zero flag is set, branch
	if c.getFlag(Zero) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// bit tests bits in memory with accumulator
func (c *CPU) bit() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the bitwise and operation
	c.i.temp = uint16(c.r.a) & uint16(c.i.fetched)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the 7th bit of the fetched byte is set
	c.setFlag(Negative, c.i.fetched&(1<<7) != 0)

	// Set the overflow flag if the 6th bit of the fetched byte is set
	c.setFlag(Overflow, c.i.fetched&(1<<6) != 0)

	return 0
}

// bmi branches if minus
func (c *CPU) bmi() uint8 {
	extraCycles := uint8(0)

	// If the negative flag is set, branch
	if c.getFlag(Negative) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// bne branches if not equal
func (c *CPU) bne() uint8 {
	extraCycles := uint8(0)

	// If the zero flag is not set, branch
	if !c.getFlag(Zero) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// bpl branches if positive
func (c *CPU) bpl() uint8 {
	extraCycles := uint8(0)

	// If the negative flag is not set, branch
	if !c.getFlag(Negative) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// brk forces an interrupt
func (c *CPU) brk() uint8 {
	// increment the program counter
	c.r.pc++

	// Set the interrupt disable flag to 1
	c.setFlag(InterruptDisable, true)

	// Push the PC to the stack
	c.pushWord(c.r.pc)

	// Set the break flag
	c.setFlag(Break, true)

	// Push the processor status to the stack
	c.pushByte(c.r.p)

	// Clear the break flag
	c.setFlag(Break, false)

	// Set the PC to the data at the interrupt vector
	c.r.pc = uint16(c.bus.Read(0xFFFE)) | uint16(c.bus.Read(0xFFFF))<<8

	return 0
}

// bvc branches if overflow clear
func (c *CPU) bvc() uint8 {
	extraCycles := uint8(0)

	// If the overflow flag is not set, branch
	if !c.getFlag(Overflow) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// bvs branches if overflow set
func (c *CPU) bvs() uint8 {
	extraCycles := uint8(0)

	// If the overflow flag is set, branch
	if c.getFlag(Overflow) {
		// We branched, so add a cycle
		extraCycles++

		// Calculate the new address
		c.i.addr_absolute = c.r.pc + c.i.addr_relative

		// If the new address crosses a page boundary, add another cycle
		if (c.i.addr_absolute & 0xFF00) != (c.r.pc & 0xFF00) {
			extraCycles++
		}

		// Set the program counter to the new address
		c.r.pc = c.i.addr_absolute
	}

	return extraCycles
}

// clc clears carry flag
func (c *CPU) clc() uint8 {
	// Clear the carry flag
	c.setFlag(Carry, false)

	return 0
}

// cld clears decimal mode
func (c *CPU) cld() uint8 {
	// Clear the decimal flag
	c.setFlag(Decimal, false)

	return 0
}

// cli clears interrupt disable
func (c *CPU) cli() uint8 {
	// Clear the interrupt disable flag
	c.setFlag(InterruptDisable, false)

	return 0
}

// clv clears overflow flag
func (c *CPU) clv() uint8 {
	// Clear the overflow flag
	c.setFlag(Overflow, false)

	return 0
}

// cmp compares accumulator with memory
func (c *CPU) cmp() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the comparison, wrapping around if necessary
	c.i.temp = uint16(c.r.a) - uint16(c.i.fetched)

	// Set the carry flag if the result is greater than or equal to the fetched byte
	c.setFlag(Carry, c.r.a >= c.i.fetched)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	return 1
}

// cpx compares X register
func (c *CPU) cpx() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the comparison, wrapping around if necessary
	c.i.temp = uint16(c.r.x) - uint16(c.i.fetched)

	// Set the carry flag if the result is greater than or equal to the fetched byte
	c.setFlag(Carry, c.r.x >= c.i.fetched)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	return 0
}

// cpy compares Y register
func (c *CPU) cpy() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the comparison, wrapping around if necessary
	c.i.temp = uint16(c.r.y) - uint16(c.i.fetched)

	// Set the carry flag if the result is greater than or equal to the fetched byte
	c.setFlag(Carry, c.r.y >= c.i.fetched)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	return 0
}

// dec decrements memory
func (c *CPU) dec() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// decrement the fetched byte
	c.i.temp = uint16(c.i.fetched) - 1

	// Store the result in memory
	c.bus.Write(c.i.addr_absolute, uint8(c.i.temp&0x00FF))

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	return 0
}

// dex decrements X register
func (c *CPU) dex() uint8 {
	// decrement the X register
	c.r.x--

	// Set the zero flag if the X register is zero
	c.setFlag(Zero, c.r.x == 0)

	// Set the negative flag if the 7th bit of the X register is set
	c.setFlag(Negative, c.r.x&0x80 != 0)

	return 0
}

// dey decrements Y register
func (c *CPU) dey() uint8 {
	// decrement the Y register
	c.r.y--

	// Set the zero flag if the Y register is zero
	c.setFlag(Zero, c.r.y == 0)

	// Set the negative flag if the 7th bit of the Y register is set
	c.setFlag(Negative, c.r.y&0x80 != 0)

	return 0
}

// eor exclusive ors accumulator
func (c *CPU) eor() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the bitwise XOR operation
	c.r.a = c.r.a ^ c.i.fetched

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.r.a == 0)

	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.r.a&0x80 != 0)

	return 1
}

// inc increments memory
func (c *CPU) inc() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// increment the fetched byte
	c.i.temp = uint16(c.i.fetched) + 1

	// Store the result in memory
	c.bus.Write(c.i.addr_absolute, uint8(c.i.temp&0x00FF))

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	return 0
}

// inx increments X register
func (c *CPU) inx() uint8 {
	// increment the X register
	c.r.x++

	// Set the zero flag if the X register is zero
	c.setFlag(Zero, c.r.x == 0)

	// Set the negative flag if the 7th bit of the X register is set
	c.setFlag(Negative, c.r.x&0x80 != 0)

	return 0
}

// iny increments Y register
func (c *CPU) iny() uint8 {
	// increment the Y register
	c.r.y++

	// Set the zero flag if the Y register is zero
	c.setFlag(Zero, c.r.y == 0)

	// Set the negative flag if the 7th bit of the Y register is set
	c.setFlag(Negative, c.r.y&0x80 != 0)

	return 0
}

// jmp jumps to location
func (c *CPU) jmp() uint8 {
	// Set the program counter to the absolute address
	c.r.pc = c.i.addr_absolute

	return 0
}

// jsr jumps to subroutine
func (c *CPU) jsr() uint8 {
	// Push the program counter to the stack
	c.pushWord(c.r.pc - 1)

	// Set the program counter to the absolute address
	c.r.pc = c.i.addr_absolute

	return 0
}

// lda loads accumulator
func (c *CPU) lda() uint8 {
	// Fetch the next byte and store it in the accumulator
	c.r.a = c.fetchByte()

	// Set the zero and negative flags as appropriate
	c.setFlag(Zero, c.r.a == 0)
	c.setFlag(Negative, c.r.a&0x80 != 0)

	return 1
}

// ldx loads X register
func (c *CPU) ldx() uint8 {
	// Fetch the next byte and store it in the X register
	c.r.x = c.fetchByte()

	// Set the zero and negative flags as appropriate
	c.setFlag(Zero, c.r.x == 0)
	c.setFlag(Negative, c.r.x&0x80 != 0)

	return 1
}

// ldy loads Y register
func (c *CPU) ldy() uint8 {
	// Fetch the next byte and store it in the Y register
	c.r.y = c.fetchByte()

	// Set the zero and negative flags as appropriate
	c.setFlag(Zero, c.r.y == 0)
	c.setFlag(Negative, c.r.y&0x80 != 0)

	return 1
}

// lsr shifts right one bit
func (c *CPU) lsr() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Shift the fetched byte right by 1 bit
	c.i.temp = uint16(c.i.fetched) >> 1

	// Set the carry flag if the 9th bit is set
	c.setFlag(Carry, c.i.fetched&0x01 != 0)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the 7th bit is set
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	// If we are in accumulator mode, store the result in the accumulator
	if c.status.currentInstruction.Mode == Accumulator {
		c.r.a = uint8(c.i.temp & 0x00FF)
	} else {
		// Otherwise, store the result in memory
		c.bus.Write(c.i.addr_absolute, uint8(c.i.temp&0x00FF))
	}

	return 0
}

// nop no operation
func (c *CPU) nop() uint8 {
	// Perform no operation
	return 0
}

// ora ors accumulator
func (c *CPU) ora() uint8 {
	// Fetch the next byte
	c.fetchByte()

	// Perform the bitwise OR operation
	c.r.a = c.r.a | c.i.fetched

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.r.a == 0)
	// Set the negative flag if the result is negative
	c.setFlag(Negative, c.r.a&0x80 != 0)

	return 1
}

// PHA pushes accumulator
func (c *CPU) PHA() uint8 {
	// Push the accumulator to the stack
	c.pushByte(c.r.a)

	return 0
}

// php pushes processor status
func (c *CPU) php() uint8 {
	// Push the flags to the stack
	c.pushByte(c.r.p | Break | Unused)

	// Clear the break flag
	c.setFlag(Break, false)

	// Set the unused flag
	c.setFlag(Unused, true)

	return 0
}

// pla pulls accumulator
func (c *CPU) pla() uint8 {
	// Pop the next byte from the stack and store it in the accumulator
	c.r.a = c.popByte()

	// Set the zero flag if the accumulator is zero
	c.setFlag(Zero, c.r.a == 0)

	// Set the negative flag if the 7th bit of the accumulator is set
	c.setFlag(Negative, c.r.a&0x80 != 0)

	return 0
}

// plp pulls processor status
func (c *CPU) plp() uint8 {
	// Pop the status flags from the stack
	c.r.p = c.popByte()

	// Set the unused flag
	c.setFlag(Unused, true)

	return 0
}

// rol rotates left one bit
func (c *CPU) rol() uint8 {
	// If we are in Accumulator mode, load the accumulator into the fetched byte
	if c.status.currentInstruction.Mode == Accumulator {
		c.i.fetched = c.r.a
	} else {
		// Otherwise, fetch the next byte
		c.fetchByte()
	}

	// Shift the fetched byte left by 1 bit
	c.i.temp = uint16(c.i.fetched) << 1

	// If the carry flag is set, set the 0th bit of the temp variable
	if c.getFlag(Carry) {
		c.i.temp |= 0x01
	}

	// Set the carry flag if the 9th bit of the temp variable is set
	c.setFlag(Carry, c.i.temp&0xFF00 != 0)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the 7th bit of the temp variable is set
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	// If we are in accumulator mode, store the result in the accumulator
	if c.status.currentInstruction.Mode == Accumulator {
		c.r.a = uint8(c.i.temp & 0x00FF)
	} else {
		// Otherwise, store the result in memory
		c.bus.Write(c.i.addr_absolute, uint8(c.i.temp&0x00FF))
	}

	return 0
}

// ror rotates right one bit
func (c *CPU) ror() uint8 {
	// If we are in Accumulator mode, load the accumulator into the fetched byte
	if c.status.currentInstruction.Mode == Accumulator {
		c.i.fetched = c.r.a
	} else {
		// Otherwise, fetch the next byte
		c.fetchByte()
	}

	// Shift the fetched byte right by 1 bit
	c.i.temp = uint16(c.i.fetched) >> 1

	// If the carry flag is set, set the 7th bit of the temp variable
	if c.getFlag(Carry) {
		c.i.temp |= 0x80
	}

	// Set the carry flag if the 0th bit of the temp variable is set
	c.setFlag(Carry, c.i.temp&0x01 != 0)

	// Set the zero flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Set the negative flag if the 7th bit of the temp variable is set
	c.setFlag(Negative, c.i.temp&0x80 != 0)

	// If we are in accumulator mode, store the result in the accumulator
	if c.status.currentInstruction.Mode == Accumulator {
		c.r.a = uint8(c.i.temp & 0x00FF)
	} else {
		// Otherwise, store the result in memory
		c.bus.Write(c.i.addr_absolute, uint8(c.i.temp&0x00FF))
	}

	return 0
}

// rti returns from interrupt
func (c *CPU) rti() uint8 {
	// Pop the status register from the stack
	c.r.p = c.popByte()

	// Clear the break flag
	c.setFlag(Break, false)

	// Clear the unused flag
	c.setFlag(Unused, false)

	// Pop the program counter from the stack
	c.r.pc = c.popWord()

	return 0
}

// rts returns from subroutine
func (c *CPU) rts() uint8 {
	// Pop the program counter from the stack
	c.r.pc = c.popWord() + 1

	return 0
}

// sbc subtracts with carry
func (c *CPU) sbc() uint8 {
	var extraCycle uint8 = 0

	// Fetch the next byte
	c.fetchByte()

	// Subtract the fetched byte from the accumulator
	c.i.temp = uint16(c.r.a) - uint16(c.i.fetched)

	// Subtract the carry flag
	if c.getFlag(Carry) {
		c.i.temp--
	}

	// Set the Z flag if the result is zero
	c.setFlag(Zero, c.i.temp&0x00FF == 0)

	// Are we in decimal mode?
	if c.getFlag(Decimal) {
		// If the result is greater than 99, subtract 96 from it
		if (c.r.a&0x0F)-(c.i.fetched&0x0F)-uint8(func() uint8 {
			if c.getFlag(Carry) {
				return 0
			}
			return 1
		}()) < uint8(0) {
			c.i.temp -= 96 - uint16(func() uint8 {
				if c.getFlag(Carry) {
					return 0
				}
				return 1
			}())
		}

		// Set the negative flag if the result is negative
		c.setFlag(Negative, c.i.temp&0x80 != 0)

		// Set the overflow flag if the result is greater than 127 or less than -128
		c.setFlag(Overflow, (c.r.a^uint8(c.i.temp))&(uint8(c.i.fetched)^uint8(c.i.temp))&0x80 != 0)

		// Adjust the result for BCD
		if c.i.temp > 99 {
			c.i.temp -= 96
		}

		// Set the carry flag if the result is greater than 99
		c.setFlag(Carry, c.i.temp > 99)

		// We used an extra cycle
		extraCycle++
	}

	// Store the result in the accumulator
	c.r.a = uint8(c.i.temp & 0x00FF)

	// Return the number of extra cycles
	return extraCycle
}

// sec sets carry flag
func (c *CPU) sec() uint8 {
	// Set the carry flag
	c.setFlag(Carry, true)

	return 0
}

// sed sets decimal mode
func (c *CPU) sed() uint8 {
	// Set the decimal flag
	c.setFlag(Decimal, true)

	return 0
}

// sei sets interrupt disable
func (c *CPU) sei() uint8 {
	// Set the interrupt disable flag
	c.setFlag(InterruptDisable, true)

	return 0
}

// sta stores accumulator
func (c *CPU) sta() uint8 {
	// Store the accumulator at the absolute address
	c.bus.Write(c.i.addr_absolute, c.r.a)

	return 0
}

// stx stores X register
func (c *CPU) stx() uint8 {
	// Store the X register at the absolute address
	c.bus.Write(c.i.addr_absolute, c.r.x)

	return 0
}

// sty stores Y register
func (c *CPU) sty() uint8 {
	// Store the Y register at the absolute address
	c.bus.Write(c.i.addr_absolute, c.r.y)

	return 0
}

// tax transfers accumulator to X register
func (c *CPU) tax() uint8 {
	// Load the accumulator into the X register
	c.r.x = c.r.a

	return 0
}

// tay transfers accumulator to Y register
func (c *CPU) tay() uint8 {
	// Load the accumulator into the Y register
	c.r.y = c.r.a

	return 0
}

// tsx transfers stack pointer to X register
func (c *CPU) tsx() uint8 {
	// Load the stack pointer into the X register
	c.r.x = c.r.sp

	return 0
}

// txa transfers X register to accumulator
func (c *CPU) txa() uint8 {
	// Load the X register into the accumulator
	c.r.a = c.r.x

	return 0
}

// txs transfers X register to stack pointer
func (c *CPU) txs() uint8 {
	// Load the X register into the stack pointer
	c.r.sp = c.r.x

	return 0
}

// tya transfers Y register to accumulator
func (c *CPU) tya() uint8 {
	// Load the Y register into the accumulator
	c.r.a = c.r.y

	return 0
}
