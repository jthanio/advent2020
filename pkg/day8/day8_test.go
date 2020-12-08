package day8

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day8.txt"

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay8Part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "small input",
			want:    5,
			wantErr: false,
			args: args{

				input: []string{
					"nop +0",
					"acc +1",
					"jmp +4",
					"acc +3",
					"jmp -3",
					"acc -99",
					"acc +1",
					"jmp -4",
					"acc +6",
				},
			},
		},
		{
			name:    "real input",
			want:    1818,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay8Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay8Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay8Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay8Part2(t *testing.T) {
	type args struct {
		input []string
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
			got, err := SolveDay8Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay8Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay8Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
