package main

import "testing"

func Test_plusMinusWork(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
		want2 float64
	}{
		{
			name:  "Test case 0",
			args:  args{[]int32{-4, 3, -9, 0, 4, 1}},
			want:  0.500000,
			want1: 0.333333,
			want2: 0.166667,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := plusMinusWork(tt.args.arr)
			if got != tt.want {
				t.Errorf("plusMinusWork() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("plusMinusWork() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("plusMinusWork() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
