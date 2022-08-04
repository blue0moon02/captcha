package captcha

import (
	"github.com/fogleman/gg"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func GenerateCaptchaImage(width, height int, fontPath string, opts ...map[string]interface{}) (image.Image, string) {
	if width <= 0 || height <= 0 {
		panic("A negative number or zero was entered")
	}
	// default text option
	var (
		bgColor   = color.RGBA{R: 220, G: 220, B: 220, A: 255}
		textCount = 6
		lineCount = 6
	)

	if opts != nil {
		for k, v := range opts[0] {
			switch k {
			case "bgColor":
				bgColor = v.(color.RGBA)
				break

			case "textCount":
				textCount = v.(int)
				break

			case "lineCount":
				lineCount = v.(int)
				break
			}
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	textImg, key := generateRandomTextImage(width, height, textCount, fontPath, opts...)
	draw.Draw(dst, dst.Bounds(), textImg, image.Point{}, draw.Over)

	lineImg := generateRandomLineImage(width, height, lineCount, opts...)
	draw.Draw(dst, dst.Bounds(), lineImg, image.Point{}, draw.Over)

	return dst, key
}

func generateRandomTextImage(width, height, count int, fontPath string, opts ...map[string]interface{}) (*image.RGBA, string) {
	if width <= 0 || height <= 0 || count <= 0 {
		panic("A negative number or zero was entered")
	}

	// default text option
	var (
		fontSize        = 72.0
		textColor       = color.RGBA{R: 0, G: 0, B: 0, A: 255}
		moveRange       = [2]float64{26.0, 26.0}
		randomColorMode = false
		randomAlphaMode = false

		arrText = []rune{
			'a', 'b', 'd', 'e', 'f', 'g', 'h', 'j', 'm', 'n', 'p', 'q', 'r', 't', 'y',
			'A', 'B', 'D', 'E', 'F', 'G', 'H', 'J', 'M', 'N', 'P', 'Q', 'R', 'T', 'Y',
			'1', '2', '3', '4', '5', '7', '8',
		}
	)

	if opts != nil {
		for k, v := range opts[0] {
			switch k {
			case "fontSize":
				fontSize = v.(float64)
				break

			case "textColor":
				textColor = v.(color.RGBA)
				break

			case "textMoveRange":
				moveRange = v.([2]float64)
				break

			case "textRandomColorMode":
				randomColorMode = v.(bool)
				break

			case "textRandomAlphaMode":
				randomAlphaMode = v.(bool)
				break

			case "arrText":
				arrText = v.([]rune)
				break
			}
		}
	}

	dst := image.NewNRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContext(width, height)
	dc.DrawImage(dst, 0, 0)
	dc.SetColor(textColor)

	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		return nil, ""
	}

	var (
		x   = float64(width / (count * 2))
		key string
	)
	for c := count; c > 0; c-- {
		if randomColorMode {
			textColor.R, textColor.G, textColor.B = uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
		}
		if randomAlphaMode {
			textColor.A = uint8(rand.Intn(255))
		}
		dc.SetColor(textColor)

		y := float64(height / 2)
		maxWidth := float64(width) - 60.0
		randText := string(arrText[rand.Intn(len(arrText))])
		dc.DrawStringWrapped(randText, x+randFloat(-moveRange[0], moveRange[0]), y+randFloat(-moveRange[1], moveRange[1]), 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
		key += randText

		x += float64(width/(count*2)) * 2
	}

	return dc.Image().(*image.RGBA), key
}

func generateDummyTextImage(width, height, count int, fontPath string, opts ...map[string]interface{}) *image.RGBA {
	if width <= 0 || height <= 0 || count <= 0 {
		panic("A negative number or zero was entered")
	}

	// default text option
	var (
		fontSize        = 64.0
		textColor       = color.RGBA{R: 0, G: 0, B: 0, A: 255}
		moveRange       = [2]float64{26.0, 26.0}
		randomColorMode = false
		randomAlphaMode = false

		arrText = []rune{
			'a', 'b', 'd', 'e', 'f', 'g', 'h', 'j', 'm', 'n', 'p', 'q', 'r', 't', 'y',
			'A', 'B', 'D', 'E', 'F', 'G', 'H', 'J', 'M', 'N', 'P', 'Q', 'R', 'T', 'Y',
			'1', '2', '3', '4', '5', '7', '8',
		}
	)

	if opts != nil {
		for key, value := range opts[0] {
			switch key {
			case "fontSize":
				fontSize = value.(float64)
				break

			case "textColor":
				textColor = value.(color.RGBA)
				break

			case "textMoveRange":
				moveRange = value.([2]float64)
				break

			case "textRandomColorMode":
				randomColorMode = value.(bool)
				break

			case "textRandomAlphaMode":
				randomAlphaMode = value.(bool)
				break

			case "arrText":
				arrText = value.([]rune)
				break
			}
		}
	}

	dst := image.NewNRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContext(width, height)
	dc.DrawImage(dst, 0, 0)
	dc.SetColor(textColor)

	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		return nil
	}

	var x = float64(width / (count * 2))
	for c := count; c > 0; c-- {
		if randomColorMode {
			textColor.R, textColor.G, textColor.B = uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
		}
		if randomAlphaMode {
			textColor.A = uint8(rand.Intn(255))
		}
		dc.SetColor(textColor)

		y := float64(height / 2)
		maxWidth := float64(width) - 60.0
		dc.DrawStringWrapped(string(arrText[rand.Intn(len(arrText))]), x+randFloat(-moveRange[0], moveRange[0]), y+randFloat(-moveRange[1], moveRange[1]), 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

		x += float64(width/(count*2)) * 2
	}

	return dc.Image().(*image.RGBA)
}

func generateRandomLineImage(width, height, count int, opts ...map[string]interface{}) *image.RGBA {
	if width <= 0 || height <= 0 || count <= 0 {
		panic("A negative number or zero was entered")
	}

	// default line option
	var (
		lineThickness   = 5.0
		lineColor       = color.RGBA{R: 0, G: 0, B: 0, A: 255}
		randomColorMode = false
		randomAlphaMode = false
	)

	if opts != nil {
		for key, value := range opts[0] {
			switch key {
			case "lineThickness":
				lineThickness = value.(float64)
				break

			case "lineColor":
				lineColor = value.(color.RGBA)
				break

			case "lineRandomColorMode":
				randomColorMode = value.(bool)
				break

			case "lineRandomAlphaMode":
				randomAlphaMode = value.(bool)
				break
			}
		}
	}

	// line point position
	var x1, y1, x2, y2 float64

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(dst)

	for c := count; c > 0; c-- {
		if randBool() {
			x1 = float64(rand.Intn(width))
			y1 = 0
		} else {
			x1 = 0
			y1 = float64(rand.Intn(height))
		}
		if randBool() {
			x2 = float64(rand.Intn(width))
			y2 = float64(height)
		} else {
			x2 = float64(width)
			y2 = float64(rand.Intn(height))
		}
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y2)

		if randomColorMode {
			lineColor.R, lineColor.G, lineColor.B = uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
		}
		if randomAlphaMode {
			lineColor.A = uint8(rand.Intn(255))
		}
		gc.SetStrokeColor(lineColor)

		gc.SetLineWidth(lineThickness)
		gc.Stroke()
	}

	return dst
}

func distortImage() {
	return
}

func randBool() bool {
	if rand.Intn(2) == 0 {
		return true
	}
	return false
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
