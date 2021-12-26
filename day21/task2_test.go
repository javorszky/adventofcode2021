package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniverse_String(t *testing.T) {
	type fields struct {
		p1Step  int
		p1Score int
		p2Step  int
		p2Score int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "correctly encodes a universe",
			fields: fields{
				p1Step:  1,
				p1Score: 2,
				p2Step:  3,
				p2Score: 4,
			},
			want: "1/2/3/4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Universe{
				p1Step:  tt.fields.p1Step,
				p1Score: tt.fields.p1Score,
				p2Step:  tt.fields.p2Step,
				p2Score: tt.fields.p2Score,
			}
			assert.Equalf(t, tt.want, u.String(), "String()")
		})
	}
}

func TestUniverse_Marshal(t *testing.T) {
	type fields struct {
		p1Step  int
		p1Score int
		p2Step  int
		p2Score int
	}

	type args struct {
		s string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "correctly unmarshals a text into a universe",
			fields: fields{
				p1Step:  1,
				p1Score: 2,
				p2Step:  3,
				p2Score: 4,
			},
			args: args{s: "1/2/3/4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Universe{}
			u.Marshal(tt.args.s)

			want := Universe{
				p1Step:  tt.fields.p1Step,
				p1Score: tt.fields.p1Score,
				p2Step:  tt.fields.p2Step,
				p2Score: tt.fields.p2Score,
			}

			assert.Equalf(t, want, u, "marshal to universe failed")
		})
	}
}
