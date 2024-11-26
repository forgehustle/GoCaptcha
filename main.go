package main

import (
	RotateCaptcha "GoCaptcha/RotateCaptcha"
	"fmt"
	"log"
)

func main() {
	// Generate the CAPTCHA
	captchaResponse, err := RotateCaptcha.GenerateRotateCaptcha()
	if err != nil {
		log.Fatalf("Error generating CAPTCHA: %v", err)
	}

	// Print the CAPTCHA response
	fmt.Println("Captcha Image (Base64):", captchaResponse.CaptchaImage)
	fmt.Println("Thumb Image (Base64):", captchaResponse.ThumbImage)
	fmt.Println("Dot Data:", captchaResponse.Angle)
}
