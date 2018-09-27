package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	s := "foo\nbar\nbaz\nbar\n"
	r := strings.NewReader(s)
	var w bytes.Buffer

	if run(r, &w) != 0 {
		t.Fatal("run exit with failure")
	}
	got := w.String()
	want := "foo\nbar\nbaz\n"
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}
