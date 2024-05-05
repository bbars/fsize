package units

import (
	"fmt"
	"strconv"
)

const (
	Bytes Size = 1
	KiB        = 1024
	MiB        = 1024 * 1024
	GiB        = 1024 * 1024 * 1024
	TiB        = 1024 * 1024 * 1024 * 1024
	PiB        = 1024 * 1024 * 1024 * 1024 * 1024
)

type Size int64

var _ fmt.Stringer = Size(0)

func (s Size) String() string {
	return s.format(0, 2, true)
}

func (s Size) format(units Size, prec int, trim bool) string {
	var str string
	switch {
	case units > 0:
		str = strconv.FormatFloat(float64(s)/float64(units), 'f', prec, 64)
	case s >= PiB:
		str = strconv.FormatFloat(float64(s)/float64(PiB), 'f', prec, 64)
		units = PiB
	case s >= TiB:
		str = strconv.FormatFloat(float64(s)/float64(TiB), 'f', prec, 64)
		units = TiB
	case s >= GiB:
		str = strconv.FormatFloat(float64(s)/float64(GiB), 'f', prec, 64)
		units = GiB
	case s >= MiB:
		str = strconv.FormatFloat(float64(s)/float64(MiB), 'f', prec, 64)
		units = MiB
	case s >= KiB:
		str = strconv.FormatFloat(float64(s)/float64(KiB), 'f', prec, 64)
		units = KiB
	default:
		return strconv.FormatInt(int64(s), 10)
	}

	if trim {
		var i int
		for i = len(str) - 1; i > 0; i-- {
			if str[i] != '0' {
				break
			}
		}

		if str[i] == '.' {
			i--
		}
		str = str[:i+1]
	}

	return str + units.units()
}

var _ fmt.Formatter = Size(0)

// Format formats file size.
//
//	%s — auto units, precision 2 (the same as String)
//	%d — always bytes
//	%f — floating point; supports + and # flags: + is KiB, # is MiB, +# is GiB (auto units by default)
func (s Size) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		_, _ = f.Write([]byte(s.String()))
	case 'd':
		_, _ = f.Write([]byte(strconv.FormatInt(int64(s), 10)))
	case 'f':
		prec, ok := f.Precision()
		if !ok {
			prec = 2
		}
		_, _ = f.Write([]byte(s.format(fmtUnits(f), prec, !ok)))
	case 'v':
		prec, ok := f.Precision()
		if !ok {
			prec = -1
		}
		_, _ = f.Write([]byte(s.format(0, prec, false)))
	default:
		_, _ = f.Write([]byte(s.String()))
	}
}

func fmtUnits(f fmt.State) Size {
	m := f.Flag('#')
	k := f.Flag('+')

	switch {
	case m && k:
		return GiB
	case m:
		return MiB
	case k:
		return KiB
	default:
		return 0
	}
}

func (s Size) units() string {
	switch s {
	case PiB:
		return "P"
	case TiB:
		return "T"
	case GiB:
		return "G"
	case MiB:
		return "M"
	case KiB:
		return "K"
	case Bytes:
		return ""
	default:
		return "?"
	}
}
