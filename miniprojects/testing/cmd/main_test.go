package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	x int
	y int
}

func Test_SumFunc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg args
		answ int 
	}{
		{
			name : "Test n1",
			arg : args{x:1, y:2},
			answ : 3,
		},
		{
			name : "Test n2",
			arg : args{x:-1, y:22},
			answ : 21,
		},
	}

	for _, res := range tests {
		t.Run(res.name, func(t *testing.T) {
			assert.Equal(t, res.answ, SumFunc(res.arg.x, res.arg.y))
		})
	}
}

