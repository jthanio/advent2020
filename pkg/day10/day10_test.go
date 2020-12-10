package day10

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day10.txt"

var realInput []int

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadIntInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay10Part1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "small example",
			want:    35,
			wantErr: false,
			args: args{
				input: []int{
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
		},
		{
			name:    "real input",
			want:    1980,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay10Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay10Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay10Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay10Part2(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "real input",
			want:    0,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay10Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay10Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay10Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
