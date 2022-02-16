package xcolor

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"math"
	"os"
	"strings"
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

	Style.unexportedXXX(0)
	Color.unexportedXXX(0)
	Background.unexportedXXX(0)
	MixCode.unexportedXXX(nil)
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

	for _, tc := range []struct {
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

func TestPrintEmpty(t *testing.T) {
	xtesting.NotPanic(t, func() {
		Bold.Print()
		Bold.Printf("")
		Bold.Println()
		Bold.Sprint()
		Bold.Sprintf("")
		Bold.Sprintln()
		_, _ = Bold.Fprint(os.Stdout)
		_, _ = Bold.Fprintf(os.Stdout, "")
		_, _ = Bold.Fprintln(os.Stdout)

		Red.Print()
		Red.Printf("")
		Red.Println()
		Red.Sprint()
		Red.Sprintf("")
		Red.Sprintln()
		_, _ = Red.Fprint(os.Stdout)
		_, _ = Red.Fprintf(os.Stdout, "")
		_, _ = Red.Fprintln(os.Stdout)

		BGRed.Print()
		BGRed.Printf("")
		BGRed.Println()
		BGRed.Sprint()
		BGRed.Sprintf("")
		BGRed.Sprintln()
		_, _ = BGRed.Fprint(os.Stdout)
		_, _ = BGRed.Fprintf(os.Stdout, "")
		_, _ = BGRed.Fprintln(os.Stdout)

		Bold.WithStyle(Italic).Print()
		Bold.WithStyle(Italic).Printf("")
		Bold.WithStyle(Italic).Println()
		Bold.WithStyle(Italic).Sprint()
		Bold.WithStyle(Italic).Sprintf("")
		Bold.WithStyle(Italic).Sprintln()
		_, _ = Bold.WithStyle(Italic).Fprint(os.Stdout)
		_, _ = Bold.WithStyle(Italic).Fprintf(os.Stdout, "")
		_, _ = Bold.WithStyle(Italic).Fprintln(os.Stdout)
	})

	xtesting.NotPanic(t, func() {
		Bold.AlignedPrint(5)
		Bold.AlignedPrintf(5, "")
		Bold.AlignedPrintln(5)
		Bold.AlignedSprint(5)
		Bold.AlignedSprintf(5, "")
		Bold.AlignedSprintln(5)
		_, _ = Bold.AlignedFprint(5, os.Stdout)
		_, _ = Bold.AlignedFprintf(5, os.Stdout, "")
		_, _ = Bold.AlignedFprintln(5, os.Stdout)

		Red.AlignedPrint(-5)
		Red.AlignedPrintf(-5, "")
		Red.AlignedPrintln(-5)
		Red.AlignedSprint(-5)
		Red.AlignedSprintf(-5, "")
		Red.AlignedSprintln(-5)
		_, _ = Red.AlignedFprint(-5, os.Stdout)
		_, _ = Red.AlignedFprintf(-5, os.Stdout, "")
		_, _ = Red.AlignedFprintln(-5, os.Stdout)

		BGRed.AlignedPrint(5)
		BGRed.AlignedPrintf(5, "")
		BGRed.AlignedPrintln(5)
		BGRed.AlignedSprint(5)
		BGRed.AlignedSprintf(5, "")
		BGRed.AlignedSprintln(5)
		_, _ = BGRed.AlignedFprint(5, os.Stdout)
		_, _ = BGRed.AlignedFprintf(5, os.Stdout, "")
		_, _ = BGRed.AlignedFprintln(5, os.Stdout)

		Bold.WithStyle(Italic).AlignedPrint(-5)
		Bold.WithStyle(Italic).AlignedPrintf(-5, "")
		Bold.WithStyle(Italic).AlignedPrintln(-5)
		Bold.WithStyle(Italic).AlignedSprint(-5)
		Bold.WithStyle(Italic).AlignedSprintf(-5, "")
		Bold.WithStyle(Italic).AlignedSprintln(-5)
		_, _ = Bold.WithStyle(Italic).AlignedFprint(-5, os.Stdout)
		_, _ = Bold.WithStyle(Italic).AlignedFprintf(-5, os.Stdout, "")
		_, _ = Bold.WithStyle(Italic).AlignedFprintln(-5, os.Stdout)
	})
}

func TestPrint(t *testing.T) {
	// style
	fmt.Println(">>> style")
	Bold.Print("bold")
	fmt.Println()
	Faint.Printf("%s", "faint")
	fmt.Println()
	Italic.Println("italic")
	fmt.Print(Underline.Sprint("underline"))
	fmt.Println()
	fmt.Print(Reverse.Sprintf("%s", "reverse"))
	fmt.Println()
	fmt.Print(Strikethrough.Sprintln("strikethrough"))
	_, _ = Bold.Fprint(os.Stdout, "bold")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = Faint.Fprintf(os.Stdout, "%s", "faint")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = Italic.Fprintln(os.Stdout, "italic")

	// color
	fmt.Println(">>> color")
	Red.Print("red")
	fmt.Println()
	Green.Printf("%s", "green")
	fmt.Println()
	Yellow.Println("yellow")
	fmt.Print(Blue.Sprint("blue"))
	fmt.Println()
	fmt.Print(Magenta.Sprintf("%s", "magenta"))
	fmt.Println()
	fmt.Print(Cyan.Sprintln("cyan"))
	_, _ = Black.Fprint(os.Stdout, "black")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = White.Fprintf(os.Stdout, "%s", "white")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = Default.Fprintln(os.Stdout, "default")
	BrightBlack.Println("bright_black")
	BrightRed.Println("bright_red")
	BrightGreen.Println("bright_green")
	BrightYellow.Println("bright_yellow")
	BrightBlue.Println("bright_blue")
	BrightMagenta.Println("bright_magenta")
	BrightCyan.Println("bright_cyan")
	BrightWhite.Println("bright_white")

	// background
	fmt.Println(">>> background")
	BGRed.Print("bg_red")
	fmt.Println()
	BGGreen.Printf("%s", "bg_green")
	fmt.Println()
	BGYellow.Println("bg_yellow")
	fmt.Print(BGBlue.Sprint("bg_blue"))
	fmt.Println()
	fmt.Print(BGMagenta.Sprintf("%s", "bg_magenta"))
	fmt.Println()
	fmt.Print(BGCyan.Sprintln("bg_cyan"))
	_, _ = BGBlack.Fprint(os.Stdout, "bg_black")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = BGWhite.Fprintf(os.Stdout, "%s", "bg_white")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = BGDefault.Fprintln(os.Stdout, "bg_default")
	BGBrightBlack.Println("bg_bright_black")
	BGBrightRed.Println("bg_bright_red")
	BGBrightGreen.Println("bg_bright_green")
	BGBrightYellow.Println("bg_bright_yellow")
	BGBrightBlue.Println("bg_bright_blue")
	BGBrightMagenta.Println("bg_bright_magenta")
	BGBrightCyan.Println("bg_bright_cyan")
	BGBrightWhite.Println("bg_bright_white")

	// mix code
	fmt.Println(">>> mix code")
	Bold.WithStyle(Italic).Print("bold;italic")
	fmt.Println()
	Bold.WithColor(Red).Printf("%s", "bold;red")
	fmt.Println()
	Bold.WithBackground(BGWhite).Println("bold;bg_white")
	fmt.Print(Red.WithStyle(Bold).Sprint("red;bold"))
	fmt.Println()
	fmt.Print(Red.WithBackground(BGWhite).Sprintf("%s", "red;bg_white"))
	fmt.Println()
	fmt.Print(BGWhite.WithStyle(Bold).Sprintln("bg_white;bold"))
	BGWhite.WithColor(Red).Println("bg_white;red")
	Bold.WithColor(Red).WithBackground(BGWhite).Println("bold;red;bg_white")
	_, _ = Bold.WithStyle(Italic).Fprint(os.Stdout, "bold;italic")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = Bold.WithColor(Red).Fprintf(os.Stdout, "%s", "bold;red")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = Bold.WithBackground(BGWhite).Fprintln(os.Stdout, "bold;bg_white")
}

func TestAlignedPrint(t *testing.T) {
	for _, alignment := range []int{20, -20} {
		fmt.Println(strings.Repeat("=", int(math.Abs(float64(alignment)))))

		// style
		fmt.Println(">>> style")
		Bold.AlignedPrint(alignment, "bold")
		fmt.Println()
		Faint.AlignedPrintf(alignment, "%s", "faint")
		fmt.Println()
		Italic.AlignedPrintln(alignment, "italic")
		fmt.Print(Underline.AlignedSprint(alignment, "underline"))
		fmt.Println()
		fmt.Print(Reverse.AlignedSprintf(alignment, "%s", "reverse"))
		fmt.Println()
		fmt.Print(Strikethrough.AlignedSprintln(alignment, "strikethrough"))
		_, _ = Bold.AlignedFprint(alignment, os.Stdout, "bold")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Faint.AlignedFprintf(alignment, os.Stdout, "%s", "faint")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Italic.AlignedFprintln(alignment, os.Stdout, "italic")

		// color
		fmt.Println(">>> color")
		Red.AlignedPrint(alignment, "red")
		fmt.Println()
		Green.AlignedPrintf(alignment, "%s", "green")
		fmt.Println()
		Yellow.AlignedPrintln(alignment, "yellow")
		fmt.Print(Blue.AlignedSprint(alignment, "blue"))
		fmt.Println()
		fmt.Print(Magenta.AlignedSprintf(alignment, "%s", "magenta"))
		fmt.Println()
		fmt.Print(Cyan.AlignedSprintln(alignment, "cyan"))
		_, _ = Black.AlignedFprint(alignment, os.Stdout, "black")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = White.AlignedFprintf(alignment, os.Stdout, "%s", "white")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Default.AlignedFprintln(alignment, os.Stdout, "default")
		BrightBlack.AlignedPrintln(alignment, "bright_black")
		BrightRed.AlignedPrintln(alignment, "bright_red")
		BrightGreen.AlignedPrintln(alignment, "bright_green")
		BrightYellow.AlignedPrintln(alignment, "bright_yellow")
		BrightBlue.AlignedPrintln(alignment, "bright_blue")
		BrightMagenta.AlignedPrintln(alignment, "bright_magenta")
		BrightCyan.AlignedPrintln(alignment, "bright_cyan")
		BrightWhite.AlignedPrintln(alignment, "bright_white")

		// background
		fmt.Println(">>> background")
		BGRed.AlignedPrint(alignment, "bg_red")
		fmt.Println()
		BGGreen.AlignedPrintf(alignment, "%s", "bg_green")
		fmt.Println()
		BGYellow.AlignedPrintln(alignment, "bg_yellow")
		fmt.Print(BGBlue.AlignedSprint(alignment, "bg_blue"))
		fmt.Println()
		fmt.Print(BGMagenta.AlignedSprintf(alignment, "%s", "bg_magenta"))
		fmt.Println()
		fmt.Print(BGCyan.AlignedSprintln(alignment, "bg_cyan"))
		_, _ = BGBlack.AlignedFprint(alignment, os.Stdout, "bg_black")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = BGWhite.AlignedFprintf(alignment, os.Stdout, "%s", "bg_white")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = BGDefault.AlignedFprintln(alignment, os.Stdout, "bg_default")
		BGBrightBlack.AlignedPrintln(alignment, "bg_bright_black")
		BGBrightRed.AlignedPrintln(alignment, "bg_bright_red")
		BGBrightGreen.AlignedPrintln(alignment, "bg_bright_green")
		BGBrightYellow.AlignedPrintln(alignment, "bg_bright_yellow")
		BGBrightBlue.AlignedPrintln(alignment, "bg_bright_blue")
		BGBrightMagenta.AlignedPrintln(alignment, "bg_bright_magenta")
		BGBrightCyan.AlignedPrintln(alignment, "bg_bright_cyan")
		BGBrightWhite.AlignedPrintln(alignment, "bg_bright_white")

		// mix code
		fmt.Println(">>> mix code")
		Bold.WithStyle(Italic).AlignedPrint(alignment, "bold;italic")
		fmt.Println()
		Bold.WithColor(Red).AlignedPrintf(alignment, "%s", "bold;red")
		fmt.Println()
		Bold.WithBackground(BGWhite).AlignedPrintln(alignment, "bold;bg_white")
		fmt.Print(Red.WithStyle(Bold).AlignedSprint(alignment, "red;bold"))
		fmt.Println()
		fmt.Print(Red.WithBackground(BGWhite).AlignedSprintf(alignment, "%s", "red;bg_white"))
		fmt.Println()
		fmt.Print(BGWhite.WithStyle(Bold).AlignedSprintln(alignment, "bg_white;bold"))
		BGWhite.WithColor(Red).AlignedPrintln(alignment, "bg_white;red")
		Bold.WithColor(Red).WithBackground(BGWhite).AlignedPrintln(alignment, "bold;red;bg_white")
		_, _ = Bold.WithStyle(Italic).AlignedFprint(alignment, os.Stdout, "bold;italic")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Bold.WithColor(Red).AlignedFprintf(alignment, os.Stdout, "%s", "bold;red")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Bold.WithBackground(BGWhite).AlignedFprintln(alignment, os.Stdout, "bold;bg_white")
	}
}
