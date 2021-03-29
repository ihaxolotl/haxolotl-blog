// util/slugify_test.go - Test slugify
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package util

import (
	"testing"
)

func TestMarshalSlug(t *testing.T) {
	s := "test->àèâ<-test"
	slug := MarshalSlug(s)

	expected := "test-aea-test"
	if slug != expected {
		t.Fatal("Return string is not slugified as expected", expected, slug)
	}
}

func TestLowerOption(t *testing.T) {
	s := "Test->àèâ<-Test"
	slug := MarshalSlug(s, true)
	expected := "test-aea-test"

	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}

	slug = MarshalSlug(s, false)

	expected = "Test-aea-Test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}

	slug = MarshalSlug(s)

	expected = "Test-aea-Test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
}