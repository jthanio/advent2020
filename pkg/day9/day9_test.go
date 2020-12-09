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

var out int // A global variable to avoid optimizations

// Before optimization
// goos: linux
// goarch: amd64
// pkg: github.com/jthanio/advent2020/pkg/day9
// BenchmarkSolveDay9Part2-16    	      27	  41431141 ns/op	 2954917 B/op	    2205 allocs/op
// PASS
// ok  	github.com/jthanio/advent2020/pkg/day9	1.165s

// After optimization
// goos: linux
// goarch: amd64
// pkg: github.com/jthanio/advent2020/pkg/day9
// BenchmarkSolveDay9Part2-16    	     146	   8203799 ns/op	 2954811 B/op	    2202 allocs/op
// PASS
// ok  	github.com/jthanio/advent2020/pkg/day9	2.030s

func BenchmarkSolveDay9Part2(b *testing.B) {
	var x int
	for n := 0; n < b.N; n++ {
		var err error
		x, err = SolveDay9Part2(realInput)
		if err != nil {
			b.Fatal(err)
		}
	}
	out = x // Assign to the global variable to avoid compiler optimization
}
