package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello 6502 World!")

	m := NewMachine()
	m.clock.SetCPUFrequency(50)
	m.clock.SetSysFrequency(m.clock.cpuFrequency * 10)

	// Set the reset vector to 0x8000
	m.bus.Write(0xFFFC, 0x00)
	m.bus.Write(0xFFFD, 0x80)

	// A very simple 6502 program that increments A by 1 forever
	// 0x8000: LDA #$01
	m.bus.Write(0x8000, 0xA9)
	m.bus.Write(0x8001, 0x01)
	// 0x8002: ADC #$01
	m.bus.Write(0x8002, 0x69)
	m.bus.Write(0x8003, 0x01)
	// 0x8004: JMP $8002
	m.bus.Write(0x8004, 0x4C)
	m.bus.Write(0x8005, 0x02)
	m.bus.Write(0x8006, 0x80)

	m.Reset()
	fmt.Println(m.cpu.String())
	m.Run()
}
