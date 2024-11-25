package AlphaClickCaptcha

import (
	AlphaClickCaptcha "GoCaptcha/ClickCaptcha/AlphaClickCaptcha"
	"testing"
)

// TestGenerateAlphaClickCaptcha tests the GenerateAlphaClickCaptcha function.
func TestGenerateAlphaClickCaptcha(t *testing.T) {
	// Call the function
	result, err := AlphaClickCaptcha.GenerateAlphaClickCaptcha()

	// Check for errors
	if err != nil {
		t.Fatalf("GenerateAlphaClickCaptcha returned an error: %v", err)
	}

	// Verify the result is not nil
	if result == nil {
		t.Fatalf("GenerateAlphaClickCaptcha returned nil result")
	}

	// Verify CAPTCHA image is not empty
	if len(result.CaptchaImage) == 0 {
		t.Errorf("CaptchaImage is empty")
	}

	// Verify thumbnail image is not empty
	if len(result.ThumbImage) == 0 {
		t.Errorf("ThumbImage is empty")
	}

	// Verify dot data exists
	if len(result.DotData) == 0 {
		t.Errorf("DotData is empty")
	}

	// Check the contents of DotData
	for _, dot := range result.DotData {
		if dot.Text == "" {
			t.Errorf("DotData item is missing text value: %+v", dot)
		}
		if dot.X < 0 || dot.Y < 0 {
			t.Errorf("DotData item has invalid coordinates: %+v", dot)
		}
		if dot.Angle < 0 || dot.Angle > 360 {
			t.Errorf("DotData item has an invalid angle: %+v", dot)
		}
	}
}
