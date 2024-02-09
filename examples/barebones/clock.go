package main

import (
	"fmt"

	"github.com/drewwalton19216801/goemu6502"
)

type Clock struct {
	cpu          *goemu6502.CPU
	sysFrequency int // ticks per second for the system clock
	cpuFrequency int // ticks per second for the CPU

	cpuTick chan struct{}
	sysTick chan struct{}

	sysTicks int
	cpuTicks int

	pause chan bool
}

func NewClock(cpu *goemu6502.CPU) *Clock {
	c := &Clock{
		cpu:          cpu,
		sysFrequency: 0,
		cpuFrequency: 0,
		cpuTick:      make(chan struct{}),
		sysTick:      make(chan struct{}),
		pause:        make(chan bool),
	}

	go c.run()
	return c
}

func (c *Clock) run() {
	paused := false
	for {
		select {
		case <-c.sysTick:
			c.sysTicks++

			if c.sysTicks%(c.sysFrequency/c.cpuFrequency) == 0 {
				select {
				case c.cpuTick <- struct{}{}:
				default:
				}
			}

		case <-c.cpuTick:
			if !paused {
				println(c.cpu.String())
				c.cpu.Tick()
				c.cpuTicks++
			}

		case paused = <-c.pause:
		// Do nothing

		case <-c.pause:
			paused = true
		}
	}
}

func (c *Clock) SysTick() {
	c.sysTick <- struct{}{}
}

func (c *Clock) CpuTick() {
	c.cpuTick <- struct{}{}
}

func (c *Clock) Pause() {
	c.pause <- true
}

func (c *Clock) Resume() {
	c.pause <- false
}

func (c *Clock) Step() {
	c.sysTick <- struct{}{}
	c.cpuTick <- struct{}{}

	<-c.sysTick
	<-c.cpuTick

	<-c.pause
}

func (c *Clock) SetCPUFrequency(frequency int) {
	fmt.Println("CPU Frequency: ", frequency)
	c.cpuFrequency = frequency
}

func (c *Clock) SetSysFrequency(frequency int) {
	fmt.Println("System Frequency: ", frequency)
	c.sysFrequency = frequency
}
