package ChineseClickCaptcha

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/bindata/chars"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/click"
)

// ChinesClickCaptchaData represents the data returned from the CAPTCHA generation process.
type ChineseClickCaptchaData struct {
	CaptchaImage string        `json:"captcha_image"` // Base64-encoded CAPTCHA image
	ThumbImage   string        `json:"thumb_image"`   // Base64-encoded thumbnail image
	DotData      []DotDataItem `json:"dot_data"`      // Slice of dot data (characters)
}

type DotDataItem struct {
	Index  int     `json:"index"`
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Text   string  `json:"text"`
	Angle  float64 `json:"angle"`
	Color  string  `json:"color"`
	Color2 string  `json:"color2"`
}

var textCapt click.Captcha

func init() {
	builder := click.NewBuilder()

	// fonts
	fonts, err := fzshengsksjw.GetFont()
	if err != nil {
		log.Fatalln(err)
	}

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		click.WithChars(chars.GetChineseChars()),
		click.WithFonts([]*truetype.Font{fonts}),
		click.WithBackgrounds(imgs),
	)

	textCapt = builder.Make()
}

func GenerateChineseClickCaptcha() (*ChineseClickCaptchaData, error) {
	// Generate CAPTCHA data
	captchaData, err := textCapt.Generate()
	if err != nil {
		return nil, fmt.Errorf("error generating CAPTCHA: %v", err)
	}

	// Extract dot data
	dotData := captchaData.GetData()
	if dotData == nil {
		return nil, fmt.Errorf("generated CAPTCHA data is empty")
	}

	// Convert dotData (JSON) to Go objects
	var dotMap map[string]DotDataItem
	dotBytes, _ := json.Marshal(dotData)
	if err := json.Unmarshal(dotBytes, &dotMap); err != nil {
		return nil, fmt.Errorf("failed to parse dot data: %v", err)
	}

	// Convert map to slice
	dots := make([]DotDataItem, 0, len(dotMap))
	for _, item := range dotMap {
		dots = append(dots, item)
	}

	// Generate base64 CAPTCHA image
	CaptchaImage, err := captchaData.GetMasterImage().ToBase64DataWithQuality(option.QualityNone)
	if err != nil {
		return nil, fmt.Errorf("failed to get captcha image: %v", err)
	}

	// Generate base64 thumbnail image
	ThumbImage, err := captchaData.GetThumbImage().ToBase64Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get puzzle image: %v", err)
	}

	// Return the response
	return &ChineseClickCaptchaData{
		CaptchaImage: CaptchaImage,
		ThumbImage:   ThumbImage,
		DotData:      dots,
	}, nil
}
