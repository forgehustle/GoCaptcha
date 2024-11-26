package RotateCaptchaTest

import (
	"testing"

	"GoCaptcha/RotateCaptcha" // Update this import path based on your project structure
)

func TestGenerateRotateCaptcha(t *testing.T) {
	// Call the GenerateRotateCaptcha function
	captchaData, err := RotateCaptcha.GenerateRotateCaptcha()
	if err != nil {
		t.Fatalf("Failed to generate rotate CAPTCHA: %v", err)
	}

	// Assert that the returned data is not nil
	if captchaData == nil {
		t.Fatal("Returned RotateCaptchaData is nil")
	}

	// Assert that CaptchaImage and ThumbImage are not empty
	if captchaData.CaptchaImage == "" {
		t.Error("CaptchaImage is empty")
	}
	if captchaData.ThumbImage == "" {
		t.Error("ThumbImage is empty")
	}

	// Assert that ParentWidth, ParentHeight, Width, Height, and Angle have valid values
	if captchaData.ParentWidth <= 0 {
		t.Errorf("Invalid ParentWidth: %d", captchaData.ParentWidth)
	}
	if captchaData.ParentHeight <= 0 {
		t.Errorf("Invalid ParentHeight: %d", captchaData.ParentHeight)
	}
	if captchaData.Width <= 0 {
		t.Errorf("Invalid Width: %d", captchaData.Width)
	}
	if captchaData.Height <= 0 {
		t.Errorf("Invalid Height: %d", captchaData.Height)
	}
	if captchaData.Angle < 0 || captchaData.Angle > 360 {
		t.Errorf("Invalid Angle: %d", captchaData.Angle)
	}
}
