package units

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilesize_Format(t *testing.T) {
	tests := []struct {
		name   string
		value  Size
		format string
		expect string
	}{
		{
			name:   "s bytes",
			value:  42,
			format: "string: %s",
			expect: "string: 42",
		},
		{
			name:   "s KiB",
			value:  1_500,
			format: "string: %s",
			expect: "string: 1.46K",
		},
		{
			name:   "s GiB",
			value:  1_500_000_000,
			format: "string: %s",
			expect: "string: 1.4G",
		},
		{
			name:   "s TiB",
			value:  2 * TiB,
			format: "string: %s",
			expect: "string: 2T",
		},
		{
			name:   "d bytes",
			value:  42,
			format: "digits: %d",
			expect: "digits: 42",
		},
		{
			name:   "d MiB",
			value:  2 * MiB,
			format: "digits: %d",
			expect: "digits: 2097152",
		},
		{
			name:   "f bytes auto",
			value:  42,
			format: "float: %f",
			expect: "float: 42",
		},
		{
			name:   "f MiB auto",
			value:  1 * 1024 * 1024,
			format: "float: %f",
			expect: "float: 1M",
		},
		{
			name:   "f auto precision",
			value:  1_500_000,
			format: "float: %f",
			expect: "float: 1.43M",
		},
		{
			name:   "f explicit KiB",
			value:  1_500_000,
			format: "float: %+f",
			expect: "float: 1464.84K",
		},
		{
			name:   "f explicit precision",
			value:  1_500_000,
			format: "float: %.4f",
			expect: "float: 1.4305M",
		},
		{
			name:   "f explicit GiB and precision",
			value:  1_500_000,
			format: "float: %+#.6f",
			expect: "float: 0.001397G",
		},
		{
			name:   "f zero GiB",
			value:  0,
			format: "float: %+#f",
			expect: "float: 0G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := fmt.Sprintf(tt.format, tt.value)
			assert.Equal(t, tt.expect, s)
		})
	}
}
