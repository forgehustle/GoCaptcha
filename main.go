package main

import (
	PuzzleSliderCaptcha "GoCaptcha/Captcha"
	"fmt"
	"log"
)

func main() {
	captchaData, err := PuzzleSliderCaptcha.GeneratePuzzleSliderCaptcha()
	if err != nil {
		log.Fatalln("Error generating captcha:", err)
	}

	fmt.Printf("CaptchaImage: %s\n", captchaData.CaptchaImage)
	fmt.Printf("PuzzleImage: %s\n", captchaData.PuzzleImage)
}
