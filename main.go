package main
import (
	"fmt"
	"log"
	 PuzzleSliderCaptcha "GoCaptcha/Captcha"
)

func main() {
    captchaData, err := PuzzleSliderCaptcha.GeneratePuzzleSliderCaptcha()
    if err != nil {
        log.Fatalln("Error generating captcha:", err)
    }

    fmt.Printf("CaptchaImage: %s\n", captchaData.CaptchaImage)
    fmt.Printf("PuzzleImage: %s\n", captchaData.PuzzleImage)
}