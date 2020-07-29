package main

import "testing"

func Test_angryProfessor(t *testing.T) {
	type args struct {
		k int32
		a []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test Case 0a", args: args{
			k: 3,
			a: []int32{-1, -3, 4, 2},
		}, want: "YES"},
		{name: "Test Case 0b", args: args{
			k: 2,
			a: []int32{0, -1, 2, 1},
		}, want: "NO"},
		{name: "Test Case 1a", args: args{
			k: 4,
			a: []int32{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67},
		}, want: "NO"},
		{name: "Test Case 1b", args: args{
			k: 10,
			a: []int32{23, -35, -2, 58, -67, -56, -42, -73, -19, 37},
		}, want: "YES"},
		{name: "Test Case 1c", args: args{
			k: 9,
			a: []int32{13, 91, 56, -62, 96, -5, -84, -36, -46, -13},
		}, want: "YES"},
		{name: "Test Case 1d", args: args{
			k: 8,
			a: []int32{45, 67, 64, -50, -8, 78, 84, -51, 99, 60},
		}, want: "YES"},
		{name: "Test Case 1e", args: args{
			k: 7,
			a: []int32{26, 94, -95, 34, 67, -97, 17, 52, 1, 86},
		}, want: "YES"},
		{name: "Test Case 1f", args: args{
			k: 2,
			a: []int32{19, 29, -17, -71, -75, -27, -56, -53, 65, 83},
		}, want: "NO"},
		{name: "Test Case 1g", args: args{
			k: 10,
			a: []int32{-32, -3, -70, 8, -40, -96, -18, -46, -21, -79},
		}, want: "YES"},
		{name: "Test Case 1h", args: args{
			k: 9,
			a: []int32{-50, 0, 64, 14, -56, -91, -65, -36, 51, -28},
		}, want: "YES"},
		{name: "Test Case 1i", args: args{
			k: 6,
			a: []int32{-58, -29, -35, -18, 43, -28, -76, 43, -13, 6},
		}, want: "NO"},
		{name: "Test Case 1j", args: args{
			k: 1,
			a: []int32{88, -17, -96, 43, 83, 99, 25, 90, -39, 86},
		}, want: "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angryProfessor(tt.args.k, tt.args.a); got != tt.want {
				t.Errorf("angryProfessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
