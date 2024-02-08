package main

import (
	"fmt"
	"time"

	"github.com/drewwalton19216801/goemu6502"
)

type Machine struct {
	bus *Bus
	cpu *goemu6502.CPU
}

func NewMachine() *Machine {
	m := &Machine{}
	m.bus = NewBus()
	m.cpu = goemu6502.NewCPU(m.bus)
	return m
}

func (m *Machine) Run() {
	// Clock the CPU five times per second, printing the CPU string every time
	for {
		m.cpu.Clock()
		fmt.Println(m.cpu.String())
		time.Sleep(time.Millisecond * 200)
	}
}

func (m *Machine) Reset() {
	m.cpu.Reset()
}
