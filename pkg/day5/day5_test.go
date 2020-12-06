package day5

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day5.txt"

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay5Part1(t *testing.T) {
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
			name:    "empty",
			want:    0,
			wantErr: false,
			args: args{
				input: nil,
			},
		},
		{
			name:    "basic",
			want:    820,
			wantErr: false,
			args: args{
				input: []string{
					"BFFFBBFRRR",
					"FFFBBBFRRR",
					"BBFFBBFRLL",
				},
			},
		},
		{
			name: "real input",
			want: 816,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay5Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay5Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay5Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay5Part2(t *testing.T) {
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
			name: "real input",
			want: 539,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay5Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay5Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay5Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
