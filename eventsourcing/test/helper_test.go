package test

import (
	"math"
	"testing"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/helper"
)

func TestVersionSplitting(t *testing.T) {

	var b1 int64 = int64(math.Pow(2, 64)) - 1

	_, _, err := helper.SplitVersion(b1)
	if err == nil {
		t.Error("Should have thrown an error with 64 bit integer")
		t.Fail()
	}
	b1 = int64(math.Pow(2, 63)) - 1
	_, _, err = helper.SplitVersion(b1)
	if err == nil {
		t.Error("Should have thrown an error with 63 bit integer")
		t.Fail()
	}
	b1 = int64(math.Pow(2, 62)) - 1

	i1, i2, err := helper.SplitVersion(b1)
	if err != nil {
		t.Error("Should be able to handle 62 bit integer")
		t.Fail()
	}

	t.Logf("%v & %v", i1, i2)
	t.Log("hallo")

}
