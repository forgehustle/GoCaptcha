package ShapeClickCaptcha

import (
	"testing"
	ShapeClickCaptcha "GoCaptcha/ClickCaptcha/ShapeClickCaptcha"
)

// TestGenerateShapeClickCaptcha tests the GenerateShapeClickCaptcha function.
func TestGenerateShapeClickCaptcha(t *testing.T) {
	// Call the function to test
	captchaResponse, err := ShapeClickCaptcha.GenerateShapeClickCaptcha()

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	// Verify that CaptchaImage is not empty
	if captchaResponse.CaptchaImage == "" {
		t.Error("CaptchaImage is empty")
	}

	// Verify that ThumbImage is not empty
	if captchaResponse.ThumbImage == "" {
		t.Error("ThumbImage is empty")
	}

	// Verify that DotData has at least one item
	if len(captchaResponse.DotData) == 0 {
		t.Error("DotData is empty")
	}

	// Validate properties of the first dot (if exists)
	if len(captchaResponse.DotData) > 0 {
		firstDot := captchaResponse.DotData[0]
		if firstDot.X < 0 || firstDot.Y < 0 {
			t.Errorf("Invalid position for the first dot: X=%d, Y=%d", firstDot.X, firstDot.Y)
		}
		if firstDot.Width <= 0 || firstDot.Height <= 0 {
			t.Errorf("Invalid size for the first dot: Width=%d, Height=%d", firstDot.Width, firstDot.Height)
		}
	}
}
