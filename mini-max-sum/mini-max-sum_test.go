package main

import "testing"

func Test_miniMaxSumCalculator(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		min  int
		max  int
	}{
		{name: "Test Case 0", args: args{[]int32{1, 2, 3, 4, 5}}, min: 10, max: 14},
		{name: "Test Case 1", args: args{[]int32{7, 69, 2, 221, 8974}}, min: 299, max: 9271},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, max := miniMaxSumCalculator(tt.args.arr)
			if min != tt.min {
				t.Errorf("miniMaxSumCalculator() min = %v, min %v", min, tt.min)
			}
			if max != tt.max {
				t.Errorf("miniMaxSumCalculator() max = %v, max %v", max, tt.max)
			}
		})
	}
}
