package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testOutputNormal = `
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

var testOutputPowerFailure = `
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

func TestParse(t *testing.T) {
	t.Parallel()

	batteryData, err := parse(testOutputNormal)
	assert.NoError(t, err)

	assert.Equal(t, "ONLINE", batteryData.Status)
	assert.Equal(t, float32(123.0), batteryData.Linev)
	assert.Equal(t, float32(29.0), batteryData.Loadpct)
	assert.Equal(t, float32(97.0), batteryData.Bcharge)
	assert.Equal(t, 21*time.Minute, batteryData.Timeleft)
	assert.Equal(t, uint8(5), batteryData.Mbattchg)
	assert.Equal(t, 3*time.Minute, batteryData.Mintimel)
	assert.Equal(t, 0*time.Second, batteryData.Maxtime)
	assert.Equal(t, float32(88.0), batteryData.Lotrans)
	assert.Equal(t, float32(142.0), batteryData.Hitrans)
	assert.Equal(t, float32(27.3), batteryData.Battv)
	assert.Equal(t, 0*time.Second, batteryData.Tonbatt)
	assert.Equal(t, 0*time.Second, batteryData.Cumonbatt)
	assert.Equal(t, uint8(120), batteryData.Nominv)
	assert.Equal(t, float32(24.0), batteryData.Nombattv)
	assert.Equal(t, uint16(900), batteryData.Nompower)
}

func TestParseFailure(t *testing.T) {
	t.Parallel()

	batteryData, err := parse(testOutputPowerFailure)
	assert.NoError(t, err)

	assert.Equal(t, "ONBATT", batteryData.Status)
	assert.Equal(t, float32(0.0), batteryData.Linev)
	assert.Equal(t, float32(12.0), batteryData.Loadpct)
	assert.Equal(t, float32(100.0), batteryData.Bcharge)
	assert.Equal(t, 56*time.Minute, batteryData.Timeleft)
	assert.Equal(t, uint8(5), batteryData.Mbattchg)
	assert.Equal(t, 3*time.Minute, batteryData.Mintimel)
	assert.Equal(t, 0*time.Second, batteryData.Maxtime)
	assert.Equal(t, float32(88.0), batteryData.Lotrans)
	assert.Equal(t, float32(142.0), batteryData.Hitrans)
	assert.Equal(t, float32(26.1), batteryData.Battv)
	assert.Equal(t, 6*time.Second, batteryData.Tonbatt)
	assert.Equal(t, 6*time.Second, batteryData.Cumonbatt)
	assert.Equal(t, uint8(120), batteryData.Nominv)
	assert.Equal(t, float32(24.0), batteryData.Nombattv)
	assert.Equal(t, uint16(900), batteryData.Nompower)
}
