package SlideCaptcha

import (
	"fmt"
	"log"

	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/slide"
)

// PuzzleSliderCaptcha defines the return structure
type SlideCaptchaData struct {
	CaptchaImage string `json:"CaptchaImage"`
	PuzzleImage  string `json:"PuzzleImage"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Angle        int    `json:"angle"`
	TileX        int    `json:"tile_x"`
	TileY        int    `json:"tile_y"`
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

func GenerateSlideCaptcha() (*SlideCaptchaData, error) {
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

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate error")
	}

	return &SlideCaptchaData{
		CaptchaImage: captchaImage,
		PuzzleImage:  puzzleImage,
		X:            dotData.X,
		Y:            dotData.Y,
		Width:        dotData.Width,
		Height:       dotData.Height,
		Angle:        dotData.Angle,
		TileX:        dotData.TileX,
		TileY:        dotData.TileY,
	}, nil
}
