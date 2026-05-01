package core

import (
	"reflect"
	"testing"
)

func TestFormatShellStringPowerShell(t *testing.T) {
	program, args := FormatShellString("powershell -NoProfile -Command", "Write-Output hello")

	if program != "powershell" {
		t.Fatalf("expected powershell program, got %q", program)
	}

	expected := []string{"-NoProfile", "-Command", "Write-Output hello"}
	if !reflect.DeepEqual(args, expected) {
		t.Fatalf("expected args %v, got %v", expected, args)
	}
}

func TestFormatShellAddsPowerShellCommandFlag(t *testing.T) {
	program, args := FormatShellString("powershell", "Write-Output hello")

	if program != "powershell" {
		t.Fatalf("expected powershell program, got %q", program)
	}

	expected := []string{"-NoProfile", "-Command", "Write-Output hello"}
	if !reflect.DeepEqual(args, expected) {
		t.Fatalf("expected args %v, got %v", expected, args)
	}
}

func TestFormatShellStringCmd(t *testing.T) {
	program, args := FormatShellString("cmd", "echo hello")

	if program != "cmd" {
		t.Fatalf("expected cmd program, got %q", program)
	}

	expected := []string{"/C", "echo hello"}
	if !reflect.DeepEqual(args, expected) {
		t.Fatalf("expected args %v, got %v", expected, args)
	}
}
