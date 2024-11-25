package main

import (
	"fmt"
	"log"
	AlphaClickCaptcha "GoCaptcha/ClickCaptcha/AlphaClickCaptcha"
)

func main() {
	// Generate the CAPTCHA
	captchaResponse, err := AlphaClickCaptcha.GenerateAlphaClickCaptcha()
	if err != nil {
		log.Fatalf("Error generating CAPTCHA: %v", err)
	}

	// Print the CAPTCHA response
	fmt.Println("Captcha Image (Base64):", captchaResponse.CaptchaImage)
	fmt.Println("Thumb Image (Base64):", captchaResponse.ThumbImage)
	fmt.Println("Dot Data:", captchaResponse.DotData)

	// Additional logic here if needed
}
