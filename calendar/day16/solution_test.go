package main

import (
	"testing"
)

func TestParseField(t *testing.T){
	line := "arrival track: 29-467 or 477-950"
	got := ParseField(line)
	expected := Field{"arrival track", []Range{Range{29,467}, Range{477,950}}}

	if(got.Name != expected.Name){
		t.Errorf("expected: %v got %v", expected, got)
	}

	for i, _ := range got.Ranges {
		if(got.Ranges[i] != expected.Ranges[i]){
			t.Errorf("expected: %v got %v", expected, got)
		}
	}
}
