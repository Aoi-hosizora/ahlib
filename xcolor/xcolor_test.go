package xcolor

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestForceColor(t *testing.T) {
	fmt.Println("\x1b[31mhello world\x1b[0m")
	fmt.Println("\x1b[4;31mhello world\x1b[0m")
	fmt.Println("\x1b[31;103mhello world\x1b[0m")
	fmt.Println("\x1b[4;31;103mhello world\x1b[0m")
	ForceColor()
	fmt.Println("\x1b[31mhello world\x1b[0m")
	fmt.Println("\x1b[4;31mhello world\x1b[0m")
	fmt.Println("\x1b[31;103mhello world\x1b[0m")
	fmt.Println("\x1b[4;31;103mhello world\x1b[0m")
}

func TestCode(t *testing.T) {
	for _, tc := range []struct {
		give uint8
		want uint8
	}{
		{Bold.Code(), 1},
		{Faint.Code(), 2},
		{Italic.Code(), 3},
		{Underline.Code(), 4},
		{Reverse.Code(), 7},
		{Strikethrough.Code(), 9},

		{Black.Code(), 30},
		{Red.Code(), 31},
		{Green.Code(), 32},
		{Yellow.Code(), 33},
		{Blue.Code(), 34},
		{Magenta.Code(), 35},
		{Cyan.Code(), 36},
		{White.Code(), 37},
		{Default.Code(), 39},
		{BrightBlack.Code(), 90},
		{BrightRed.Code(), 91},
		{BrightGreen.Code(), 92},
		{BrightYellow.Code(), 93},
		{BrightBlue.Code(), 94},
		{BrightMagenta.Code(), 95},
		{BrightCyan.Code(), 96},
		{BrightWhite.Code(), 97},

		{BGBlack.Code(), 40},
		{BGRed.Code(), 41},
		{BGGreen.Code(), 42},
		{BGYellow.Code(), 43},
		{BGBlue.Code(), 44},
		{BGMagenta.Code(), 45},
		{BGCyan.Code(), 46},
		{BGWhite.Code(), 47},
		{BGDefault.Code(), 49},
		{BGBrightBlack.Code(), 100},
		{BGBrightRed.Code(), 101},
		{BGBrightGreen.Code(), 102},
		{BGBrightYellow.Code(), 103},
		{BGBrightBlue.Code(), 104},
		{BGBrightMagenta.Code(), 105},
		{BGBrightCyan.Code(), 106},
		{BGBrightWhite.Code(), 107},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}
}
func TestString(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{Bold.String(), "1"},
		{Faint.String(), "2"},
		{Italic.String(), "3"},
		{Underline.String(), "4"},
		{Reverse.String(), "7"},
		{Strikethrough.String(), "9"},

		{Black.String(), "30"},
		{Red.String(), "31"},
		{Green.String(), "32"},
		{Yellow.String(), "33"},
		{Blue.String(), "34"},
		{Magenta.String(), "35"},
		{Cyan.String(), "36"},
		{White.String(), "37"},
		{Default.String(), "39"},
		{BrightBlack.String(), "90"},
		{BrightRed.String(), "91"},
		{BrightGreen.String(), "92"},
		{BrightYellow.String(), "93"},
		{BrightBlue.String(), "94"},
		{BrightMagenta.String(), "95"},
		{BrightCyan.String(), "96"},
		{BrightWhite.String(), "97"},

		{BGBlack.String(), "40"},
		{BGRed.String(), "41"},
		{BGGreen.String(), "42"},
		{BGYellow.String(), "43"},
		{BGBlue.String(), "44"},
		{BGMagenta.String(), "45"},
		{BGCyan.String(), "46"},
		{BGWhite.String(), "47"},
		{BGDefault.String(), "49"},
		{BGBrightBlack.String(), "100"},
		{BGBrightRed.String(), "101"},
		{BGBrightGreen.String(), "102"},
		{BGBrightYellow.String(), "103"},
		{BGBrightBlue.String(), "104"},
		{BGBrightMagenta.String(), "105"},
		{BGBrightCyan.String(), "106"},
		{BGBrightWhite.String(), "107"},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}
}

func TestMixCode(t *testing.T) {
	for _, tc := range []struct {
		give []uint8
		want []uint8
	}{
		{MixCode{}.Codes(), []uint8{}},
		{MixCode{Bold.Code()}.Codes(), []uint8{Bold.Code()}},
		{MixCode{Bold.Code(), Bold.Code()}.Codes(), []uint8{Bold.Code(), Bold.Code()}},
		{MixCode{Bold.Code(), Red.Code()}.Codes(), []uint8{Bold.Code(), Red.Code()}},
		{MixCode{Bold.Code(), Red.Code(), BGWhite.Code()}.Codes(), []uint8{Bold.Code(), Red.Code(), BGWhite.Code()}},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}

	for _, tc := range []struct {
		give string
		want string
	}{
		{MixCode{}.String(), "0"},
		{MixCode{Bold.Code(), Bold.Code()}.String(), Bold.String() + ";" + Bold.String()},
		{MixCode{Bold.Code(), Red.Code()}.String(), Bold.String() + ";" + Red.String()},
		{MixCode{Bold.Code(), Red.Code(), BGWhite.Code()}.String(), Bold.String() + ";" + Red.String() + ";" + BGWhite.String()},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}

	for _, tc := range  []struct {
		give MixCode
		want MixCode
	}{
		{Bold.WithStyle(Italic), []uint8{Bold.Code(), Italic.Code()}},
		{Bold.WithColor(Red), []uint8{Bold.Code(), Red.Code()}},
		{Bold.WithBackground(BGWhite), []uint8{Bold.Code(), BGWhite.Code()}},
		{Red.WithStyle(Bold), []uint8{Red.Code(), Bold.Code()}},
		{Red.WithBackground(BGWhite), []uint8{Red.Code(), BGWhite.Code()}},
		{BGWhite.WithStyle(Bold), []uint8{BGWhite.Code(), Bold.Code()}},
		{BGWhite.WithColor(Red), []uint8{BGWhite.Code(), Red.Code()}},
		{MixCode{}.WithStyle(Bold).WithColor(Red).WithBackground(BGWhite), []uint8{Bold.Code(), Red.Code(), BGWhite.Code()}},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}
}

func TestPrint(t *testing.T) {
	Bold.Print("bold\n")
	Faint.Printf("%s\n", "faint")
	Italic.Println("italic")
	fmt.Print(Underline.Sprint("underline\n"))
	fmt.Print(Reverse.Sprintf("%s\n", "reverse"))
	fmt.Print(Strikethrough.Sprintln("strikethrough"))
	fmt.Println()

	Red.Print("red\n")
	Green.Printf("%s\n", "green")
	Yellow.Println("yellow")
	fmt.Print(Blue.Sprint("blue\n"))
	fmt.Print(Magenta.Sprintf("%s\n", "magenta"))
	fmt.Print(Cyan.Sprintln("cyan"))
	Black.Println("black")
	White.Println("white")
	Default.Println("default")
	BrightBlack.Println("bright_black")
	BrightRed.Println("bright_red")
	BrightGreen.Println("bright_green")
	BrightYellow.Println("bright_yellow")
	BrightBlue.Println("bright_blue")
	BrightMagenta.Println("bright_magenta")
	BrightCyan.Println("bright_cyan")
	BrightWhite.Println("bright_white")
	fmt.Println()

	BGRed.Print("bg_red\n")
	BGGreen.Printf("%s\n", "bg_green")
	BGYellow.Println("bg_yellow")
	fmt.Print(BGBlue.Sprint("bg_blue\n"))
	fmt.Print(BGMagenta.Sprintf("%s\n", "bg_magenta"))
	fmt.Print(BGCyan.Sprintln("bg_cyan"))
	BGBlack.Println("bg_black")
	BGWhite.Println("bg_white")
	BGDefault.Println("bg_default")
	BGBrightBlack.Println("bg_bright_black")
	BGBrightRed.Println("bg_bright_red")
	BGBrightGreen.Println("bg_bright_green")
	BGBrightYellow.Println("bg_bright_yellow")
	BGBrightBlue.Println("bg_bright_blue")
	BGBrightMagenta.Println("bg_bright_magenta")
	BGBrightCyan.Println("bg_bright_cyan")
	BGBrightWhite.Println("bg_bright_white")
	fmt.Println()

	Bold.WithStyle(Italic).Print("bold;italic\n")
	Bold.WithColor(Red).Printf("%s\n", "bold;red")
	Bold.WithBackground(BGWhite).Println("bold;bg_white")
	fmt.Print(Red.WithStyle(Bold).Sprint("red;bold\n"))
	fmt.Print(Red.WithBackground(BGWhite).Sprintf("%s\n", "red;bg_white"))
	fmt.Print(BGWhite.WithStyle(Bold).Sprintln("bg_white;bold"))
	BGWhite.WithColor(Red).Println("bg_white;red")
	Bold.WithColor(Red).WithBackground(BGWhite).Println("bold;red;bg_white")
}
