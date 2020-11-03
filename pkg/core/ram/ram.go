/*

	n64 memory map (http://n64.icequake.net/mirror/www.jimb.de/Projects/N64TEK.htm#memorymapoverview)
	00000000-03EFFFFF   RDRAM Memory (RDRAM = Rambus DRAM)
	03F00000-03FFFFFF   RDRAM Registers
	04000000-040FFFFF   SP Registers
	04100000-041FFFFF   DP Command Registers
	04200000-042FFFFF   DP Span Registers
	04300000-043FFFFF   MIPS Interface (MI) Registers
	04400000-044FFFFF   Video Interface (VI) Registers
	04500000-045FFFFF   Audio Interface (AI) Registers
	04600000-046FFFFF   Peripheral Interface (PI) Registers
	04700000-047FFFFF   RDRAM Interface (RI) Registers
	04800000-048FFFFF   Serial Interface (SI) Registers
	04900000-04FFFFFF   Unused
	05000000-05FFFFFF   Cartridge Domain 2 Address 1
	06000000-07FFFFFF   Cartridge Domain 1 Address 1
	08000000-0FFFFFFF   Cartridge Domain 2 Address 2
	10000000-1FBFFFFF   Cartridge Domain 1 Address 2
	1FC00000-1FC007BF   PIF Boot ROM
	1FC007C0-1FC007FF   PIF RAM
	1FC00800-1FCFFFFF   Reserved
	1FD00000-7FFFFFFF   Cartridge Domain 1 Address 3
	80000000-FFFFFFFF   External SysAD Device

*/

package ram
