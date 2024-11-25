package SlideCaptcha

import (
	"testing"
)

func TestGenerateSlideCaptcha(t *testing.T) {
	// Call the function
	captchaData, err := GenerateSlideCaptcha()
	if err != nil {
		t.Fatalf("Error generating captcha: %v", err)
	}

	// Validate the output
	if captchaData == nil {
		t.Fatalf("Expected non-nil captcha data, got nil")
	}

	if captchaData.CaptchaImage == "" {
		t.Errorf("CaptchaImage is empty")
	}

	if captchaData.PuzzleImage == "" {
		t.Errorf("PuzzleImage is empty")
	}
}
