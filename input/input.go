package input

import "github.com/faiface/pixel/pixelgl"

const (
	Left = iota
	Right
	Up
	Down
	Space
	Shift
	Zed
	Escape
	Enter
)

func initInputArray() []pixelgl.Button {
	inputArray := make([]pixelgl.Button, 0, 32)

	inputArray = append(inputArray, pixelgl.KeyLeft)
	inputArray = append(inputArray, pixelgl.KeyRight)
	inputArray = append(inputArray, pixelgl.KeyUp)
	inputArray = append(inputArray, pixelgl.KeyDown)
	inputArray = append(inputArray, pixelgl.KeySpace)
	inputArray = append(inputArray, pixelgl.KeyLeftShift)
	inputArray = append(inputArray, pixelgl.KeyZ)
	inputArray = append(inputArray, pixelgl.KeyEscape)
	inputArray = append(inputArray, pixelgl.KeyEnter)

	return inputArray
}

var InputArray []pixelgl.Button = initInputArray()
