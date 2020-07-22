package main

import "testing"

func Test_solveMeFirst(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "Test case 0", args: args{
			a: 2,
			b: 3,
		}, want: 5},
		{name: "Test case 1", args: args{
			a: 100,
			b: 1000,
		}, want: 1100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveMeFirst(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveMeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
