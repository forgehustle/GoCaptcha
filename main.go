package main

import (
	"fmt"
	"log"
	ShapeClickCaptcha "GoCaptcha/ClickCaptcha/ShapeClickCaptcha"
)

func main() {
	// Generate the CAPTCHA
	captchaResponse, err := ShapeClickCaptcha.GenerateShapeClickCaptcha()
	if err != nil {
		log.Fatalf("Error generating CAPTCHA: %v", err)
	}

	// Print the CAPTCHA response
	fmt.Println("Captcha Image (Base64):", captchaResponse.CaptchaImage)
	fmt.Println("Thumb Image (Base64):", captchaResponse.ThumbImage)
	fmt.Println("Dot Data:", captchaResponse.DotData)

	// Additional logic here if needed
}
