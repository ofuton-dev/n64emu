package reg

type Register struct {
	// General-purpose registers
	GPR

	// Floating Point general-purpose registers
	FGR

	// System Control Coprocessor(CP0) Register
	CP0

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
