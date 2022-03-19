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
		Bold.APrint(5)
		Bold.APrintf(5, "")
		Bold.APrintln(5)
		Bold.ASprint(5)
		Bold.ASprintf(5, "")
		Bold.ASprintln(5)
		_, _ = Bold.AFprint(5, os.Stdout)
		_, _ = Bold.AFprintf(5, os.Stdout, "")
		_, _ = Bold.AFprintln(5, os.Stdout)

		Red.APrint(-5)
		Red.APrintf(-5, "")
		Red.APrintln(-5)
		Red.ASprint(-5)
		Red.ASprintf(-5, "")
		Red.ASprintln(-5)
		_, _ = Red.AFprint(-5, os.Stdout)
		_, _ = Red.AFprintf(-5, os.Stdout, "")
		_, _ = Red.AFprintln(-5, os.Stdout)

		BGRed.APrint(5)
		BGRed.APrintf(5, "")
		BGRed.APrintln(5)
		BGRed.ASprint(5)
		BGRed.ASprintf(5, "")
		BGRed.ASprintln(5)
		_, _ = BGRed.AFprint(5, os.Stdout)
		_, _ = BGRed.AFprintf(5, os.Stdout, "")
		_, _ = BGRed.AFprintln(5, os.Stdout)

		Bold.WithStyle(Italic).APrint(-5)
		Bold.WithStyle(Italic).APrintf(-5, "")
		Bold.WithStyle(Italic).APrintln(-5)
		Bold.WithStyle(Italic).ASprint(-5)
		Bold.WithStyle(Italic).ASprintf(-5, "")
		Bold.WithStyle(Italic).ASprintln(-5)
		_, _ = Bold.WithStyle(Italic).AFprint(-5, os.Stdout)
		_, _ = Bold.WithStyle(Italic).AFprintf(-5, os.Stdout, "")
		_, _ = Bold.WithStyle(Italic).AFprintln(-5, os.Stdout)
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

func TestNPrint(t *testing.T) {
	for _, alignment := range []int{20, -20} {
		fmt.Println(strings.Repeat("=", int(math.Abs(float64(alignment)))))

		// style
		fmt.Println(">>> style")
		Bold.APrint(alignment, "bold")
		fmt.Println()
		Faint.APrintf(alignment, "%s", "faint")
		fmt.Println()
		Italic.APrintln(alignment, "italic")
		fmt.Print(Underline.ASprint(alignment, "underline"))
		fmt.Println()
		fmt.Print(Reverse.ASprintf(alignment, "%s", "reverse"))
		fmt.Println()
		fmt.Print(Strikethrough.ASprintln(alignment, "strikethrough"))
		_, _ = Bold.AFprint(alignment, os.Stdout, "bold")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Faint.AFprintf(alignment, os.Stdout, "%s", "faint")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Italic.AFprintln(alignment, os.Stdout, "italic")

		// color
		fmt.Println(">>> color")
		Red.APrint(alignment, "red")
		fmt.Println()
		Green.APrintf(alignment, "%s", "green")
		fmt.Println()
		Yellow.APrintln(alignment, "yellow")
		fmt.Print(Blue.ASprint(alignment, "blue"))
		fmt.Println()
		fmt.Print(Magenta.ASprintf(alignment, "%s", "magenta"))
		fmt.Println()
		fmt.Print(Cyan.ASprintln(alignment, "cyan"))
		_, _ = Black.AFprint(alignment, os.Stdout, "black")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = White.AFprintf(alignment, os.Stdout, "%s", "white")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Default.AFprintln(alignment, os.Stdout, "default")
		BrightBlack.APrintln(alignment, "bright_black")
		BrightRed.APrintln(alignment, "bright_red")
		BrightGreen.APrintln(alignment, "bright_green")
		BrightYellow.APrintln(alignment, "bright_yellow")
		BrightBlue.APrintln(alignment, "bright_blue")
		BrightMagenta.APrintln(alignment, "bright_magenta")
		BrightCyan.APrintln(alignment, "bright_cyan")
		BrightWhite.APrintln(alignment, "bright_white")

		// background
		fmt.Println(">>> background")
		BGRed.APrint(alignment, "bg_red")
		fmt.Println()
		BGGreen.APrintf(alignment, "%s", "bg_green")
		fmt.Println()
		BGYellow.APrintln(alignment, "bg_yellow")
		fmt.Print(BGBlue.ASprint(alignment, "bg_blue"))
		fmt.Println()
		fmt.Print(BGMagenta.ASprintf(alignment, "%s", "bg_magenta"))
		fmt.Println()
		fmt.Print(BGCyan.ASprintln(alignment, "bg_cyan"))
		_, _ = BGBlack.AFprint(alignment, os.Stdout, "bg_black")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = BGWhite.AFprintf(alignment, os.Stdout, "%s", "bg_white")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = BGDefault.AFprintln(alignment, os.Stdout, "bg_default")
		BGBrightBlack.APrintln(alignment, "bg_bright_black")
		BGBrightRed.APrintln(alignment, "bg_bright_red")
		BGBrightGreen.APrintln(alignment, "bg_bright_green")
		BGBrightYellow.APrintln(alignment, "bg_bright_yellow")
		BGBrightBlue.APrintln(alignment, "bg_bright_blue")
		BGBrightMagenta.APrintln(alignment, "bg_bright_magenta")
		BGBrightCyan.APrintln(alignment, "bg_bright_cyan")
		BGBrightWhite.APrintln(alignment, "bg_bright_white")

		// mix code
		fmt.Println(">>> mix code")
		Bold.WithStyle(Italic).APrint(alignment, "bold;italic")
		fmt.Println()
		Bold.WithColor(Red).APrintf(alignment, "%s", "bold;red")
		fmt.Println()
		Bold.WithBackground(BGWhite).APrintln(alignment, "bold;bg_white")
		fmt.Print(Red.WithStyle(Bold).ASprint(alignment, "red;bold"))
		fmt.Println()
		fmt.Print(Red.WithBackground(BGWhite).ASprintf(alignment, "%s", "red;bg_white"))
		fmt.Println()
		fmt.Print(BGWhite.WithStyle(Bold).ASprintln(alignment, "bg_white;bold"))
		BGWhite.WithColor(Red).APrintln(alignment, "bg_white;red")
		Bold.WithColor(Red).WithBackground(BGWhite).APrintln(alignment, "bold;red;bg_white")
		_, _ = Bold.WithStyle(Italic).AFprint(alignment, os.Stdout, "bold;italic")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Bold.WithColor(Red).AFprintf(alignment, os.Stdout, "%s", "bold;red")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Bold.WithBackground(BGWhite).AFprintln(alignment, os.Stdout, "bold;bg_white")
	}
}
