package helpers

import "testing"

func TestParseIDFromPath(t *testing.T) {
	p1 := "/api/v1/tweet/2"
	got, _ := ParseIDFromPath(p1, 4)
	want := 2
	if got != want {
		t.Fatalf("want %d got %d\n", want, got)
	}

	p2 := "/api/v1/tweet/2"
	got, _ = ParseIDFromPath(p2, -1)
	want = 2
	if got != want {
		t.Fatalf("want %d got %d\n", want, got)
	}

	p3 := "api/v1/users/17281/feed"
	got, _ = ParseIDFromPath(p3, 4)
	want = 17281
	if got != want {
		t.Fatalf("want %d got %d\n", want, got)
	}

	p4 := "api/v1/tweet"
	got, err := ParseIDFromPath(p4, 3)
	if err == nil {
		t.Fatalf("want %d got %d\n", want, got)
	}
	if got != 0 {
		t.Fatalf("want %d got %d\n", want, got)
	}

	p5 := "/api/v1/users/717/feed/"
	got, _ = ParseIDFromPath(p5, 4)
	want = 717
	if got != want {
		t.Fatalf("want %d got %d\n", want, got)
	}

	p6 := "/api/v1/tweet/2179513/"
	got, _ = ParseIDFromPath(p6, 4)
	want = 2179513
	if got != want {
		t.Fatalf("want %d got %d\n", want, got)
	}
}
