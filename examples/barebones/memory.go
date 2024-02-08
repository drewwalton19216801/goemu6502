package main

type Memory struct {
	data []uint8
}

func NewMemory() *Memory {
	return &Memory{
		data: make([]uint8, 0x10000), // 64K memory
	}
}

func (m *Memory) Read(addr uint16) uint8 {
	return m.data[addr]
}

func (m *Memory) Write(addr uint16, value uint8) {
	m.data[addr] = value
}
