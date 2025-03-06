package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type BatteryData struct {
	Model     string        `description:"UPS model derived from UPS information"`
	Status    string        `description:"UPS status (online, charging, on battery etc)"`
	Linev     float32       `description:"Current input line voltage"`
	Loadpct   float32       `description:"Percentage of UPS load capacity used as estimated by UPS"`
	Bcharge   float32       `description:"Current battery capacity charge percentage"`
	Timeleft  time.Duration `description:"Remaining runtime left on battery as estimated by the UPS"`
	Mbattchg  uint8         `description:"Min battery charge % (BCHARGE) required for system shutdown"`
	Mintimel  time.Duration `description:"Min battery runtime (MINUTES) required for system shutdown"`
	Maxtime   time.Duration `description:"Max battery runtime (TIMEOUT) after which system is shutdown"`
	Lotrans   float32       `description:"Input line voltage below which UPS will switch to battery"`
	Hitrans   float32       `description:"Input line voltage above which UPS will switch to battery"`
	Battv     float32       `description:"Current battery voltage"`
	Tonbatt   time.Duration `description:"Seconds currently on battery"`
	Cumonbatt time.Duration `description:"Cumulative seconds on battery since apcupsd startup"`
	Nominv    uint8         `description:"Nominal input voltage"`
	Nombattv  float32       `description:"Nominal battery voltage"`
	Nompower  uint16        `description:"Nominal power output in watts"`
}

func parse(cmdOutput string) (BatteryData, error) {

	var batteryData BatteryData
	var err error
	var lines = strings.Split(cmdOutput, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var keyVal = strings.SplitN(line, ":", 2)
		var valueString = strings.TrimSpace(keyVal[1])

		switch strings.TrimSpace(keyVal[0]) {
		case "MODEL":
			batteryData.Model = strings.TrimSpace(valueString)

		case "STATUS":
			batteryData.Status = strings.TrimSpace(valueString)

		case "LINEV":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Volts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing volts: %s, err: %w", valueString, err)
			}
			batteryData.Linev = val

		case "LOADPCT":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Percent", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing load: %s, err: %w", valueString, err)
			}
			batteryData.Loadpct = val

		case "BCHARGE":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Percent", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing charge: %s, err: %w", valueString, err)
			}
			batteryData.Bcharge = val

		case "TIMELEFT":
			batteryData.Timeleft, err = parseDuration(valueString)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing timeleft: %s, err: %w", valueString, err)
			}

		case "MBATTCHG":
			var val uint8
			_, err := fmt.Sscanf(valueString, "%d Percent", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing mbattchg: %s, err: %w", valueString, err)
			}
			batteryData.Mbattchg = val

		case "MINTIMEL":
			batteryData.Mintimel, err = parseDuration(valueString)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing mintimel: %s, err: %w", valueString, err)
			}

		case "MAXTIME":
			batteryData.Maxtime, err = parseDuration(valueString)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing maxtime: %s, err: %w", valueString, err)
			}

		case "LOTRANS":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Volts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing lotrans: %s, err: %w", valueString, err)
			}
			batteryData.Lotrans = val

		case "HITRANS":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Volts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing hitrans: %s, err: %w", valueString, err)
			}
			batteryData.Hitrans = val

		case "BATTV":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Volts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing battv: %s, err: %w", valueString, err)
			}
			batteryData.Battv = val

		case "TONBATT":
			batteryData.Tonbatt, err = parseDuration(valueString)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing tonbatt: %s, err: %w", valueString, err)
			}

		case "CUMONBATT":
			batteryData.Cumonbatt, err = parseDuration(valueString)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing cumonbatt: %s, err: %w", valueString, err)
			}

		case "NOMINV":
			var val uint8
			_, err := fmt.Sscanf(valueString, "%d Volts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing nominv: %s, err: %w", valueString, err)
			}
			batteryData.Nominv = val

		case "NOMBATTV":
			var val float32
			_, err := fmt.Sscanf(valueString, "%f Volts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing nombattv: %s, err: %w", valueString, err)
			}
			batteryData.Nombattv = val

		case "NOMPOWER":
			var val uint16
			_, err := fmt.Sscanf(valueString, "%d Watts", &val)
			if err != nil {
				return batteryData, fmt.Errorf("error parsing nompower: %s, err: %w", valueString, err)
			}
			batteryData.Nompower = val

		}
	}

	return batteryData, nil
}

func parseDuration(line string) (time.Duration, error) {

	var val float32
	var unit string
	_, err := fmt.Sscanf(line, "%f %s", &val, &unit)
	if err != nil {
		return 0, fmt.Errorf("error parsing duration: %s, err: %w", line, err)
	}

	switch unit {
	case "Seconds":
		return time.Duration(val) * time.Second, nil
	case "Minutes":
		return time.Duration(val) * time.Minute, nil
	case "Hours":
		return time.Duration(val) * time.Hour, nil
	}

	return 0, errors.New("unable to parse duration: " + line) //nolint:goerr113
}
