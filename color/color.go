package color

import "github.com/fatih/color"

func Colorprint(a, b string) { //a는 텍스트 , b는 색깔
	switch b {
	case "red":
		color.Red(a)
	case "yellow":
		color.Yellow(a)
	case "green":
		color.Green(a)
	case "cyan":
		color.Cyan(a)
	case "blue":
		color.Blue(a)
	case "magenta":
		color.Magenta(a)
	case "white":
		color.White(a)
	case "red+underline":
		redunderline := color.New(color.FgRed).Add(color.Underline)
		redunderline.Println(a)
	case "yellow+underline":
		yellowunderline := color.New(color.FgYellow).Add(color.Underline)
		yellowunderline.Println(a)
	case "green+underline":
		greenunderline := color.New(color.FgGreen).Add(color.Underline)
		greenunderline.Println(a)
	case "cyan+underline":
		cyanunderline := color.New(color.FgCyan).Add(color.Underline)
		cyanunderline.Println(a)
	case "blue+underline":
		blueunderline := color.New(color.FgBlue).Add(color.Underline)
		blueunderline.Println(a)
	case "magenta+underline":
		magentaunderline := color.New(color.FgMagenta).Add(color.Underline)
		magentaunderline.Println(a)
	case "white+underline":
		whiteunderline := color.New(color.FgWhite).Add(color.Underline)
		whiteunderline.Println(a)
	}
}
