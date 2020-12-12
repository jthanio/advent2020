package day11

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day11.txt"

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay11Part1(t *testing.T) {
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
			name:    "small example",
			want:    37,
			wantErr: false,
			args: args{
				input: []string{
					"L.LL.LL.LL",
					"LLLLLLL.LL",
					"L.L.L..L..",
					"LLLL.LL.LL",
					"L.LL.LL.LL",
					"L.LLLLL.LL",
					"..L.L.....",
					"LLLLLLLLLL",
					"L.LLLLLL.L",
					"L.LLLLL.LL",
				},
			},
		},
		{
			name:    "real input",
			want:    2468,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay11Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay11Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay11Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay11Part2(t *testing.T) {
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
			name:    "small example",
			want:    26,
			wantErr: false,
			args: args{
				input: []string{
					"L.LL.LL.LL",
					"LLLLLLL.LL",
					"L.L.L..L..",
					"LLLL.LL.LL",
					"L.LL.LL.LL",
					"L.LLLLL.LL",
					"..L.L.....",
					"LLLLLLLLLL",
					"L.LLLLLL.L",
					"L.LLLLL.LL",
				},
			},
		},
		{
			name:    "real input",
			want:    2214,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay11Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay11Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay11Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
