# Copyright 2019 u-root Authors
#
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file

SOC := ast2400
# The UARTs are in the following order: UART5, VUART, UART1, UART2, UART3, UART4
QEMUFLAGS += -m 128 -M palmetto-bmc \
	-serial mon:stdio \
	-serial null \
	-serial null \
	-serial null \
	-serial chardev:host \
	-serial null

RAM_SIZE := 0x07000000 # 128 MB - 16 MB (VGA)

sim: flash.sim.img
	$(QEMU) $(QEMUFLAGS)
	stty sane
