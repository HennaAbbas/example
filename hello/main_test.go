package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Helper function to build the binary and run the test cases
func runTest(args []string, expected string, t *testing.T) {
	// Build the binary
	cmd := exec.Command("go", "build", "-o", "testbinary")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("testbinary")

	// Execute the binary with given arguments
	cmd = exec.Command("./testbinary", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to execute binary: %v", err)
	}

	// Verify the output
	if got := out.String(); got != expected {
		t.Errorf("For args %v, expected %q but got %q", args, expected, got)
	}
}

// Test cases
func TestMainOutputStandardCases(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"John", "Doe"}, "Hello, John Doe!\n"},
		{[]string{"Alice"}, "Hello, Alice!\n"},
		{[]string{}, "Hello, !\n"}, // No arguments
	}

	for _, test := range tests {
		runTest(test.args, test.expected, t)
	}
}

func TestMainOutputEdgeCases(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"  ", "Extra", "Spaces"}, "Hello,   Extra Spaces!\n"}, // Multiple spaces
		{[]string{"!@#$%^&*()"}, "Hello, !@#$%^&*()!\n"},               // Special characters
		{[]string{strings.Repeat("A", 1000)}, "Hello, " + strings.Repeat("A", 1000) + "!\n"}, // Very long name
		{[]string{"\tTabCharacter"}, "Hello, \tTabCharacter!\n"},       // Tab character
	}

	for _, test := range tests {
		runTest(test.args, test.expected, t)
	}
}

func TestMainOutputMixedInput(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"Hello", "World", "!"}, "Hello, World !\n"},   // Mixed words and symbols
		{[]string{"123", "456"}, "Hello, 123 456!\n"},           // Numeric inputs
		{[]string{"\nNewLine"}, "Hello, \nNewLine!\n"},           // Newline character
		{[]string{"Mixed", "Spaces ", " ", "AndTabs\t"}, "Hello, Mixed Spaces   AndTabs\t!\n"}, // Mixed spaces and tabs
	}

	for _, test := range tests {
		runTest(test.args, test.expected, t)
	}
}
