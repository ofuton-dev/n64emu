package pif

import (
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

// definition of dummy device
type MockJoyBusDevice struct {
	wantCmd   joybus.CommandType
	wantTxLen types.Byte
	wantRxLen types.Byte
	readDatas []types.Byte
}

func (m *MockJoyBusDevice) Run(cmd joybus.CommandType, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	if cmd != m.wantCmd {
		return joybus.DeviceNotPresent
	}
	if len(txBuf) != int(m.wantTxLen) {
		return joybus.UnableToTransferDatas
	}
	if len(rxBuf) != int(m.wantRxLen) {
		return joybus.UnableToTransferDatas
	}
	// copy datas
	for i := 0; i < len(rxBuf) && i < len(m.readDatas); i++ {
		rxBuf[i] = m.readDatas[i]
	}
	return joybus.Success
}

func TestReadController(t *testing.T) {
	const numOfControllers = 4
	const numOfReadBytes = 4
	tests := []struct {
		name        string
		controllers [numOfControllers]*MockJoyBusDevice
		readDatas   [numOfControllers][numOfReadBytes]types.Byte
		initialRAM  [PIFRAMSize]types.Byte
		wantRAM     [PIFRAMSize]types.Byte
	}{
		{
			name:        "DeviceNotPresent",
			controllers: [numOfControllers]*MockJoyBusDevice{nil, nil, nil, nil},
			readDatas: [numOfControllers][numOfReadBytes]types.Byte{
				{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0},
			},
			initialRAM: [PIFRAMSize]types.Byte{
				0xff, 0x01, 0x04, 0x01, 0xff, 0xff, 0xff, 0xff, // controller0: [dummy, tx:1, rx:4, cmd:ReadButtonValues], [dummy*4]
				0xff, 0x01, 0x04, 0x01, 0xff, 0xff, 0xff, 0xff, // controller1: [dummy, tx:1, rx:4, cmd:ReadButtonValues], [dummy*4]
				0xff, 0x01, 0x04, 0x01, 0xff, 0xff, 0xff, 0xff, // controller2: [dummy, tx:1, rx:4, cmd:ReadButtonValues], [dummy*4]
				0xff, 0x01, 0x04, 0x01, 0xff, 0xff, 0xff, 0xff, // controller3: [dummy, tx:1, rx:4, cmd:ReadButtonValues], [dummy*4]
				0xfe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [end of setup],                                       [dummy*7]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [dummy*8]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [dummy*8]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, //              [dummy*7],                                        [new command]
			},
			wantRAM: [PIFRAMSize]types.Byte{
				0xff, 0x01, 0x84, 0x01, 0xff, 0xff, 0xff, 0xff, // controller0: [dummy, tx:1, rx:(device not present, 4), cmd:ReadButtonValues], [dummy*4]
				0xff, 0x01, 0x84, 0x01, 0xff, 0xff, 0xff, 0xff, // controller1: [dummy, tx:1, rx:(device not present, 4), cmd:ReadButtonValues], [dummy*4]
				0xff, 0x01, 0x84, 0x01, 0xff, 0xff, 0xff, 0xff, // controller2: [dummy, tx:1, rx:(device not present, 4), cmd:ReadButtonValues], [dummy*4]
				0xff, 0x01, 0x84, 0x01, 0xff, 0xff, 0xff, 0xff, // controller3: [dummy, tx:1, rx:(device not present, 4), cmd:ReadButtonValues], [dummy*4]
				0xfe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [end of setup],                                       [dummy*7]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [dummy*8]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [dummy*8]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //              [dummy*7],                                               [idle]
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup structure
			p := PIF{}
			p.ram = tt.initialRAM
			for i := 0; i < numOfControllers; i++ {
				if tt.controllers[i] != nil {
					var c joybus.JoyBus = tt.controllers[i]
					p.controllers[i] = &c
				}
			}
			// Run
			p.Update()
			// Verify
			assert.Equal(t, tt.wantRAM, p.ram)
		})
	}
}
