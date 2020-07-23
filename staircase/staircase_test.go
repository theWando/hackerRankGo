package main

import "testing"

func Test_makeStair(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test Case 0", args: args{6}, want: `     #
    ##
   ###
  ####
 #####
######`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStair(tt.args.n); got != tt.want {
				t.Errorf("makeStair() = %v, want %v", got, tt.want)
			}
		})
	}
}
