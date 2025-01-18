package token

import (
	"testing"

	"github.com/gbburleigh/quick-card-tokenizer/internal/db"
)

func TestTokenNotFound(t *testing.T) {
	database := db.Create()
	defer database.Close()

	_, err := Query("nonexistenttoken", database)
	if err == nil {
		t.Error("Expected token not found error, but got nil")
	}
}

func TestMask(t *testing.T) {
	testCases := []struct {
		name     string
		pan      string
		expected string
	}{
		{
			name:     "Empty PAN",
			pan:      "",
			expected: "",
		},
		{
			name:     "Short PAN",
			pan:      "123",
			expected: "123",
		},
		{
			name:     "Long PAN",
			pan:      "1234567890123456",
			expected: "************3456",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			masked := Mask(tc.pan)
			if masked != tc.expected {
				t.Errorf("For PAN %s, expected %s, got %s", tc.pan, tc.expected, masked)
			}
		})
	}
}

func TestRun(t *testing.T) {
	TestMask(t)
	TestTokenNotFound(t)
}
