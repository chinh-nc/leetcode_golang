package jumpgame

import "testing"

func TestJumpGameCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	want := true

	got := CanJump(nums)
	if got != want {
		t.Fatalf("CanJump() = %v, want %v", got, want)
	}
}

func TestJumpGameCannotJump(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	want := false

	got := CanJump(nums)
	if got != want {
		t.Fatalf("CanJump() = %v, want %v", got, want)
	}
}

func TestJumpGameOutHasOneItem(t *testing.T) {
	nums := []int{0}
	want := true

	got := CanJump(nums)
	if got != want {
		t.Fatalf("CanJump() = %v, want %v", got, want)
	}
}

func TestJumpGameOutOfRange(t *testing.T) {
	nums := []int{2,5,0,0}
	want := true

	got := CanJump(nums)
	if got != want {
		t.Fatalf("CanJump() = %v, want %v", got, want)
	}
}


func TestJumpGameCannotJumpFromFist(t *testing.T) {
	nums := []int{0,1}
	want := false

	got := CanJump(nums)
	if got != want {
		t.Fatalf("CanJump() = %v, want %v", got, want)
	}
}
