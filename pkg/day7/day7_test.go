package day7

import (
	"os"
	"reflect"
	"testing"

	"github.com/jthanio/advent2020/internal"
)

const inputFile = "../../inputs/day7.txt"

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = internal.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay7Part1(t *testing.T) {
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
			want:    4,
			wantErr: false,
			args: args{
				input: []string{
					"light red bags contain 1 bright white bag, 2 muted yellow bags.",
					"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
					"bright white bags contain 1 shiny gold bag.",
					"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
					"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
					"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
					"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
					"faded blue bags contain no other bags.",
					"dotted black bags contain no other bags.",
				},
			},
		},
		{
			name:    "real input",
			want:    161,
			wantErr: false,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveDay7Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay7Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay7Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay7Part2(t *testing.T) {
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
			want:    0,
			wantErr: false,
			args: args{
				input: []string{
					"light red bags contain 1 bright white bag, 2 muted yellow bags.",
					"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
					"bright white bags contain 1 shiny gold bag.",
					"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
					"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
					"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
					"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
					"faded blue bags contain no other bags.",
					"dotted black bags contain no other bags.",
				},
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
			got, err := SolveDay7Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveDay7Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveDay7Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBag(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{
			name:    "single contents",
			wantErr: false,
			want: map[string]int{
				"shiny gold": 1,
			},
			args: args{
				line: "bright white bags contain 1 shiny gold bag.",
			},
		},
		{
			name:    "multiple contents",
			wantErr: false,
			want: map[string]int{
				"shiny gold": 2,
				"faded blue": 9,
			},
			args: args{
				line: "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			},
		},
		{
			name:    "no contents",
			wantErr: false,
			want:    nil,
			args: args{
				line: "dotted black bags contain no other bags.",
			},
		},
		{
			name:    "empty input",
			wantErr: true,
			want:    nil,
			args: args{
				line: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBagContents(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBag() = %v, want %v", got, tt.want)
			}
		})
	}
}
