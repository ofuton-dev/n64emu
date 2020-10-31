package reg

type Register struct {
	GPR
	FGR

	// Program Counter
	PC uint64

	// Multiply/Divide Registers
	HI, LO uint64

	// Implementation/Revision Information
	FCR0 uint32

	// Control/Status
	FCR31 uint32

	// Load/Link Bit
	LLBit bool
}
