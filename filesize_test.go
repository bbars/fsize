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
		want   string
	}{
		{
			name:   "s bytes",
			value:  42,
			format: "string: %s",
			want:   "string: 42",
		},
		{
			name:   "s KiB",
			value:  1_500,
			format: "string: %s",
			want:   "string: 1.46K",
		},
		{
			name:   "s GiB",
			value:  1_500_000_000,
			format: "string: %s",
			want:   "string: 1.4G",
		},
		{
			name:   "s TiB",
			value:  2 * TiB,
			format: "string: %s",
			want:   "string: 2T",
		},
		{
			name:   "d bytes",
			value:  42,
			format: "digits: %d",
			want:   "digits: 42",
		},
		{
			name:   "d MiB",
			value:  2 * MiB,
			format: "digits: %d",
			want:   "digits: 2097152",
		},
		{
			name:   "f bytes auto",
			value:  42,
			format: "float: %f",
			want:   "float: 42",
		},
		{
			name:   "f MiB auto",
			value:  1 * 1024 * 1024,
			format: "float: %f",
			want:   "float: 1M",
		},
		{
			name:   "f auto precision",
			value:  1_500_000,
			format: "float: %f",
			want:   "float: 1.43M",
		},
		{
			name:   "f explicit KiB",
			value:  1_500_000,
			format: "float: %+f",
			want:   "float: 1464.84K",
		},
		{
			name:   "f explicit precision",
			value:  1_500_000,
			format: "float: %.4f",
			want:   "float: 1.4305M",
		},
		{
			name:   "f explicit GiB and precision",
			value:  1_500_000,
			format: "float: %+#.6f",
			want:   "float: 0.001397G",
		},
		{
			name:   "f zero GiB",
			value:  0,
			format: "float: %+#f",
			want:   "float: 0G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := fmt.Sprintf(tt.format, tt.value)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestParseSize(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    Size
		wantErr bool
	}{
		{
			name: "no units",
			str:  "42",
			want: 42,
		},
		{
			name: "float KiB",
			str:  "1.465K",
			want: 1_500,
		},
		{
			name: "int TiB",
			str:  "2T",
			want: 2 * TiB,
		},
		{
			name: "zero GiB",
			str:  "0G",
			want: 0,
		},
		{
			name: "wrong case units",
			str:  "0mIb",
			want: 0,
		},
		{
			name:    "unknown units",
			str:     "0ZiB",
			want:    0,
			wantErr: true,
		},
		{
			name:    "malformed float",
			str:     "0.1.2ZiB",
			want:    0,
			wantErr: true,
		},
		// TODO: negative size
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSize(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
