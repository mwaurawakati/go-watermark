/**********************************************************
* 2024/07/05
* Author: Michael Putong
* This code free to use, share and modify
* Author not responsible for any damage caused by this code
***********************************************************/

package go_watermark

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

type Position struct {
	PosX  float64
	PosY  float64
	PosAX float64
	PosAY float64
}

type Font struct {
	FontName string
	FontSize float64
}

type Watermark struct {
	Image      string
	OutputFile string
	Logo       string
	Text       string
	Position
	Font
	Color       color.Color
	Align       Align
	LineSpacing float64
	Repeat
	Rotate float64
	ImgSize
	AddLogoFile bool
}

type Repeat struct {
	RepX, RepY, WordSpacing int
}

type ImgSize struct {
	Width, Height int
}

func AddWatermark(watermark *Watermark) error {
	srcImage, err := imageDecode(watermark.Image)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	var logoFile image.Image

	bgImage := imageResize(srcImage, watermark.ImgSize.Width, watermark.ImgSize.Height)

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	// Add logoFile
	if watermark.AddLogoFile {
		logoFile, err = imageDecode(watermark.Logo)
		if err != nil {
			return fmt.Errorf("error decoding image: %v", err)
		}
		logoFile = imaging.Resize(logoFile, 60, 60, imaging.Lanczos)
		logoHeight := logoFile.Bounds().Dy()
		x := 10.0
		y := float64(dc.Height()) - float64(logoHeight) - 10.0
		dc.DrawImage(logoFile, int(x), int(y))
	}

	// Add IG Image
	igImg, err := imageDecode("instagram.jpeg")
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	igImg = imaging.Resize(igImg, 40, 40, imaging.Lanczos)
	igImgHeight := igImg.Bounds().Dy()
	x := imgWidth - 320.0
	y := float64(dc.Height()) - float64(igImgHeight) - 10.0
	dc.DrawImage(igImg, int(x), int(y))

	// Add phone Image
	phoneImg, err := imageDecode("phone.png")
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	phoneImg = imaging.Resize(phoneImg, 40, 40, imaging.Lanczos)
	phoneImgHeight := igImg.Bounds().Dy()
	x = imgWidth - 320.0
	y = float64(dc.Height()) - float64(phoneImgHeight) - 60.0
	dc.DrawImage(phoneImg, int(x), int(y))

	fontByte := goregular.TTF
	if len(watermark.FontName) > 0 {
		fontByte, err = loadFont(watermark.FontName)
		if err != nil {
			return fmt.Errorf("error loading font %q: %v", watermark.FontName, err)
		}
	}

	font, err := truetype.Parse(fontByte)
	if err != nil {
		return fmt.Errorf("error in truetype.Parse: %v", err)
	}
	fontSize := 40.0
	f, _ := gg.LoadFontFace("./arial.ttf", fontSize)
	dc.SetFontFace(f)
	dc.SetColor(image.NewUniform(color.RGBA{255, 255, 0, 255}))

	text := "@nazar.hajy"
	textX := float64(imgWidth) - 50.0
	textY := float64(imgHeight) - 60.0
	dc.DrawStringAnchored(text, textX, textY+5, 1, 1)
	dc.DrawStringAnchored("862-60-50-50", textX+20, textY-45, 1, 1)
	DrawWatermark(font, watermark, dc, float64(imgWidth), float64(imgHeight))

	err = dc.SavePNG(watermark.OutputFile)
	if err != nil {
		return fmt.Errorf("error saving image: %v", err)
	}

	return nil
}

func DrawWatermark(font *truetype.Font, watermark *Watermark, dc *gg.Context, imgWidth, imgHeight float64) {
	maxWidth := imgWidth - 60.0
	posY := int(watermark.PosY)
	if watermark.RepY < 2 {
		watermark.RepY = 0
	}

	if watermark.RepX < 1 {
		watermark.RepX = 1
	}

	y := float64(posY)

	dc.Rotate(gg.Radians(watermark.Rotate))

	for divY := 0; divY <= watermark.RepY-1; divY++ {
		wordSpaces := strings.Repeat(" ", watermark.WordSpacing)
		repTextX := strings.Repeat(watermark.Text+wordSpaces, watermark.RepX)

		face := truetype.NewFace(font, &truetype.Options{Size: watermark.FontSize})
		dc.SetFontFace(face)
		dc.SetColor(watermark.Color)
		dc.DrawStringWrapped(
			repTextX,
			watermark.PosX, y, watermark.PosAX, watermark.PosAY,
			maxWidth,
			watermark.LineSpacing,
			gg.Align(watermark.Align),
		)

		y += watermark.LineSpacing
	}
	face := truetype.NewFace(font, &truetype.Options{Size: watermark.FontSize})
	dc.SetFontFace(face)
	dc.SetColor(watermark.Color)
	dc.DrawString("IG:@MwauraWakati", float64(dc.Width()-200), float64(y))
}

func imageDecode(imageFile string) (image.Image, error) {
	imgFile, err := os.Open(imageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}

	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image file: %v", err)
	}

	return img, nil
}

func imageResize(srcImage image.Image, width, height int) image.Image {
	return imaging.Resize(srcImage, width, height, imaging.Lanczos)
}

func loadFont(fontFile string) ([]byte, error) {
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %v", err)
	}

	return fontBytes, nil
}
