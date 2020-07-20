package main

import "testing"

func Test_repeatedString(t *testing.T) {
	type args struct {
		s string
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "Input 0", args: args{
			s: "aba",
			n: 10,
		}, want: 7},
		{name: "Input 1", args: args{
			s: "a",
			n: 1000000000000,
		}, want: 1000000000000},
		{name: "Input 2", args: args{
			s: "abacadadra",
			n: 3,
		}, want: 2},
		{name: "Test Case 4", args: args{
			s: "gfcaaaecbg",
			n: 547602,
		}, want: 164280},
		{name: "Test Case 18", args: args{
			s: "beeaabc",
			n: 711560125001,
		}, want: 203302892858},
		{name: "Test Case 9", args: args{
			s: "epsxyyflvrrrxzvnoenvpegvuonodjoxfwdmcvwctmekpsnamchznsoxaklzjgrqruyzavshfbmuhdwwmpbkwcuomqhiyvuztwvq",
			n: 549382313570,
		}, want: 16481469408},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedString(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("repeatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
