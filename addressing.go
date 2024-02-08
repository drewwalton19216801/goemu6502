package goemu6502

type (
	AddressingMode uint8
)

const (
	_ AddressingMode = iota
	Accumulator
	Immediate
	Implied
	Indirect
	Relative
	IndexedIndirect
	IndirectIndexed
	Absolute
	AbsoluteX
	AbsoluteY
	ZeroPage
	ZeroPageX
	ZeroPageY
)

// AddressingModeNames is a map of addressing mode names
var AddressingModeNames = map[AddressingMode]string{
	Accumulator:     "Accumulator",
	Immediate:       "Immediate",
	Implied:         "Implied",
	Indirect:        "Indirect",
	Relative:        "Relative",
	IndexedIndirect: "IndexedIndirect",
	IndirectIndexed: "IndirectIndexed",
	Absolute:        "Absolute",
	AbsoluteX:       "AbsoluteX",
	AbsoluteY:       "AbsoluteY",
	ZeroPage:        "ZeroPage",
	ZeroPageX:       "ZeroPageX",
	ZeroPageY:       "ZeroPageY",
}

// --- Addressing modes ---
// Addressing mode functions calculate the effective address of an instruction,
// which is the address of the data that the instruction will operate on.
// Functions return the number of extra cycles that may be required to fetch
// the data.

// accumulator gets the accumulator value.
//
// No parameters.
// Returns uint8.
func (c *CPU) accumulator() uint8 {
	// Get the accumulator value, doesn't really matter what it is
	c.i.fetched = c.r.a

	return 0
}

// immediate Get the immediate value
//
// None
// uint8
func (c *CPU) immediate() uint8 {
	// Get the immediate value
	c.i.addr_absolute = c.r.pc

	// Advance the program counter
	c.r.pc++

	return 0
}

// implied gets the implied value.
//
// None.
// uint8.
func (c *CPU) implied() uint8 {
	c.i.fetched = c.r.a
	return 0
}

func (c *CPU) indirect() uint8 {
	// Read the indirect address
	c.i.temp = (uint16(c.bus.Read(c.r.pc+1)) << 8) | uint16(c.bus.Read(c.r.pc))

	// Get the indirect address
	if c.i.temp&0x00FF == 0x00FF {
		// Simulate page boundary bug
		c.i.addr_absolute = uint16(c.bus.Read(c.i.temp&0xFF00))<<8 | uint16(c.bus.Read(c.i.temp))
	} else {
		// Proceed as normal
		c.i.addr_absolute = uint16(c.bus.Read(c.i.temp+1))<<8 | uint16(c.bus.Read(c.i.temp))
	}

	return 0
}

// relative gets the relative address.
//
// No parameters.
// Returns a uint8.
func (c *CPU) relative() uint8 {
	// Get the relative address
	c.i.temp = uint16(int8(c.bus.Read(c.r.pc)))

	// Store the address
	c.i.addr_relative = c.i.temp

	// Increment the program counter
	c.r.pc += 2

	return 0
}

// indexedIndirect Get the address indexed by the X register.
//
// No parameters.
// Return type uint8.
func (c *CPU) indexedIndirect() uint8 {
	// Get the address indexed by the X register
	c.i.temp = uint16(c.bus.Read(c.r.pc))
	c.i.addr_absolute = uint16((c.i.temp + uint16(c.r.x)&0xFF))

	// Increment the program counter
	c.r.pc++

	return 0
}

// indirectIndexed returns the address indexed by the Y register.
//
// No parameters.
// Returns uint8.
func (c *CPU) indirectIndexed() uint8 {
	// Get the address indexed by the Y register
	c.i.temp = uint16(c.bus.Read(c.r.pc))
	c.i.addr_absolute = uint16((c.i.temp + uint16(c.r.y)&0xFF))

	// Increment the program counter
	c.r.pc++

	// Check if the page boundary was crossed, if so add another cycle
	if c.i.addr_absolute&0xFF00 != c.i.temp&0xFF00 {
		return 1
	}

	return 0
}

// zeroPage Get the zero page address.
//
// None.
// uint8.
func (c *CPU) zeroPage() uint8 {
	// Get the zero page address
	c.i.addr_absolute = uint16(c.bus.Read(c.r.pc) & 0xFF)

	// Increment the program counter
	c.r.pc++

	return 0
}

// absolute gets the absolute address
//
// No parameters
// Returns uint8
func (c *CPU) absolute() uint8 {
	// Get the absolute address
	c.i.addr_absolute = (uint16(c.bus.Read(c.r.pc+1)) << 8) | uint16(c.bus.Read(c.r.pc))

	// Increment the program counter
	c.r.pc += 2

	return 0
}

// absoluteX gets the absolute address and adds the X register. It also increments the program counter and checks if the page boundary was crossed, adding another cycle if so.
//
// None.
// uint8.
func (c *CPU) absoluteX() uint8 {
	// Get the absolute address
	c.i.addr_absolute = (uint16(c.bus.Read(c.r.pc+1)) << 8) | uint16(c.bus.Read(c.r.pc))
	// Add the X register
	c.i.addr_absolute += uint16(c.r.x)

	// Increment the program counter
	c.r.pc += 2

	// Check if the page boundary was crossed, if so add another cycle
	if (c.i.addr_absolute & 0xFF00) != ((c.i.addr_absolute - uint16(c.r.x)) & 0xFF00) {
		return 1
	}

	return 0
}

// absoluteY gets the absolute address and adds the Y register. It increments the program counter and checks if the page boundary was crossed, adding another cycle if necessary.
//
// No parameters.
// Returns uint8.
func (c *CPU) absoluteY() uint8 {
	// Get the absolute address
	c.i.addr_absolute = (uint16(c.bus.Read(c.r.pc+1)) << 8) | uint16(c.bus.Read(c.r.pc))
	// Add the Y register
	c.i.addr_absolute += uint16(c.r.y)

	// Increment the program counter
	c.r.pc += 2

	// Check if the page boundary was crossed, if so add another cycle
	if (c.i.addr_absolute & 0xFF00) != ((c.i.addr_absolute - uint16(c.r.y)) & 0xFF00) {
		return 1
	}

	return 0
}

// zeroPageX gets the zero page address, wrapping if necessary.
//
// No parameters.
// Returns uint8.
func (c *CPU) zeroPageX() uint8 {
	// Get the zero page address, wrapping if necessary
	c.i.addr_absolute = uint16((c.bus.Read(c.r.pc) + c.r.x) & 0xFF)

	// Increment the program counter
	c.r.pc++

	return 0
}

// zeroPageY Get the zero page address, wrapping if necessary.
//
// No parameters.
// Returns uint8.
func (c *CPU) zeroPageY() uint8 {
	// Get the zero page address, wrapping if necessary
	c.i.addr_absolute = uint16((c.bus.Read(c.r.pc) + c.r.y) & 0xFF)

	// Increment the program counter
	c.r.pc++

	return 0
}
