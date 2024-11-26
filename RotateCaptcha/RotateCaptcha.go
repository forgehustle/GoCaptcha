package RotateCaptcha

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/rotate"
)

// RotateCaptchaData represents the data returned from the CAPTCHA generation process.
type RotateCaptchaData struct {
	CaptchaImage string `json:"captcha_image"` // Base64-encoded CAPTCHA image
	ThumbImage   string `json:"thumb_image"`   // Base64-encoded thumbnail image
	ParentWidth  int    `json:"parent_width"`  // Parent container width
	ParentHeight int    `json:"parent_height"` // Parent container height
	Width        int    `json:"width"`         // Width of the draggable component
	Height       int    `json:"height"`        // Height of the draggable component
	Angle        int    `json:"angle"`         // Angle of rotation
}

var rotateCapt rotate.Captcha

func init() {
	builder := rotate.NewBuilder(rotate.WithRangeAnglePos([]option.RangeVal{
		{Min: 20, Max: 330},
	}))

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	// set resources
	builder.SetResources(
		rotate.WithImages(imgs),
	)

	rotateCapt = builder.Make()
}

// RotateCaptcha generates a rotate CAPTCHA and returns a struct with image data and dots.
func GenerateRotateCaptcha() (*RotateCaptchaData, error) {
	// Generate CAPTCHA data
	captchaData, err := rotateCapt.Generate()
	if err != nil {
		return nil, fmt.Errorf("error generating CAPTCHA: %v", err)
	}

	// Get dot data
	dotData := captchaData.GetData()
	if dotData == nil {
		return nil, fmt.Errorf("generated CAPTCHA dot data is empty")
	}

	// Parse dot data
	var dotMap map[string]interface{}
	dotBytes, _ := json.Marshal(dotData)
	if err := json.Unmarshal(dotBytes, &dotMap); err != nil {
		return nil, fmt.Errorf("failed to parse dot data: %v", err)
	}

	// Extract specific fields from dot data
	parentWidth := int(dotMap["parent_width"].(float64))
	parentHeight := int(dotMap["parent_height"].(float64))
	width := int(dotMap["width"].(float64))
	height := int(dotMap["height"].(float64))
	angle := int(dotMap["angle"].(float64))

	// Generate base64 CAPTCHA image
	CaptchaImage, err := captchaData.GetMasterImage().ToBase64Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get captcha image: %v", err)
	}

	// Generate base64 thumbnail image
	ThumbImage, err := captchaData.GetThumbImage().ToBase64Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get puzzle image: %v", err)
	}

	// Return the response
	return &RotateCaptchaData{
		CaptchaImage: CaptchaImage,
		ThumbImage:   ThumbImage,
		ParentWidth:  parentWidth,
		ParentHeight: parentHeight,
		Width:        width,
		Height:       height,
		Angle:        angle,
	}, nil
}
