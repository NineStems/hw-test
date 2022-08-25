package main

import (
	"os"
	"testing"
)

func TestRunCmd(t *testing.T) {
	code := RunCmd([]string{"ls"}, nil)
	if code != 0 {
		t.Errorf("RunCmd() = %v, want %v", code, 0)
	}
}

func TestBadRunCmd(t *testing.T) {
	code := RunCmd([]string{""}, nil)
	if code != -1 {
		t.Errorf("RunCmd() = %v, want %v", code, -1)
	}
}

func TestSetEnvRunCmd(t *testing.T) {
	os.Setenv("EMPTY", "empty")
	os.Setenv("BAR", "notBar")
	empty := os.Getenv("EMPTY")
	bar := os.Getenv("BAR")
	if empty != "empty" {
		t.Errorf("empty = %v, want %v", empty, "empty")
	}
	if bar != "notBar" {
		t.Errorf("bar = %v, want %v", bar, "notBar")
	}

	envs := Environment{
		"EMPTY": EnvValue{Value: "", NeedRemove: true},
		"BAR":   EnvValue{Value: "bar"},
	}
	code := RunCmd([]string{"ls"}, envs)
	if code != 0 {
		t.Errorf("RunCmd() = %v, want %v", code, 0)
	}
	empty = os.Getenv("EMPTY")
	bar = os.Getenv("BAR")
	if empty != "" {
		t.Errorf("empty = %v, want %v", empty, "")
	}
	if bar != "bar" {
		t.Errorf("bar = %v, want %v", bar, "bar")
	}
}
