package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_literal(t *testing.T) {
	type fields struct {
		packetVersion int
		packetType    int
		value         int
	}

	type wants struct {
		packetVersion int
		packetType    int
		subPackets    []packet
		allVersions   int
		lengthType    lengthType
	}

	tests := []struct {
		name   string
		fields fields
		wants  wants
	}{
		{
			name: "creates a literal, and tests all the methods",
			fields: fields{
				packetVersion: 2,
				packetType:    3,
				value:         98,
			},
			wants: wants{
				packetVersion: 2,
				packetType:    3,
				subPackets:    nil,
				allVersions:   2,
				lengthType:    unknownLengthType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &literal{
				packetVersion: tt.fields.packetVersion,
				packetType:    tt.fields.packetType,
				value:         tt.fields.value,
			}

			assert.Equalf(t, tt.wants.packetVersion, l.Version(), "Version()")
			assert.Equalf(t, tt.wants.packetType, l.Type(), "Type()")
			assert.Equalf(t, tt.wants.subPackets, l.SubPackets(), "SubPackets()")
			assert.Equalf(t, tt.wants.allVersions, l.AllVersions(), "AllVersions()")
			assert.Equalf(t, tt.wants.lengthType, l.LengthType(), "LengthType()")
		})
	}
}

func Test_operator(t *testing.T) {
	type fields struct {
		packetVersion int
		packetType    int
		subPackets    []packet
		lengthTypeID  lengthType
	}

	type wants struct {
		packetVersion int
		packetType    int
		subPackets    []packet
		allVersions   int
		lengthTypeID  lengthType
	}

	tests := []struct {
		name   string
		fields fields
		wants  wants
	}{
		{
			name: "creates an operator, and tests all the methods",
			fields: fields{
				packetVersion: 2,
				packetType:    3,
				subPackets: []packet{
					&literal{
						packetVersion: 1,
						packetType:    2,
						value:         3,
					},
					&operator{
						packetVersion: 2,
						packetType:    3,
						lengthTypeID:  subPacketNumber,
						subPackets:    nil,
					},
				},
				lengthTypeID: subPacketLength,
			},
			wants: wants{
				packetVersion: 2,
				packetType:    3,
				subPackets: []packet{
					&literal{
						packetVersion: 1,
						packetType:    2,
						value:         3,
					},
					&operator{
						packetVersion: 2,
						packetType:    3,
						lengthTypeID:  subPacketNumber,
						subPackets:    nil,
					},
				},
				allVersions:  5,
				lengthTypeID: subPacketLength,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &operator{
				packetVersion: tt.fields.packetVersion,
				packetType:    tt.fields.packetType,
				lengthTypeID:  tt.fields.lengthTypeID,
				subPackets:    tt.fields.subPackets,
			}

			assert.Equalf(t, tt.wants.packetVersion, l.Version(), "Version()")
			assert.Equalf(t, tt.wants.packetType, l.Type(), "Type()")
			assert.Equalf(t, tt.wants.subPackets, l.SubPackets(), "SubPackets()")
			assert.Equalf(t, tt.wants.allVersions, l.AllVersions(), "AllVersions()")
			assert.Equalf(t, tt.wants.lengthTypeID, l.LengthType(), "LengthType()")
		})
	}
}
