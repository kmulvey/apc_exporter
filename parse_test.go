package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// nolint: gochecknoglobals
var (
	testOutputNormal = `
APC      : 001,036,0872
DATE     : 2024-12-15 08:39:55 -0700  
HOSTNAME : deepthought
VERSION  : 3.14.14 (31 May 2016) redhat
UPSNAME  : deepthought
CABLE    : USB Cable
DRIVER   : USB UPS Driver
UPSMODE  : Stand Alone
STARTTIME: 2024-12-15 07:11:03 -0700  
MODEL    : Back-UPS NS 1500M2 
STATUS   : ONLINE 
LINEV    : 123.0 Volts
LOADPCT  : 29.0 Percent
BCHARGE  : 97.0 Percent
TIMELEFT : 21.0 Minutes
MBATTCHG : 5 Percent
MINTIMEL : 3 Minutes
MAXTIME  : 0 Seconds
SENSE    : Medium
LOTRANS  : 88.0 Volts
HITRANS  : 142.0 Volts
ALARMDEL : No alarm
BATTV    : 27.3 Volts
LASTXFER : No transfers since turnon
NUMXFERS : 0
TONBATT  : 0 Seconds
CUMONBATT: 0 Seconds
XOFFBATT : N/A
SELFTEST : NO
STATFLAG : 0x05000008
SERIALNO : 0B2413L37262  
BATTDATE : 2024-03-28
NOMINV   : 120 Volts
NOMBATTV : 24.0 Volts
NOMPOWER : 900 Watts
FIRMWARE : 957.e5 .D USB FW:e5
END APC  : 2024-12-15 08:40:03 -0700  
`

	testOutputPowerFailure = `
APC      : 001,037,0901
DATE     : 2024-12-16 20:37:41 -0700  
HOSTNAME : deepthought
VERSION  : 3.14.14 (31 May 2016) redhat
UPSNAME  : deepthought
CABLE    : USB Cable
DRIVER   : USB UPS Driver
UPSMODE  : Stand Alone
STARTTIME: 2024-12-16 08:05:02 -0700  
MODEL    : Back-UPS NS 1500M2 
STATUS   : ONBATT 
LINEV    : 0.0 Volts
LOADPCT  : 12.0 Percent
BCHARGE  : 100.0 Percent
TIMELEFT : 56.4 Minutes
MBATTCHG : 5 Percent
MINTIMEL : 3 Minutes
MAXTIME  : 0 Seconds
SENSE    : Medium
LOTRANS  : 88.0 Volts
HITRANS  : 142.0 Volts
ALARMDEL : No alarm
BATTV    : 26.1 Volts
LASTXFER : Low line voltage
NUMXFERS : 1
XONBATT  : 2024-12-16 20:37:38 -0700  
TONBATT  : 6 Seconds
CUMONBATT: 6 Seconds
XOFFBATT : N/A
SELFTEST : NO
STATFLAG : 0x05060010
SERIALNO : 0B2413L37262  
BATTDATE : 2024-03-28
NOMINV   : 120 Volts
NOMBATTV : 24.0 Volts
NOMPOWER : 900 Watts
FIRMWARE : 957.e5 .D USB FW:e5
END APC  : 2024-12-16 20:37:44 -0700  
`
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected BatteryData
	}{
		{
			name:  "Normal",
			input: testOutputNormal,
			expected: BatteryData{
				Status:    "ONLINE",
				Linev:     123.0,
				Loadpct:   29.0,
				Bcharge:   97.0,
				Timeleft:  21 * time.Minute,
				Mbattchg:  5,
				Mintimel:  3 * time.Minute,
				Maxtime:   0 * time.Second,
				Lotrans:   88.0,
				Hitrans:   142.0,
				Battv:     27.3,
				Tonbatt:   0 * time.Second,
				Cumonbatt: 0 * time.Second,
				Nominv:    120,
				Nombattv:  24.0,
				Nompower:  900,
			},
		},
		{
			name:  "PowerFailure",
			input: testOutputPowerFailure,
			expected: BatteryData{
				Status:    "ONBATT",
				Linev:     0.0,
				Loadpct:   12.0,
				Bcharge:   100.0,
				Timeleft:  56 * time.Minute,
				Mbattchg:  5,
				Mintimel:  3 * time.Minute,
				Maxtime:   0 * time.Second,
				Lotrans:   88.0,
				Hitrans:   142.0,
				Battv:     26.1,
				Tonbatt:   6 * time.Second,
				Cumonbatt: 6 * time.Second,
				Nominv:    120,
				Nombattv:  24.0,
				Nompower:  900,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			batteryData, err := parse(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Status, batteryData.Status)
			assert.Equal(t, tt.expected.Linev, batteryData.Linev)       // nolint: testifylint
			assert.Equal(t, tt.expected.Loadpct, batteryData.Loadpct)   // nolint: testifylint
			assert.Equal(t, tt.expected.Bcharge, batteryData.Bcharge)   // nolint: testifylint
			assert.Equal(t, tt.expected.Timeleft, batteryData.Timeleft) // nolint: testifylint
			assert.Equal(t, tt.expected.Mbattchg, batteryData.Mbattchg)
			assert.Equal(t, tt.expected.Mintimel, batteryData.Mintimel)
			assert.Equal(t, tt.expected.Maxtime, batteryData.Maxtime)
			assert.Equal(t, tt.expected.Lotrans, batteryData.Lotrans)
			assert.Equal(t, tt.expected.Hitrans, batteryData.Hitrans)
			assert.Equal(t, tt.expected.Battv, batteryData.Battv)
			assert.Equal(t, tt.expected.Tonbatt, batteryData.Tonbatt)
			assert.Equal(t, tt.expected.Cumonbatt, batteryData.Cumonbatt)
			assert.Equal(t, tt.expected.Nominv, batteryData.Nominv)
			assert.Equal(t, tt.expected.Nombattv, batteryData.Nombattv)
			assert.Equal(t, tt.expected.Nompower, batteryData.Nompower)
		})
	}
}
