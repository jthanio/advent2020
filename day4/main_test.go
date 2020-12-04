package main

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020"
)

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = advent2020.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay4Part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			want: 0,
			args: args{
				input: nil,
			},
		},
		{
			name: "real input",
			want: 0,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay4Part1(tt.args.input); got != tt.want {
				t.Errorf("SolveDay4Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
