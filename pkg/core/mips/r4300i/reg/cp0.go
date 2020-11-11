/*

System Control Coprocessor(CP0) Register

Register map:
	cp0[0](Index)    :  Programmable pointer into TLB array
	cp0[1](Random)   :  Pseudorandom pointer into TLB array (read only)
	cp0[2](EntryLo0) :  Low half of TLB entry for even virtual address (VPN)
	cp0[3](EntryLo1) :  Low half of TLB entry for odd virtual address (VPN)
	cp0[4](Context)  :  Pointer to kernel virtual page table entry (PTE) in 32-bit mode
	cp0[5](PageMask) :  Page size specification
	cp0[6](Wired)    :  Number of wired TLB entries
	cp0[7](-)        :  Reserved for future use
	cp0[8](BadVAddr) :  Display of virtual address that occurred an error last
	cp0[9](Count)    :  Timer Count
	cp0[10](EntryHi) :  High half of TLB entry (including ASID)
	cp0[11](Compare) :  Timer Compare Value
	cp0[12](Status)  :  Operation status setting
	cp0[13](Cause)   :  Display of cause of last exception
	cp0[14](EPC)     :  Exception Program Counter
	cp0[15](PRId)    :  Processor Revision Identifier
	cp0[16](Config)  :  Memory system mode setting
	cp0[17](LLAddr)  :  Load Linked instruction address display
	cp0[18](WatchLo) :  Memory reference trap address low bits
	cp0[19](WatchHi) :  Memory reference trap address high bits
	cp0[20](XContext):  Pointer to Kernel virtual PTE table in 64-bit mode
	cp0[21–25](-)    :  Reserved for future use
	cp0[26](Parity)  :  Error Cache parity bits
	cp0[27](Cache)   :  Error* Cache Error and Status register
	cp0[28](TagLo)   :  Cache Tag register low
	cp0[29](TagHi)   :  Cache Tag register high
	cp0[30](ErrorEPC):  Error Exception Program Counter
	cp0[31](-)       :  Reserved for future use

*/

package reg

const (
	NumOfRegsInCp0 = 32
)

type CP0 struct {
	cp0 [NumOfRegsInCp0]uint32
}

// Read value of the register.
func (cp0 *CP0) Read(index int) uint32 {
	return cp0.cp0[index]
}

// Write value in register
func (cp0 *CP0) Write(index int, value uint32) {
	cp0.cp0[index] = value
}
