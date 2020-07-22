package main

import "testing"

func Test_simpleArraySum(t *testing.T) {
	type args struct {
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "Test Case 0", args: args{ar: []int32{1, 2, 3, 4, 10, 11}}, want: 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simpleArraySum(tt.args.ar); got != tt.want {
				t.Errorf("simpleArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
