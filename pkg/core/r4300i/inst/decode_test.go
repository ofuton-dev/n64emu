package inst

import (
	"fmt"
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeocdeI(t *testing.T) {
	tests := []struct {
		src  types.Word
		want InstI
	}{
		{
			src: 0x0,
			want: InstI{
				Opcode: 0, Rs: 0, Rt: 0, Immediate: 0,
			},
		},
		{
			src: 0xffff_ffff,
			want: InstI{
				Opcode: 0b11_1111, Rs: 0b1_1111, Rt: 0b1_1111, Immediate: 0xffff,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			got := DecodeI(tt.src)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeocdeJ(t *testing.T) {
	tests := []struct {
		src  types.Word
		want InstJ
	}{
		{
			src: 0x0,
			want: InstJ{
				Opcode: 0, Address: 0,
			},
		},
		{
			src: 0xffff_ffff,
			want: InstJ{
				Opcode: 0b11_1111, Address: 0x03_ffffff,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			got := DecodeJ(tt.src)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeocdeR(t *testing.T) {
	tests := []struct {
		src  types.Word
		want InstR
	}{
		{
			src: 0x0,
			want: InstR{
				Opcode: 0, Rs: 0, Rt: 0, Sa: 0, Funct: 0,
			},
		},
		{
			src: 0xffff_ffff,
			want: InstR{
				Opcode: 0b11_1111, Rs: 0b1_1111, Rt: 0b1_1111, Rd: 0b1_1111, Sa: 0b1_1111, Funct: 0b11_1111,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			got := DecodeR(tt.src)
			assert.Equal(t, tt.want, got)
		})
	}
}
