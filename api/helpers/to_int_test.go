package helpers

import "testing"

func TestStrToInt(t *testing.T) {
	a1 := "16"
	want := 16
	got, _ := StrToInt(a1)
	if want != got {
		t.Fatalf("want %d, got %d\n", want, got)
	}
	a2 := ""
	_, err := StrToInt(a2)
	if err == nil {
		t.Fatalf("should have returned an error")
	}
}
