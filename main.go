package main

import (
	"fmt"
	"image"
	"image-overlay/file"
	myImg "image-overlay/image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	backgroundImage, err := getPngImageFromSourceFileName("background.png")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	rgba := image.NewRGBA(image.Rectangle{image.Point{0, 0}, backgroundImage.Bounds().Size()})
	draw.Draw(
		rgba,
		image.Rectangle{image.Point{0, 0}, backgroundImage.Bounds().Size()},
		backgroundImage, image.Point{0, 0},
		draw.Src,
	)

	list := map[int]string{
		myImg.OverlayStartPointTopLeft:      "top-left.png",
		myImg.OverlayStartPointTopMiddle:    "top-middle.png",
		myImg.OverlayStartPointTopRight:     "top-right.png",
		myImg.OverlayStartPointMiddleLeft:   "middle-left.png",
		myImg.OverlayStartPointCenter:       "center.png",
		myImg.OverlayStartPointMiddleRight:  "middle-right.png",
		myImg.OverlayStartPointBottomLeft:   "bottom-left.png",
		myImg.OverlayStartPointBottomMiddle: "bottom-middle.png",
		myImg.OverlayStartPointBottomRight:  "bottom-right.png",
	}
	for overlayStartPoint, fileName := range list {
		err = overlayImage(overlayStartPoint, rgba, fileName)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}

	outputFilePath, err := file.GetOutputFilePath("output.png")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer outputFile.Close()

	png.Encode(outputFile, rgba)
}

func getPngImageFromSourceFileName(sourceFileName string) (image.Image, error) {
	filePath, err := file.GetSourceFilePath(sourceFileName)
	if err != nil {
		return nil, err
	}
	fileImage, err := myImg.GetPngImage(filePath)
	if err != nil {
		return nil, err
	}
	return fileImage, nil
}

func overlayImage(overlayStartPoint int, rgba *image.RGBA, sourceFileName string) (err error) {
	fileImage, err := getPngImageFromSourceFileName(sourceFileName)
	if err != nil {
		return err
	}
	myImg.OverlayImage(overlayStartPoint, rgba, &fileImage)
	if err != nil {
		return err
	}
	return nil
}
