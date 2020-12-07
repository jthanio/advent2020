package day6

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day6.txt"

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay6Part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small example",
			want: 11,
			args: args{
				input: []string{
					"abc",
					"",
					"a",
					"b",
					"c",
					"",
					"ab",
					"ac",
					"",
					"a",
					"a",
					"a",
					"a",
					"",
					"b",
				},
			},
		},
		{
			name: "real input",
			want: 7283,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay6Part1(tt.args.input); got != tt.want {
				t.Errorf("SolveDay6Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay6Part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small example",
			want: 6,
			args: args{
				input: []string{
					"abc",
					"",
					"a",
					"b",
					"c",
					"",
					"ab",
					"ac",
					"",
					"a",
					"a",
					"a",
					"a",
					"",
					"b",
				},
			},
		},
		{
			name: "real input",
			want: 3520,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay6Part2(tt.args.input); got != tt.want {
				t.Errorf("SolveDay6Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
