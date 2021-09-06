package image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

const (
	OverlayStartPointTopLeft = iota + 1
	OverlayStartPointTopMiddle
	OverlayStartPointTopRight
	OverlayStartPointMiddleLeft
	OverlayStartPointCenter
	OverlayStartPointMiddleRight
	OverlayStartPointBottomLeft
	OverlayStartPointBottomMiddle
	OverlayStartPointBottomRight
)

func getStartPoint(pointType int, rgba *image.RGBA, i *image.Image) (point *image.Point, err error) {
	xPoint := 0
	yPoint := 0

	switch pointType {
	case OverlayStartPointTopLeft:
	case OverlayStartPointTopMiddle:
		xPoint = (rgba.Bounds().Size().X - (*i).Bounds().Size().X) / 2
	case OverlayStartPointTopRight:
		xPoint = rgba.Bounds().Size().X - (*i).Bounds().Size().X
	case OverlayStartPointMiddleLeft:
		yPoint = (rgba.Bounds().Size().Y - (*i).Bounds().Size().Y) / 2
	case OverlayStartPointCenter:
		xPoint = (rgba.Bounds().Size().X - (*i).Bounds().Size().X) / 2
		yPoint = (rgba.Bounds().Size().Y - (*i).Bounds().Size().Y) / 2
	case OverlayStartPointMiddleRight:
		xPoint = rgba.Bounds().Size().X - (*i).Bounds().Size().X
		yPoint = (rgba.Bounds().Size().Y - (*i).Bounds().Size().Y) / 2
	case OverlayStartPointBottomLeft:
		yPoint = rgba.Bounds().Size().Y - (*i).Bounds().Size().Y
	case OverlayStartPointBottomMiddle:
		xPoint = (rgba.Bounds().Size().X - (*i).Bounds().Size().X) / 2
		yPoint = rgba.Bounds().Size().Y - (*i).Bounds().Size().Y
	case OverlayStartPointBottomRight:
		xPoint = rgba.Bounds().Size().X - (*i).Bounds().Size().X
		yPoint = rgba.Bounds().Size().Y - (*i).Bounds().Size().Y
	default:
		return nil, fmt.Errorf("重ね合わせ開始ポイントが不正です 値:%v", pointType)
	}

	return &(image.Point{xPoint, yPoint}), nil
}

func OverlayImage(pointType int, rgba *image.RGBA, i *image.Image) error {
	point, err := getStartPoint(pointType, rgba, i)
	if err != nil {
		return err
	}
	rct := image.Rectangle{*point, rgba.Bounds().Size()}
	draw.Draw(rgba, rct, *i, image.Point{0, 0}, draw.Over)
	return nil
}

func GetPngImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}
