package main

import "testing"

func TestSum(t *testing.T)  {
	for i := 0 ; i < 10 ; i++ {
		if Sum(i , i +1) != 2*i +1 {
			t.Errorf("it's wierd to got error")
		}
	}
}

