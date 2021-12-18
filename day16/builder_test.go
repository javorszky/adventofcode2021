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
			b := newBuilder(tt.args.reader)

			got, got1, got2 := b.parseHeader()
			assert.Equalf(t, tt.want, got, "parseHeader(%v)", tt.args.reader)
			assert.Equalf(t, tt.want1, got1, "parseHeader(%v)", tt.args.reader)
			assert.Equalf(t, tt.want2, got2, "parseHeader(%v)", tt.args.reader)
		})
	}
}

func Test_builder_build(t *testing.T) {
	type args struct {
		reader io.Reader
	}

	tests := []struct {
		name string
		args args
		want packet
	}{
		{
			name: "parses a literal value correctly",
			args: args{
				reader: strings.NewReader("110100101111111000101000"),
			},
			want: literal{
				packetVersion: 6,
				packetType:    4,
				value:         2021,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newBuilder(tt.args.reader)

			assert.Equalf(t, tt.want, b.build(), "build(%v)", tt.args.reader)
		})
	}
}
