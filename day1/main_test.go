package main

import (
	"os"
	"testing"
)

var realInput []int

func TestMain(m *testing.M) {
	var err error
	realInput, err = loadInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay1Part1(t *testing.T) {
	type args struct {
		input []int
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
			name: "only 1 entry",
			want: 0,
			args: args{
				input: []int{
					2020,
				},
			},
		},
		{
			name: "add 0",
			want: 0,
			args: args{
				input: []int{
					0,
					2020,
				},
			},
		},
		{
			name: "add 1",
			want: 2019,
			args: args{
				input: []int{
					1,
					2019,
				},
			},
		},
		{
			name: "real input",
			want: 145875,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay1Part1(tt.args.input); got != tt.want {
				t.Errorf("SolveDay1Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay1Part2(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "real input",
			want: 69596112,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay1Part2(tt.args.input); got != tt.want {
				t.Errorf("SolveDay1Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
