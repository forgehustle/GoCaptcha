package ShapeClickCaptcha

import (
	ShapeClickCaptcha "GoCaptcha/ClickCaptcha/ShapeClickCaptcha"
	"testing"
)

// TestGenerateShapeClickCaptcha tests the GenerateShapeClickCaptcha function.
func TestGenerateShapeClickCaptcha(t *testing.T) {
	// Call the function
	result, err := ShapeClickCaptcha.GenerateShapeClickCaptcha()

	// Check for errors
	if err != nil {
		t.Fatalf("GenerateShapeClickCaptcha returned an error: %v", err)
	}

	// Verify the result is not nil
	if result == nil {
		t.Fatalf("GenerateShapeClickCaptcha returned nil result")
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
		if dot.X < 0 || dot.Y < 0 {
			t.Errorf("DotData item has invalid coordinates: %+v", dot)
		}
		if dot.Angle < 0 || dot.Angle > 360 {
			t.Errorf("DotData item has an invalid angle: %+v", dot)
		}
		if dot.Size <= 0 {
			t.Errorf("DotData item has invalid size: %+v", dot)
		}
		if dot.Width <= 0 || dot.Height <= 0 {
			t.Errorf("DotData item has invalid width or height: %+v", dot)
		}
	}
}
