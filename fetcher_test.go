package main

import (
	"testing"

	"github.com/hkairi/adjust/fetcher"
)

func TestNewFetcher(t *testing.T) {
	f := fetcher.New(0, []string{})

	if f.Limit != 0 {
		t.Errorf("default value should be 0 but got %d", f.Limit)
	}

	if len(f.Urls) != 0 {
		t.Errorf("[]urls should be empty. But this slice contains %d item(s)", len(f.Urls))
	}
}

func TestHashingText(t *testing.T) {
	f := fetcher.New(0, []string{})
	h := f.HashText("adjust.com")

	if h != "26fe2876c7a97831ace94b4748e1ece4" {
		t.Fail()
	}
}
