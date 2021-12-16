package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newPathNode(t *testing.T) {
	type args struct {
		col  int
		row  int
		cost int
	}

	tests := []struct {
		name            string
		args            args
		wantDesignation int
		wantCost        int
	}{
		{
			name: "adds a node, does all the things to test",
			args: args{
				col:  14,
				row:  65,
				cost: 9,
			},
			wantDesignation: 140065,
			wantCost:        9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newPathNode(tt.args.col, tt.args.row, tt.args.cost)

			assert.Equal(t, tt.wantDesignation, n.Designation())
			assert.Equal(t, tt.wantCost, n.Cost())
		})
	}
}

func Test_pathNode_SetCost(t *testing.T) {
	type fields struct {
		col, row, cost int
	}

	type args struct {
		c int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "sets up a node, and updates cost with a lower one",
			fields: fields{
				col:  14,
				row:  53,
				cost: 837,
			},
			args: args{c: 44},
			want: 44,
		},
		{
			name: "sets up a node, and update gets ignored silently with larger one",
			fields: fields{
				col:  14,
				row:  53,
				cost: 837,
			},
			args: args{c: 988},
			want: 837,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newPathNode(tt.fields.col, tt.fields.row, tt.fields.cost)
			assert.Equalf(t, tt.want, n.SetCost(tt.args.c).Cost(), "SetCost(%v)", tt.args.c)
		})
	}
}

func Test_newPathFinder(t *testing.T) {
	tests := []struct {
		name     string
		wantNext *pathNode
	}{
		{
			name:     "creates a new pathfinder struct and returns the pointer",
			wantNext: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantNext, newPathFinder().Next(), "newPathFinder()")
		})
	}
}

func Test_pathFinder_AddElement(t *testing.T) {
	type args struct {
		col  int
		row  int
		cost int
	}

	type states struct {
		visited     map[int]struct{}
		inQueue     map[int]int
		queueLength int
	}

	type want struct {
		col    int
		row    int
		cost   int
		before states
		after  states
	}

	tests := []struct {
		name string
		args []args
		want
	}{
		{
			name: "adds two elements, returns the smaller one",
			args: []args{
				{
					col:  11,
					row:  12,
					cost: 9,
				},
				{
					col:  14,
					row:  33,
					cost: 1,
				},
			},
			want: want{
				before: states{
					visited: map[int]struct{}{},
					inQueue: map[int]int{
						110012: 9,
						140033: 1,
					},
					queueLength: 2,
				},
				col:  14,
				row:  33,
				cost: 1,
				after: states{
					visited: map[int]struct{}{
						140033: {},
					},
					inQueue: map[int]int{
						110012: 9,
					},
					queueLength: 1,
				},
			},
		},
		{
			name: "adds the same element, updates cost",
			args: []args{
				{
					col:  11,
					row:  12,
					cost: 9,
				},
				{
					col:  11,
					row:  12,
					cost: 3,
				},
			},
			want: want{
				before: states{
					visited: map[int]struct{}{},
					inQueue: map[int]int{
						110012: 3,
					},
					queueLength: 1,
				},
				col:  11,
				row:  12,
				cost: 3,
				after: states{
					visited: map[int]struct{}{
						110012: {},
					},
					inQueue:     map[int]int{},
					queueLength: 0,
				},
			},
		},
		{
			name: "adds the same element, ignores updating cost",
			args: []args{
				{
					col:  11,
					row:  12,
					cost: 9,
				},
				{
					col:  11,
					row:  12,
					cost: 83,
				},
			},
			want: want{
				before: states{
					visited: map[int]struct{}{},
					inQueue: map[int]int{
						110012: 9,
					},
					queueLength: 1,
				},
				col:  11,
				row:  12,
				cost: 9,
				after: states{
					visited: map[int]struct{}{
						110012: {},
					},
					inQueue:     map[int]int{},
					queueLength: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newPathFinder()

			for _, a := range tt.args {
				p.AddElement(a.col, a.row, a.cost)
			}

			assert.Equal(t, tt.want.before.queueLength, p.Left())
			assert.Equal(t, len(tt.want.before.inQueue), len(p.elementsInQueue))

			els := make(map[int]int)
			for k, v := range p.elementsInQueue {
				els[k] = v.Cost()
			}

			assert.Equal(t, tt.want.before.inQueue, els)
			assert.Equal(t, tt.want.before.visited, p.visitedElements)

			pn := p.Next()
			assert.Equal(t, tt.want.cost, pn.Cost())
			assert.Equal(t, tt.want.col*10000+tt.want.row, pn.Designation())

			assert.Equal(t, len(tt.want.after.inQueue), len(p.elementsInQueue))

			elsAfter := make(map[int]int)
			for k, v := range p.elementsInQueue {
				elsAfter[k] = v.Cost()
			}

			assert.Equal(t, tt.want.after.inQueue, elsAfter)
			assert.Equal(t, tt.want.after.visited, p.visitedElements)
			assert.Equal(t, tt.want.after.queueLength, p.Left())
		})
	}
}

func Test_getNeighbours(t *testing.T) {
	field := map[int]map[int]int{
		0: {
			0: 1,
			1: 4,
			2: 3,
		},
		1: {
			0: 3,
			1: 1,
			2: 2,
		},
		2: {
			0: 3,
			1: 5,
			2: 3,
		},
	}

	type args struct {
		n     *pathNode
		field map[int]map[int]int
	}

	tests := []struct {
		name string
		args args
		want []neighbour
	}{
		{
			name: "gets neighbours in top left corner",
			args: args{
				n:     newPathNode(0, 0, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  1,
					row:  0,
					cost: 4,
				},
				{
					col:  0,
					row:  1,
					cost: 3,
				},
			},
		},
		{
			name: "gets neighbours in top right corner",
			args: args{
				n:     newPathNode(2, 0, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  1,
					row:  0,
					cost: 4,
				},
				{
					col:  2,
					row:  1,
					cost: 2,
				},
			},
		},
		{
			name: "gets neighbours in bottom left corner",
			args: args{
				n:     newPathNode(0, 2, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  1,
					row:  2,
					cost: 5,
				},
				{
					col:  0,
					row:  1,
					cost: 3,
				},
			},
		},
		{
			name: "gets neighbours in bottom right corner",
			args: args{
				n:     newPathNode(2, 2, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  1,
					row:  2,
					cost: 5,
				},
				{
					col:  2,
					row:  1,
					cost: 2,
				},
			},
		},
		{
			name: "gets neighbours in top middle",
			args: args{
				n:     newPathNode(1, 0, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  0,
					row:  0,
					cost: 1,
				},
				{
					col:  2,
					row:  0,
					cost: 3,
				},
				{
					col:  1,
					row:  1,
					cost: 1,
				},
			},
		},
		{
			name: "gets neighbours in right middle",
			args: args{
				n:     newPathNode(2, 1, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  2,
					row:  0,
					cost: 3,
				},
				{
					col:  1,
					row:  1,
					cost: 1,
				},
				{
					col:  2,
					row:  2,
					cost: 3,
				},
			},
		},
		{
			name: "gets neighbours in bottom middle",
			args: args{
				n:     newPathNode(1, 2, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  2,
					row:  2,
					cost: 3,
				},
				{
					col:  0,
					row:  2,
					cost: 3,
				},
				{
					col:  1,
					row:  1,
					cost: 1,
				},
			},
		},
		{
			name: "gets neighbours in left middle",
			args: args{
				n:     newPathNode(0, 1, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  0,
					row:  0,
					cost: 1,
				},
				{
					col:  0,
					row:  2,
					cost: 3,
				},
				{
					col:  1,
					row:  1,
					cost: 1,
				},
			},
		},
		{
			name: "gets neighbours in middle",
			args: args{
				n:     newPathNode(1, 1, 0),
				field: field,
			},
			want: []neighbour{
				{
					col:  1,
					row:  0,
					cost: 4,
				},
				{
					col:  0,
					row:  1,
					cost: 3,
				},
				{
					col:  1,
					row:  2,
					cost: 5,
				},
				{
					col:  2,
					row:  1,
					cost: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, getNeighbours(tt.args.n, tt.args.field),
				"getNeighbours(%v, %v)", tt.args.n, tt.args.field)
		})
	}
}
