package main

import (
	"time"

	"github.com/drewwalton19216801/goemu6502"
)

type Machine struct {
	bus   *Bus
	clock *Clock
	cpu   *goemu6502.CPU
}

func NewMachine() *Machine {
	m := &Machine{}
	m.bus = NewBus()
	m.cpu = goemu6502.NewCPU(m.bus)
	m.clock = NewClock(m.cpu)
	return m
}

func (m *Machine) Run() {
	go m.clock.run()

	m.clock.Resume()

	// Clock the CPU five times per second, printing the CPU string every time
	for {
		m.clock.SysTick()
		time.Sleep(time.Second / time.Duration(m.clock.sysFrequency))
	}
}

func (m *Machine) Reset() {
	m.cpu.Reset()
}
