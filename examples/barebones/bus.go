package main

type Bus struct {
	memory *Memory
}

func NewBus() *Bus {
	return &Bus{
		memory: NewMemory(),
	}
}

func (b *Bus) Read(addr uint16) uint8 {
	return b.memory.Read(addr)
}

func (b *Bus) Write(addr uint16, value uint8) {
	b.memory.Write(addr, value)
}
