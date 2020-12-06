package day3

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day3.txt"

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay3Part1(t *testing.T) {
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
			want: 232,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay3Part1(tt.args.input); got != tt.want {
				t.Errorf("SolveDay3Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay3Part2(t *testing.T) {
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
			want: 3952291680,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay3Part2(tt.args.input); got != tt.want {
				t.Errorf("SolveDay3Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
