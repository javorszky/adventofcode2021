package day16

import (
	"fmt"
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
		name    string
		args    args
		want    int
		want1   int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "parses version, type into ints",
			args:    args{reader: strings.NewReader("001101")},
			want:    1,
			want1:   5,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newBuilder()

			got, got1, err := b.parseHeader(tt.args.reader)
			if !tt.wantErr(t, err, fmt.Sprintf("parseHeader(%v)", tt.args.reader)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseHeader(%v)", tt.args.reader)
			assert.Equalf(t, tt.want1, got1, "parseHeader(%v)", tt.args.reader)
		})
	}
}
