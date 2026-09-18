package pf_image_2fjpeg

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type Image []any

type Cmyk struct {
	C int
	M int
	Y int
	K int
}

type Rgba struct {
	R int
	G int
	B int
	A int
}

type Gray struct {
	G int
}

type Rgb_64 struct {
	R int
	G int
	B int
}

type Rgba_64 struct {
	R int
	G int
	B int
	A int
}

type YCbCr struct {
	Y  int
	Cb int
	Cr int
}

type Rgb struct {
	R int
	G int
	B int
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Rgb":   func(t uint32, v any) any { return Rgb{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int)} },
	"Image": func(t uint32, v any) any { return Image(v.([]any)) },
	"Cmyk": func(t uint32, v any) any {
		return Cmyk{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int), v.([]any)[3].(int)}
	},
	"Rgba": func(t uint32, v any) any {
		return Rgba{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int), v.([]any)[3].(int)}
	},
	"Gray":   func(t uint32, v any) any { return Gray{v.([]any)[0].(int)} },
	"Rgb_64": func(t uint32, v any) any { return Rgb_64{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int)} },
	"Rgba_64": func(t uint32, v any) any {
		return Rgba_64{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int), v.([]any)[3].(int)}
	},
	"YCbCr": func(t uint32, v any) any { return YCbCr{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int)} },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Gray":    (*Gray)(nil),
	"Rgb_64":  (*Rgb_64)(nil),
	"Rgba_64": (*Rgba_64)(nil),
	"YCbCr":   (*YCbCr)(nil),
	"Rgb":     (*Rgb)(nil),
	"Image":   (*Image)(nil),
	"Cmyk":    (*Cmyk)(nil),
	"Rgba":    (*Rgba)(nil),
}

func GoGetJpegFile(filename string) any {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	im, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	result := Image([]any{})
	for col := im.Bounds().Min.X; col < im.Bounds().Max.X; col++ {
		newColumn := []any{}
		for row := im.Bounds().Min.Y; row < im.Bounds().Max.Y; row++ {
			newColumn = append(newColumn, goColorToPfColor(im.At(col, row)))
		}
		result = append(result, newColumn)
	}
	return result
}

func RegisterColors(a Cmyk, b Rgb, c Rgba, d Gray, e Rgb_64, f Rgba_64, g YCbCr) any {
	return struct{}{}
}

func GoSaveJpegFile(I Image, width int, height int, filename string, quality int) any {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	img := pfImageToGoImage(I, width, height)
	if quality == -1 {
		err = jpeg.Encode(file, img, nil)
	} else {
		err = jpeg.Encode(file, img, &jpeg.Options{quality})
	}
	if err != nil {
		return err
	}
	return struct{}{}
}

func goColorToPfColor(pixel color.Color) Rgba {
	r, g, b, a := pixel.RGBA()
	return Rgba{int(r) / 256, int(g) / 256, int(b) / 256, int(a) / 256}
}

func pfImageToGoImage(I Image, width, height int) image.Image {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for x, col := range I {
		for y, pixel := range col.([]any) {
			switch pixel := pixel.(type) {
			case Cmyk:
				img.Set(x, y, color.CMYK{uint8(pixel.C), uint8(pixel.M), uint8(pixel.Y), uint8(pixel.K)})
			case Gray:
				img.Set(x, y, color.Gray{uint8(pixel.G)})
			case Rgb:
				img.Set(x, y, color.RGBA{uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), 255})
			case Rgb_64:
				img.Set(x, y, color.RGBA{uint8(pixel.R / 256), uint8(pixel.G / 256), uint8(pixel.B / 256), 255})
			case Rgba:
				img.Set(x, y, color.RGBA{uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), uint8(pixel.A)})
			case Rgba_64:
				img.Set(x, y, color.RGBA{uint8(pixel.R / 256), uint8(pixel.G / 256), uint8(pixel.B / 256), uint8(pixel.A / 256)})
			case YCbCr:
				img.Set(x, y, color.YCbCr{uint8(pixel.Y), uint8(pixel.Cb), uint8(pixel.Cr)})
			default:
				panic("Unhandled case")
			}
		}
	}
	return img
}
