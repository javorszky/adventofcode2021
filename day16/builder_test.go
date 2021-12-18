package day16

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_builder_parseHeader(t *testing.T) {
	type args struct {
		reader io.Reader
	}

	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name:  "parses version, type into ints",
			args:  args{reader: strings.NewReader("001101")},
			want:  1,
			want1: 5,
			want2: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newBuilder()

			got, got1, got2 := b.parseHeader(tt.args.reader)
			assert.Equalf(t, tt.want, got, "parseHeader(%v)", tt.args.reader)
			assert.Equalf(t, tt.want1, got1, "parseHeader(%v)", tt.args.reader)
			assert.Equalf(t, tt.want2, got2, "parseHeader(%v)", tt.args.reader)
		})
	}
}
