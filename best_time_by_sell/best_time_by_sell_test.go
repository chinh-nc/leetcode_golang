package besttimebysell

import "testing"

func TestMaxProfitCorrect(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	want := 5
	result := MaxProfit(prices)
	if result != want {
		t.Fatalf("MaxProfit() = %v, want %v", result, want)
	}
}

func TestMaxProfitReturnZero(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	want := 0
	result := MaxProfit(prices)
	if result != want {
		t.Fatalf("MaxProfit() = %v, want %v", result, want)
	}
}
