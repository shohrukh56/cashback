package main

import (
	"testing"
)

func Test_cashback(t *testing.T) {
	tests := []struct {
		name string
		amount int
		want int
	}{
		//TODO: Add test cases.
		{"dfdf" ,5000,250  },
		{"no cashback", 1000, 0},
		{"cahsback om=n bound", 3000, 150},
	}
	for _, test := range tests {
		got :=cashback(test.amount)
		if got != test.want{
			t.Error("for cashback", test.name, "got", got, "want:", test.want)
		}
	}


}