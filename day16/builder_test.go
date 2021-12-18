package day16

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_builder_parseHeader(t *testing.T) {
	type args struct {
		reader *strings.Reader
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
		reader *strings.Reader
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
			want: &literal{
				packetVersion: 6,
				packetType:    4,
				value:         2021,
			},
		},
		{
			name: "parses an operator with two literal values in it correctly",
			args: args{reader: strings.NewReader("00111000000000000110111101000101001010010001001000000000")},
			want: &operator{
				packetVersion: 1,
				packetType:    6,
				lengthTypeID:  subPacketLength,
				subPackets: []packet{
					&literal{
						packetVersion: 6,
						packetType:    4,
						value:         10,
					},
					&literal{
						packetVersion: 2,
						packetType:    4,
						value:         20,
					},
				},
			},
		},
		{
			name: "parses an operator with three literal values in it correctly",
			args: args{reader: strings.NewReader("11101110000000001101010000001100100000100011000001100000")},
			want: &operator{
				packetVersion: 7,
				packetType:    3,
				lengthTypeID:  subPacketNumber,
				subPackets: []packet{
					&literal{
						packetVersion: 2,
						packetType:    4,
						value:         1,
					},
					&literal{
						packetVersion: 4,
						packetType:    4,
						value:         2,
					},
					&literal{
						packetVersion: 1,
						packetType:    4,
						value:         3,
					},
				},
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
