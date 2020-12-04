package main

import (
	"os"
	"testing"

	"github.com/jthanio/advent2020"
)

var realInput []string

func TestMain(m *testing.M) {
	var err error
	realInput, err = advent2020.LoadStringInputFile(inputFile) // Load the test input data
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestSolveDay4Part1(t *testing.T) {
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
			want: 245,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay4Part1(tt.args.input); got != tt.want {
				t.Errorf("SolveDay4Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveDay4Part2(t *testing.T) {
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
			name: "basic",
			want: 1,
			args: args{
				input: []string{
					"byr:2002 hgt:190cm hcl:#123abc",
					"iyr:2011 eyr:2030 ecl:gry pid:000000000",
				},
			},
		},
		{
			name: "pid too long",
			want: 0,
			args: args{
				input: []string{
					"pid:0874997040 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
					"hcl:#623a2f",
				},
			},
		},
		{
			name: "all wrong",
			want: 0,
			args: args{
				input: []string{
					"eyr:1972 cid:100",
					"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
					"",
					"iyr:2019",
					"hcl:#602927 eyr:1967 hgt:170cm",
					"ecl:grn pid:012533040 byr:1946",
					"",
					"hcl:dab227 iyr:2012",
					"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
					"",
					"hgt:59cm ecl:zzz",
					"eyr:2038 hcl:74454a iyr:2023",
					"pid:3556412378 byr:2007",
				},
			},
		},
		{
			name: "all right",
			want: 4,
			args: args{
				input: []string{
					"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
					"hcl:#623a2f",
					"",
					"eyr:2029 ecl:blu cid:129 byr:1989",
					"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
					"",
					"hcl:#888785",
					"hgt:164cm byr:2001 iyr:2015 cid:88",
					"pid:545766238 ecl:hzl",
					"eyr:2022",
					"",
					"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
				},
			},
		},
		{
			name: "real input",
			want: 184,
			args: args{
				input: realInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDay4Part2(tt.args.input); got != tt.want {
				t.Errorf("SolveDay4Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
