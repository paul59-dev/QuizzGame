package colors

import "github.com/fatih/color"

func Color(str string, colors string, value ...interface{}) {
	var printer *color.Color

	switch colors {
	case "yellow":
		printer = color.New(color.FgYellow)
	case "red":
		printer = color.New(color.FgRed)
	case "cyan":
		printer = color.New(color.FgCyan)
	case "green":
		printer = color.New(color.FgGreen)
	case "blue":
		printer = color.New(color.FgBlue)
	default:
		color.Red("Impossible to find the color")
	}

	if printer != nil {
		if len(value) != 0 {
			printer.Print(str, value)
		} else {
			printer.Printf(str)
		}
	}
}
