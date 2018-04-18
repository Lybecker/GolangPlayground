package main

import "testing"

func TestAvg(t *testing.T) {
	anders := Person{"Anders"}
	nameChanger(&anders)

	if anders.Name == "Anders" {
		t.Error("The name should have been changed, got ", anders.Name)
	}
}
