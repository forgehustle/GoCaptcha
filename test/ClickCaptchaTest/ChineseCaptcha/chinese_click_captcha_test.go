package ChineseClickCaptcha

import (
	"testing"
	"GoCaptcha/ClickCaptcha/ChineseClickCaptcha" // Update with the correct import path
)

func TestGenerateChineseClickCaptcha(t *testing.T) {
	// Generate the Chinese Click CAPTCHA
	data, err := ChineseClickCaptcha.GenerateChineseClickCaptcha()

	// Check if there was an error
	if err != nil {
		t.Fatalf("Error generating captcha: %v", err)
	}

	// Check that the returned data is not nil
	if data == nil {
		t.Fatalf("Captcha data should not be nil")
	}

	// Check that the CAPTCHA image is not empty
	if data.CaptchaImage == "" {
		t.Errorf("Captcha image should not be empty")
	}

	// Check that the thumbnail image is not empty
	if data.ThumbImage == "" {
		t.Errorf("Thumbnail image should not be empty")
	}

	// Assert that the dot data is not empty
	if len(data.DotData) == 0 {
		t.Errorf("Dot data should not be empty")
	}

	// Check that each dot has valid properties
	for _, dot := range data.DotData {
		if dot.X <= 0 {
			t.Errorf("Dot X coordinate should be greater than 0, got: %d", dot.X)
		}
		if dot.Y <= 0 {
			t.Errorf("Dot Y coordinate should be greater than 0, got: %d", dot.Y)
		}
		if dot.Text == "" {
			t.Errorf("Dot text should not be empty")
		}
		if dot.Angle <= 0 {
			t.Errorf("Dot angle should be greater than 0, got: %f", dot.Angle)
		}
		if dot.Color == "" {
			t.Errorf("Dot color should not be empty")
		}
		if dot.Color2 == "" {
			t.Errorf("Dot color2 should not be empty")
		}
	}
}
