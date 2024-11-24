package main

import (
	PuzzleSliderCaptcha "GoCaptcha/Captcha"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	captchaData, err := PuzzleSliderCaptcha.GeneratePuzzleSliderCaptcha()
	if err != nil {
		log.Fatalln("Error generating captcha:", err)
	}
	// Marshal dotData and print (for debugging)
	dots, _ := json.Marshal(captchaData)
	//fmt.Println(">>>>> ", string(dots))
	fmt.Printf(string(dots))
}
