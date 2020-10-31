package reg

// GPR General Purpose Register are reserved for integer operations while the other thirty-two register
type GPR struct {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31 uint64
}

// R0 = always zero. Any attempts to modify this register silently fail.
func (g *GPR) R0() uint64         { return 0 }
func (g *GPR) SetR0(value uint64) { return }

// R1 = AT = assembler temporary. Free use.
func (g *GPR) R1() uint64         { return g.r1 }
func (g *GPR) SetR1(value uint64) { g.r1 = value }

// R2 = V0 = arithmetic values, function return values.
func (g *GPR) R2() uint64         { return g.r2 }
func (g *GPR) SetR2(value uint64) { g.r2 = value }

// R3 = V1 = arithmetic values, function return values.
func (g *GPR) R3() uint64         { return g.r3 }
func (g *GPR) SetR3(value uint64) { g.r3 = value }

// R4 = A0 = parameter passing to subroutines. Formal but not rigid.
func (g *GPR) R4() uint64         { return g.r4 }
func (g *GPR) SetR4(value uint64) { g.r4 = value }

// R5 = A1 = parameter passing to subroutines. Formal but not rigid.
func (g *GPR) R5() uint64         { return g.r5 }
func (g *GPR) SetR5(value uint64) { g.r5 = value }

// R6 = A2 = parameter passing to subroutines. Formal but not rigid.
func (g *GPR) R6() uint64         { return g.r6 }
func (g *GPR) SetR6(value uint64) { g.r6 = value }

// R7 = A3 = parameter passing to subroutines. Formal but not rigid.
func (g *GPR) R7() uint64         { return g.r7 }
func (g *GPR) SetR7(value uint64) { g.r7 = value }

// R8 = T0 = scratch registers. CPU RAM.
func (g *GPR) R8() uint64         { return g.r8 }
func (g *GPR) SetR8(value uint64) { g.r8 = value }

// R9 = T1 = scratch registers. CPU RAM.
func (g *GPR) R9() uint64         { return g.r9 }
func (g *GPR) SetR9(value uint64) { g.r9 = value }

// R10 = T2 = scratch registers. CPU RAM.
func (g *GPR) R10() uint64         { return g.r10 }
func (g *GPR) SetR10(value uint64) { g.r10 = value }

// R11 = T3 = scratch registers. CPU RAM.
func (g *GPR) R11() uint64         { return g.r11 }
func (g *GPR) SetR11(value uint64) { g.r11 = value }

// R12 = T4 = scratch registers. CPU RAM.
func (g *GPR) R12() uint64         { return g.r12 }
func (g *GPR) SetR12(value uint64) { g.r12 = value }

// R13 = T5 = scratch registers. CPU RAM.
func (g *GPR) R13() uint64         { return g.r13 }
func (g *GPR) SetR13(value uint64) { g.r13 = value }

// R14 = T6 = scratch registers. CPU RAM.
func (g *GPR) R14() uint64         { return g.r14 }
func (g *GPR) SetR14(value uint64) { g.r14 = value }

// R15 = T7 = scratch registers. CPU RAM.
func (g *GPR) R15() uint64         { return g.r15 }
func (g *GPR) SetR15(value uint64) { g.r15 = value }

// R16 = S0 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R16() uint64         { return g.r16 }
func (g *GPR) SetR16(value uint64) { g.r16 = value }

// R17 = S1 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R17() uint64         { return g.r17 }
func (g *GPR) SetR17(value uint64) { g.r17 = value }

// R18 = S2 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R18() uint64         { return g.r18 }
func (g *GPR) SetR18(value uint64) { g.r18 = value }

// R19 = S3 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R19() uint64         { return g.r19 }
func (g *GPR) SetR19(value uint64) { g.r19 = value }

// R20 = S4 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R20() uint64         { return g.r20 }
func (g *GPR) SetR20(value uint64) { g.r20 = value }

// R21 = S5 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R21() uint64         { return g.r21 }
func (g *GPR) SetR21(value uint64) { g.r21 = value }

// R22 = S6 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R22() uint64         { return g.r22 }
func (g *GPR) SetR22(value uint64) { g.r22 = value }

// R23 = S7 = registers saved upon function protocol. Trash at will if you know how.
func (g *GPR) R23() uint64         { return g.r23 }
func (g *GPR) SetR23(value uint64) { g.r23 = value }

// R24 = T8 = scratch registers. CPU RAM.
func (g *GPR) R24() uint64         { return g.r24 }
func (g *GPR) SetR24(value uint64) { g.r24 = value }

// R25 = T9 = scratch registers. CPU RAM.
func (g *GPR) R25() uint64         { return g.r25 }
func (g *GPR) SetR25(value uint64) { g.r25 = value }

// R26 = K0
func (g *GPR) R26() uint64         { return g.r26 }
func (g *GPR) SetR26(value uint64) { g.r26 = value }

// R27 = K1
func (g *GPR) R27() uint64         { return g.r27 }
func (g *GPR) SetR27(value uint64) { g.r27 = value }

// R28 = GP
func (g *GPR) R28() uint64         { return g.r28 }
func (g *GPR) SetR28(value uint64) { g.r28 = value }

// R29 = SP = stack pointer
// Informal.
func (g *GPR) R29() uint64         { return g.r29 }
func (g *GPR) SetR29(value uint64) { g.r29 = value }

// R30 = S8
func (g *GPR) R30() uint64         { return g.r30 }
func (g *GPR) SetR30(value uint64) { g.r30 = value }

// R31 = RA = return address from subroutine. Not pulled from 'stack'. Change at convenience.
func (g *GPR) R31() uint64         { return g.r31 }
func (g *GPR) SetR31(value uint64) { g.r31 = value }
