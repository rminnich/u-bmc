// Copyright 2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"github.com/u-root/u-bmc/pkg/aspeed"
)

var (
	linePortMap = map[string]uint32{
		"PWR_BTN_N":           aspeed.GpioPort("C1"),
		"RST_BTN_N":           aspeed.GpioPort("E2"),
		// Pull low to have BMC own the bus
		"BIOS_MUX_N":          aspeed.GpioPort("E4"),
		"PWR_LED_N":           aspeed.GpioPort("D6"),
		"UID_LED_N":           aspeed.GpioPort("D7"),
	}

	// Reverse map of linePortMap
	portLineMap map[uint32]string
)

type Gpio struct {
}

func init() {
	portLineMap = make(map[uint32]string)
	for k, v := range linePortMap {
		portLineMap[v] = k
	}
}

func (_ *Gpio) GpioNameToPort(l string) (uint32, bool) {
	s, ok := linePortMap[l]
	return s, ok
}

func (_ *Gpio) GpioPortToName(i uint32) (string, bool) {
	s, ok := portLineMap[i]
	return s, ok
}
