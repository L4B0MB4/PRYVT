package test

import (
	"math"
	"testing"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/helper"
)

func TestVersionSplitting64Bit(t *testing.T) {
	var b1 int64 = int64(math.Pow(2, 64)) - 1

	_, _, err := helper.SplitVersion(b1)
	if err == nil {
		t.Error("Should have thrown an error with 64 bit integer")
		t.Fail()
	}
}
func TestVersionSplitting63Bit(t *testing.T) {
	var b1 int64 = int64(math.Pow(2, 63)) - 1

	_, _, err := helper.SplitVersion(b1)
	if err == nil {
		t.Error("Should have thrown an error with 63 bit integer")
		t.Fail()
	}
}

func TestVersionSplitting62Bit(t *testing.T) {

	var b1 int64 = int64(math.Pow(2, 62)) - 1

	_, _, err := helper.SplitVersion(b1)
	if err != nil {
		t.Error("Should be able to handle 62 bit integer")
		t.Fail()
	}
}

func TestVersionSplitting(t *testing.T) {

	var b1 int64 = int64(math.Pow(2, 62)) - 1

	i1, i2, err := helper.SplitVersion(b1)
	if err != nil {
		t.Error("Should be able to handle 62 bit integer")
		t.Fail()
	}
	b_created, err := helper.MergeVersion(i1, i2)

	if err != nil {
		t.Error("Should be able to handle two normal integers")
		t.Fail()
	}
	if b_created != b1 {
		t.Error("big integers should have the same values")
		t.Fail()
	}
}
func TestVersionMergeTwoNegatives(t *testing.T) {

	_, err := helper.MergeVersion(int32(math.Pow(2, 32)-1), int32(math.Pow(2, 32)-1))
	if err == nil {
		t.Error("Should have caused an error")
		t.Fail()
	}
}

func TestVersionMergeOneNegativeOnePositive(t *testing.T) {

	_, err := helper.MergeVersion(int32(math.Pow(2, 32)-1), int32(math.Pow(2, 1)))
	if err == nil {
		t.Error("Should have caused an error")
		t.Fail()
	}
}

func TestVersionMergeOnePositiveOneNegative(t *testing.T) {

	_, err := helper.MergeVersion(int32(math.Pow(2, 1)), int32(math.Pow(2, 32)-1))
	if err == nil {
		t.Error("Should have caused an error")
		t.Fail()
	}
}

func TestVersionMergeTwoPositive(t *testing.T) {

	big, err := helper.MergeVersion(int32(math.Pow(2, 31)-1), int32(math.Pow(2, 31)-1))
	if err != nil {
		t.Error("Should have been handled without error")
		t.Fail()
	}
	if big != 0x3FFF_FFFF_FFFF_FFFF {
		t.Errorf("Should have been equal to %v but was %v", 0x3FFF_FFFF_FFFF_FFFF, big)
		t.Fail()
	}
}

func TestVersionMergeTwoPositiveEx2(t *testing.T) {

	big, err := helper.MergeVersion(int32(math.Pow(2, 30)-1), int32(math.Pow(2, 31)-1))
	if err != nil {
		t.Error("Should have been handled without error")
		t.Fail()
	}
	if big != 0x1FFF_FFFF_FFFF_FFFF {
		t.Errorf("Should have been equal to %v but was %v", 0x1FFF_FFFF_FFFF_FFFF, big)
		t.Fail()
	}
}

func TestVersionMergeTwoPositiveEx3(t *testing.T) {

	big, err := helper.MergeVersion(int32(math.Pow(2, 30)-1), int32(math.Pow(2, 30)-1))
	if err != nil {
		t.Error("Should have been handled without error")
		t.Fail()
	}
	if big != 0x1FFF_FFFF_BFFF_FFFF {
		t.Errorf("Should have been equal to %v but was %v", 0x1FFF_FFFF_BFFF_FFFF, big)
		t.Fail()
	}
}
