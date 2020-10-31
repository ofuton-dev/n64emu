package reg

// FGR Floating Point General Purpose Register are reserved for floating point operations
type FGR struct {
	f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28, f29, f30, f31 uint64
}

// F0 = register reserved for floating point operations
func (f *FGR) F0() uint64         { return f.f0 }
func (f *FGR) SetF0(value uint64) { f.f0 = value }

// F1 = register reserved for floating point operations
func (f *FGR) F1() uint64         { return f.f1 }
func (f *FGR) SetF1(value uint64) { f.f1 = value }

// F2 = register reserved for floating point operations
func (f *FGR) F2() uint64         { return f.f2 }
func (f *FGR) SetF2(value uint64) { f.f2 = value }

// F3 = register reserved for floating point operations
func (f *FGR) F3() uint64         { return f.f3 }
func (f *FGR) SetF3(value uint64) { f.f3 = value }

// F4 = register reserved for floating point operations
func (f *FGR) F4() uint64         { return f.f4 }
func (f *FGR) SetF4(value uint64) { f.f4 = value }

// F5 = register reserved for floating point operations
func (f *FGR) F5() uint64         { return f.f5 }
func (f *FGR) SetF5(value uint64) { f.f5 = value }

// F6 = register reserved for floating point operations
func (f *FGR) F6() uint64         { return f.f6 }
func (f *FGR) SetF6(value uint64) { f.f6 = value }

// F7 = register reserved for floating point operations
func (f *FGR) F7() uint64         { return f.f7 }
func (f *FGR) SetF7(value uint64) { f.f7 = value }

// F8 = register reserved for floating point operations
func (f *FGR) F8() uint64         { return f.f8 }
func (f *FGR) SetF8(value uint64) { f.f8 = value }

// F9 = register reserved for floating point operations
func (f *FGR) F9() uint64         { return f.f9 }
func (f *FGR) SetF9(value uint64) { f.f9 = value }

// F10 = register reserved for floating point operations
func (f *FGR) F10() uint64         { return f.f10 }
func (f *FGR) SetF10(value uint64) { f.f10 = value }

// F11 = register reserved for floating point operations
func (f *FGR) F11() uint64         { return f.f11 }
func (f *FGR) SetF11(value uint64) { f.f11 = value }

// F12 = register reserved for floating point operations
func (f *FGR) F12() uint64         { return f.f12 }
func (f *FGR) SetF12(value uint64) { f.f12 = value }

// F13 = register reserved for floating point operations
func (f *FGR) F13() uint64         { return f.f13 }
func (f *FGR) SetF13(value uint64) { f.f13 = value }

// F14 = register reserved for floating point operations
func (f *FGR) F14() uint64         { return f.f14 }
func (f *FGR) SetF14(value uint64) { f.f14 = value }

// F15 = register reserved for floating point operations
func (f *FGR) F15() uint64         { return f.f15 }
func (f *FGR) SetF15(value uint64) { f.f15 = value }

// F16 = register reserved for floating point operations
func (f *FGR) F16() uint64         { return f.f16 }
func (f *FGR) SetF16(value uint64) { f.f16 = value }

// F17 = register reserved for floating point operations
func (f *FGR) F17() uint64         { return f.f17 }
func (f *FGR) SetF17(value uint64) { f.f17 = value }

// F18 = register reserved for floating point operations
func (f *FGR) F18() uint64         { return f.f18 }
func (f *FGR) SetF18(value uint64) { f.f18 = value }

// F19 = register reserved for floating point operations
func (f *FGR) F19() uint64         { return f.f19 }
func (f *FGR) SetF19(value uint64) { f.f19 = value }

// F20 = register reserved for floating point operations
func (f *FGR) F20() uint64         { return f.f20 }
func (f *FGR) SetF20(value uint64) { f.f20 = value }

// F21 = register reserved for floating point operations
func (f *FGR) F21() uint64         { return f.f21 }
func (f *FGR) SetF21(value uint64) { f.f21 = value }

// F22 = register reserved for floating point operations
func (f *FGR) F22() uint64         { return f.f22 }
func (f *FGR) SetF22(value uint64) { f.f22 = value }

// F23 = register reserved for floating point operations
func (f *FGR) F23() uint64         { return f.f23 }
func (f *FGR) SetF23(value uint64) { f.f23 = value }

// F24 = register reserved for floating point operations
func (f *FGR) F24() uint64         { return f.f24 }
func (f *FGR) SetF24(value uint64) { f.f24 = value }

// F25 = register reserved for floating point operations
func (f *FGR) F25() uint64         { return f.f25 }
func (f *FGR) SetF25(value uint64) { f.f25 = value }

// F26 = register reserved for floating point operations
func (f *FGR) F26() uint64         { return f.f26 }
func (f *FGR) SetF26(value uint64) { f.f26 = value }

// F27 = register reserved for floating point operations
func (f *FGR) F27() uint64         { return f.f27 }
func (f *FGR) SetF27(value uint64) { f.f27 = value }

// F28 = register reserved for floating point operations
func (f *FGR) F28() uint64         { return f.f28 }
func (f *FGR) SetF28(value uint64) { f.f28 = value }

// F29 = register reserved for floating point operations
func (f *FGR) F29() uint64         { return f.f29 }
func (f *FGR) SetF29(value uint64) { f.f29 = value }

// F30 = register reserved for floating point operations
func (f *FGR) F30() uint64         { return f.f30 }
func (f *FGR) SetF30(value uint64) { f.f30 = value }

// F31 = register reserved for floating point operations
func (f *FGR) F31() uint64         { return f.f31 }
func (f *FGR) SetF31(value uint64) { f.f31 = value }
