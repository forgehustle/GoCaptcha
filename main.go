package main

import (
	SlideCaptcha "GoCaptcha/SlideCaptcha"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	captchaData, err := SlideCaptcha.GenerateSlideCaptcha()
	if err != nil {
		log.Fatalln("Error generating captcha:", err)
	}
	// Marshal dotData and print (for debugging)
	dots, _ := json.Marshal(captchaData)
	//fmt.Println(">>>>> ", string(dots))
	fmt.Println(string(dots))
}
