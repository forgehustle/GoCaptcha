package ShapeClickCaptcha

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/resources/shapes"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/click"
)

// CaptchaResponse represents the complete CAPTCHA response structure.
type CaptchaResponse struct {
	CaptchaImage string        `json:"captcha_image"` // Base64-encoded CAPTCHA image
	ThumbImage   string        `json:"thumb_image"`   // Base64-encoded thumbnail image
	DotData      []DotDataItem `json:"dot_data"`      // Slice of dot data (shapes)
}

// DotDataItem represents a single clickable dot (shape).
type DotDataItem struct {
	Index  int     `json:"index"`
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Size   int     `json:"size"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Text   string  `json:"text"`
	Shape  string  `json:"shape"`
	Angle  float64 `json:"angle"`
	Color  string  `json:"color"`
	Color2 string  `json:"color2"`
}

var shapeCaptcha click.Captcha

func init() {
	builder := click.NewBuilder()

	// Load Shapes
	shapeMaps, err := shapes.GetShapes()
	if err != nil {
		log.Fatalln("Error loading shapes:", err)
	}

	// Load Fonts
	font, err := fzshengsksjw.GetFont()
	if err != nil {
		log.Fatalln("Error loading fonts:", err)
	}

	// Load Background Images
	bgImages, err := images.GetImages()
	if err != nil {
		log.Fatalln("Error loading background images:", err)
	}

	// Configure resources
	builder.SetResources(
		click.WithShapes(shapeMaps),             // Shapes to be used in CAPTCHA
		click.WithFonts([]*truetype.Font{font}), // Font for rendering (optional, not directly needed for shapes)
		click.WithBackgrounds(bgImages),         // Background images
	)

	// Configure CAPTCHA settings
	builder.SetOptions(
		click.WithImageSize(option.Size{Width: 300, Height: 200}), // CAPTCHA image dimensions
		click.WithRangeLen(option.RangeVal{Min: 1, Max: 10}),      // Number of clickable shapes
		click.WithRangeColors([]string{"#FF5733", "#33FF57"}),     // Shape colors
		click.WithDisplayShadow(true),                             // Enable shadow
		click.WithShadowColor("#555555"),                          // Shadow color
	)

	// Generate the CAPTCHA instance
	shapeCaptcha = builder.MakeWithShape()
}

func GenerateShapeClickCaptcha() (*CaptchaResponse, error) {
	// Generate CAPTCHA data
	captchaData, err := shapeCaptcha.Generate()
	if err != nil {
		return nil, fmt.Errorf("error generating CAPTCHA: %v", err)
	}

	// Extract dot data
	dotData := captchaData.GetData()
	if dotData == nil {
		return nil, fmt.Errorf("generated CAPTCHA data is empty")
	}

	// Convert dotData (JSON) to Go objects
	var dots []DotDataItem
	dotBytes, _ := json.Marshal(dotData)
	if err := json.Unmarshal(dotBytes, &dots); err != nil {
		return nil, fmt.Errorf("failed to parse dot data: %v", err)
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

	// Return response
	return &CaptchaResponse{
		CaptchaImage: CaptchaImage,
		ThumbImage:   ThumbImage,
		DotData:      dots,
	}, nil
}

