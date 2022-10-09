package main

import (
	"strings"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	sample := "port = 8080"
	r := strings.NewReader(sample)

	config, err := LoadConfig(r)
	if err != nil {
		t.Fatalf("failed to load config: \n%s\n", sample)
	}

	if config.Port != 8080 {
		t.Fatalf("found unexpected value")
	}
}

func TestConfigLoadError(t *testing.T) {
	sample := "potr = 8080"
	r := strings.NewReader(sample)

	_, err := LoadConfig(r)
	if err != nil {
		return
	}

	t.Fatalf("Not found expected error")
}
