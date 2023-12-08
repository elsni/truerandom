package truerandom

import (
	"fmt"
	"math"
	"testing"
)

func TestFetch(t *testing.T) {
	fmt.Println("Testing fetch...")
	result, err := fetch(1, 10, 10)
	if err != nil {
		t.Fatalf("Fetch Error %s", err.Error())
	}
	if len(result) != 10 {
		t.Fatalf("wanted 10 numbers, got %d", len(result))
	}
}

func TestRandInt(t *testing.T) {
	fmt.Println("Testing RandInt...")
	g, err := NewIntGenerator(5, 2, 100)
	if err == nil {
		t.Fatalf("Expected min/max error but got none")
	}
	g, err = NewIntGenerator(1, 6, 0)
	if err == nil {
		t.Fatalf("Expected num error but got none")
	}
	g, err = NewIntGenerator(1, 6, math.MaxInt)
	if err == nil {
		t.Fatalf("Expected num error but got none")
	}
	g, _ = NewIntGenerator(1, 20, 100)
	for i := 1; i < 200; i++ {
		r := g.Rand()
		if r < 1 || r > 20 {
			t.Fatalf("Random number out of range, got %d", r)
		}
	}

}

func TestRandFloat(t *testing.T) {
	fmt.Println("Testing RandFloat...")
	g, err := NewFloatGenerator(5.0, 2.0, 100)
	if err == nil {
		t.Fatalf("Expected min/max error but got none")
	}
	g, err = NewFloatGenerator(1.0, 6.0, 0)
	if err == nil {
		t.Fatalf("Expected num error but got none")
	}
	g, err = NewFloatGenerator(1.0, 6.0, math.MaxInt)
	if err == nil {
		t.Fatalf("Expected num error but got none")
	}
	min := 3.238457623
	max := 12342.123984723987
	g, _ = NewFloatGenerator(min, max, 100)
	for i := 1; i < 200; i++ {
		r := g.Rand()
		if r < min || r > max {
			t.Fatalf("Random number out of range, got %f", r)
		}
	}

}
