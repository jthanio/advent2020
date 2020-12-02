package main

import "testing"

func TestFind2020EntriesAndMultiply(t *testing.T) {
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
			args: args{
				input: nil,
			},
			want: 0,
		},
		{
			name: "only 1 entry",
			args: args{
				input: []int{
					2020,
				},
			},
			want: 0,
		},
		{
			name: "add 0",
			args: args{
				input: []int{
					0,
					2020,
				},
			},
			want: 0,
		},
		{
			name: "add 1",
			args: args{
				input: []int{
					1,
					2019,
				},
			},
			want: 2019,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find2020EntriesAndMultiply(tt.args.input); got != tt.want {
				t.Errorf("Find2020EntriesAndMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
