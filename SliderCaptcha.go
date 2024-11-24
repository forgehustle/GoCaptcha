package PuzzleSliderCaptcha

import (
	"fmt"
	"log"

	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/slide"
)

// PuzzleSliderCaptcha defines the return structure
type PuzzleSliderCaptcha struct {
	CaptchaImage string `json:"CaptchaImage"`
	PuzzleImage  string `json:"PuzzleImage"`
}

var slideCapt slide.Captcha

func init() {
	builder := slide.NewBuilder(
			slide.WithEnableGraphVerticalRandom(true),
	)

	// Load resources
	imgs, err := images.GetImages()
	if err != nil {
			log.Fatalln("Failed to load background images:", err)
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
			log.Fatalln("Failed to load graph tiles:", err)
	}

	// Prepare graph images
	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
	for _, graph := range graphs {
			newGraphs = append(newGraphs, &slide.GraphImage{
					OverlayImage: graph.OverlayImage,
					MaskImage:    graph.MaskImage,
					ShadowImage:  graph.ShadowImage,
			})
	}

	builder.SetResources(
			slide.WithGraphImages(newGraphs),
			slide.WithBackgrounds(imgs),
	)

	slideCapt = builder.Make()
}

// GeneratePuzzleSliderCaptcha generates a captcha and returns it as a PuzzleSliderCaptcha struct
func GeneratePuzzleSliderCaptcha() (*PuzzleSliderCaptcha, error) {
	captData, err := slideCapt.Generate()
	if err != nil {
			return nil, fmt.Errorf("failed to generate captcha: %v", err)
	}

	captchaImage, err := captData.GetMasterImage().ToBase64DataWithQuality(option.QualityNone)
	if err != nil {
			return nil, fmt.Errorf("failed to get captcha image: %v", err)
	}

	puzzleImage, err := captData.GetTileImage().ToBase64Data()
	if err != nil {
			return nil, fmt.Errorf("failed to get puzzle image: %v", err)
	}

	return &PuzzleSliderCaptcha{
			CaptchaImage: captchaImage,
			PuzzleImage:  puzzleImage,
	}, nil
}