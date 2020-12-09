package day9

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day9.txt"

var realInput []int

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadIntInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay9Part1(t *testing.T) {
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
			want:    22477624,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay9Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay9Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay9Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay9Part2(t *testing.T) {
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
			want:    2980044,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay9Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay9Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay9Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
