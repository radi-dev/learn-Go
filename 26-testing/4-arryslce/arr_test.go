package arryslce

import (
	"testing"
	"testing_test/utils"
)

func TestSum(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	got := Sum(arr)
	want := 10
	utils.AssertCorrectMessage(t, got, want)

}

func TestSumSlice(t *testing.T) {
	slc := []int{1, 2, 3, 4}
	got := SumSlice(slc)
	want := 10
	utils.AssertCorrectMessage(t, got, want)
}

func TestSumAll(t *testing.T) {
	slc1 := []int{1, 2, 3, 4} //10
	slc2 := []int{5, 6, 7, 8} //26
	got := SumAll(slc1, slc2)
	want := 36
	utils.AssertCorrectMessage(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	slc1 := []int{1, 2, 3, 4} //9, excluding 1
	slc2 := []int{5, 6, 7, 8} //21, excluding 5
	got := SumAllTails(slc1, slc2)
	want := 30

	utils.AssertCorrectMessage(t, got, want)

}
