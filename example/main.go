package main

import (
	"fmt"
	"image/color"
	"github.com/fogleman/gg"
	"image"
	"image/png"
	"os"
	goWatermark "github.com/michaelwp/go-watermark"
	"github.com/disintegration/imaging"
)

func main() {
	err := goWatermark.AddWatermark(
		&goWatermark.Watermark{
			Image:      "input.jpg",
			OutputFile: "output1.jpeg",
			Text:       "PandaExpress",
			Position: goWatermark.Position{
				PosAY: 0,
				PosAX: 0,
				PosX: -100,
				PosY:-10,
			},
			Font: goWatermark.Font{
				FontSize: 40,
			},
			Color: color.RGBA{
				R: 255,
				G: 165,
				B: 0,
				A: 200,
			},
			Align: goWatermark.AlignCenter,
			Repeat: goWatermark.Repeat{
				RepY: 20,
				RepX: 10,
				WordSpacing: 40,
			},
			LineSpacing: 200,
			Rotate:      -30,
			ImgSize: goWatermark.ImgSize{
				Width: 1024,
			},
			AddLogoFile: true,
			Logo: "icon.png",
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Watermark added successfully!")
	}
	// Paths to the images
	//imagePath := "output1.jpeg"
	//logoPath := "icon.png"
	//outputPath := "final_image.png"

	// Add logo to the image
	/*err = addLogoToImage(imagePath, logoPath, outputPath)
	if err != nil {
		panic(err)
	}*/
}

/*func addLogoToImage(imagePath, logoPath, outputPath string) error {
	// Load the original image with watermark
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	// Load the logo image
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return err
	}
	defer logoFile.Close()

	logo, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}
	logo = imaging.Resize(logo, 60,60, imaging.Lanczos)
	// Create a new context with the size of the original image
	dc := gg.NewContextForImage(img)

	// Calculate the position to place the logo (bottom left corner)
	//logoWidth := logo.Bounds().Dx()
	logoHeight := logo.Bounds().Dy()
	x := 10.0  // Add some padding from the left edge
	y := float64(dc.Height()) - float64(logoHeight) - 10.0 // Add some padding from the bottom edge

	// Draw the logo onto the image
	dc.DrawImage(logo, int(x), int(y))
	dc.DrawString("IG:@Mwau", float64(dc.Width()-200), float64(y))
	// Save the final image
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return png.Encode(outputFile, dc.Image())
}
*/

func addLogoToImage(imagePath, logoPath, outputPath string) error {
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return err
	}
	defer logoFile.Close()

	logo, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}
	logo = imaging.Resize(logo, 60, 60, imaging.Lanczos)
	dc := gg.NewContextForImage(img)

	logoHeight := logo.Bounds().Dy()
	x := 10.0 
	y := float64(dc.Height()) - float64(logoHeight) - 10.0 
	dc.DrawImage(logo, int(x), int(y))

	fontSize := 20.0
	f, _:=gg.LoadFontFace("./arial.ttf", fontSize)
	dc.SetFontFace(f) 
	dc.SetColor(image.NewUniform(color.RGBA{255, 255, 0, 255})) 

	text := "IG:@nazar.hajy"
	textX := float64(dc.Width()) - 10.0 
	textY := float64(dc.Height()) - 30.0 
	dc.DrawStringAnchored(text, textX, textY, 1, 1)
	dc.DrawStringAnchored("PhoneNumber: 862-60-50-50", textX, textY - 40, 1, 1)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return png.Encode(outputFile, dc.Image())
}