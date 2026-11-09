package main

import (
	"testing"
)

func TestTransformText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Uppercase single", "hello (up)", "HELLO"},
		{"Lowercase single", "WORLD (low)", "world"},
		{"Capitalize single", "hello (cap)", "Hello"},
		{"Uppercase multiple", "hello world (up, 2)", "HELLO WORLD"},
		{"Lowercase multiple", "HELLO WORLD (low, 2)", "hello world"},
		{"Capitalize multiple", "hello world (cap, 2)", "Hello World"},
		{"Hex to decimal", "1A (hex) files", "26 files"},
		{"Bin to decimal", "101 (bin) files", "5 files"},
		{"Attached command", "ready(up)", "READY"},
		{"No change needed", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformText(tt.input); got != tt.expected {
				t.Errorf("transformText(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFixPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Space before punctuation", "hello , world", "hello, world"},
		{"Space before grouped", "hello ... world", "hello... world"},
		{"Ensure space after", "hello,world", "hello, world"},
		{"Mixed punctuation spacing", "hello ! ? world", "hello!? world"},
		{"No spaces to fix", "hello, world.", "hello, world."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixPunctuation(tt.input); got != tt.expected {
				t.Errorf("fixPunctuation(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFixQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Spaces inside quotes", "I am ' exactly ' like this", "I am 'exactly' like this"},
		{"No spaces inside quotes", "I am 'exactly' like this", "I am 'exactly' like this"},
		{"Multiple quoted words", "The ' quick ' brown ' fox '", "The 'quick' brown 'fox'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixQuotes(tt.input); got != tt.expected {
				t.Errorf("fixQuotes(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFixArticles(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Lowercase 'a' to 'an'", "a apple", "an apple"},
		{"Uppercase 'A' to 'An'", "A apple", "An apple"},
		{"Lowercase 'a' stays 'a'", "a banana", "a banana"},
		{"Handles 'h'", "a hour", "an hour"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixArticles(tt.input); got != tt.expected {
				t.Errorf("fixArticles(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFullPipeline(t *testing.T) {
	input := "I have 1E (hex) apples ,and they are ' amazing ' !!"
	expected := "I have 30 apples, and they are 'amazing'!!"

	result := fixArticles(fixQuotes(fixPunctuation(transformText(input))))

	if result != expected {
		t.Errorf("Full pipeline = %q, want %q", result, expected)
	}
}
