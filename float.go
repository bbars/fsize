package units

import (
	"strconv"
)

func splitFloatUnits(s string) (string, string) {
	i := 0
	sepFound := false
	for _, b := range []byte(s) {
		if b == '.' {
			if !sepFound {
				sepFound = true
			} else {
				break
			}
		} else if b < '0' || b > '9' {
			break
		}

		i++
	}

	return s[0:i], s[i:]
}

func parseFloat(s string) (float64, string, error) {
	s, units := splitFloatUnits(s)

	if f, err := strconv.ParseFloat(s, 64); err != nil {
		return 0, "", err
	} else {
		return f, units, nil
	}
}
