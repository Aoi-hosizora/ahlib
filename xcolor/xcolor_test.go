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
		Bold.NPrint(5)
		Bold.NPrintf(5, "")
		Bold.NPrintln(5)
		Bold.NSprint(5)
		Bold.NSprintf(5, "")
		Bold.NSprintln(5)
		_, _ = Bold.NFprint(5, os.Stdout)
		_, _ = Bold.NFprintf(5, os.Stdout, "")
		_, _ = Bold.NFprintln(5, os.Stdout)

		Red.NPrint(-5)
		Red.NPrintf(-5, "")
		Red.NPrintln(-5)
		Red.NSprint(-5)
		Red.NSprintf(-5, "")
		Red.NSprintln(-5)
		_, _ = Red.NFprint(-5, os.Stdout)
		_, _ = Red.NFprintf(-5, os.Stdout, "")
		_, _ = Red.NFprintln(-5, os.Stdout)

		BGRed.NPrint(5)
		BGRed.NPrintf(5, "")
		BGRed.NPrintln(5)
		BGRed.NSprint(5)
		BGRed.NSprintf(5, "")
		BGRed.NSprintln(5)
		_, _ = BGRed.NFprint(5, os.Stdout)
		_, _ = BGRed.NFprintf(5, os.Stdout, "")
		_, _ = BGRed.NFprintln(5, os.Stdout)

		Bold.WithStyle(Italic).NPrint(-5)
		Bold.WithStyle(Italic).NPrintf(-5, "")
		Bold.WithStyle(Italic).NPrintln(-5)
		Bold.WithStyle(Italic).NSprint(-5)
		Bold.WithStyle(Italic).NSprintf(-5, "")
		Bold.WithStyle(Italic).NSprintln(-5)
		_, _ = Bold.WithStyle(Italic).NFprint(-5, os.Stdout)
		_, _ = Bold.WithStyle(Italic).NFprintf(-5, os.Stdout, "")
		_, _ = Bold.WithStyle(Italic).NFprintln(-5, os.Stdout)
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
		Bold.NPrint(alignment, "bold")
		fmt.Println()
		Faint.NPrintf(alignment, "%s", "faint")
		fmt.Println()
		Italic.NPrintln(alignment, "italic")
		fmt.Print(Underline.NSprint(alignment, "underline"))
		fmt.Println()
		fmt.Print(Reverse.NSprintf(alignment, "%s", "reverse"))
		fmt.Println()
		fmt.Print(Strikethrough.NSprintln(alignment, "strikethrough"))
		_, _ = Bold.NFprint(alignment, os.Stdout, "bold")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Faint.NFprintf(alignment, os.Stdout, "%s", "faint")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Italic.NFprintln(alignment, os.Stdout, "italic")

		// color
		fmt.Println(">>> color")
		Red.NPrint(alignment, "red")
		fmt.Println()
		Green.NPrintf(alignment, "%s", "green")
		fmt.Println()
		Yellow.NPrintln(alignment, "yellow")
		fmt.Print(Blue.NSprint(alignment, "blue"))
		fmt.Println()
		fmt.Print(Magenta.NSprintf(alignment, "%s", "magenta"))
		fmt.Println()
		fmt.Print(Cyan.NSprintln(alignment, "cyan"))
		_, _ = Black.NFprint(alignment, os.Stdout, "black")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = White.NFprintf(alignment, os.Stdout, "%s", "white")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Default.NFprintln(alignment, os.Stdout, "default")
		BrightBlack.NPrintln(alignment, "bright_black")
		BrightRed.NPrintln(alignment, "bright_red")
		BrightGreen.NPrintln(alignment, "bright_green")
		BrightYellow.NPrintln(alignment, "bright_yellow")
		BrightBlue.NPrintln(alignment, "bright_blue")
		BrightMagenta.NPrintln(alignment, "bright_magenta")
		BrightCyan.NPrintln(alignment, "bright_cyan")
		BrightWhite.NPrintln(alignment, "bright_white")

		// background
		fmt.Println(">>> background")
		BGRed.NPrint(alignment, "bg_red")
		fmt.Println()
		BGGreen.NPrintf(alignment, "%s", "bg_green")
		fmt.Println()
		BGYellow.NPrintln(alignment, "bg_yellow")
		fmt.Print(BGBlue.NSprint(alignment, "bg_blue"))
		fmt.Println()
		fmt.Print(BGMagenta.NSprintf(alignment, "%s", "bg_magenta"))
		fmt.Println()
		fmt.Print(BGCyan.NSprintln(alignment, "bg_cyan"))
		_, _ = BGBlack.NFprint(alignment, os.Stdout, "bg_black")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = BGWhite.NFprintf(alignment, os.Stdout, "%s", "bg_white")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = BGDefault.NFprintln(alignment, os.Stdout, "bg_default")
		BGBrightBlack.NPrintln(alignment, "bg_bright_black")
		BGBrightRed.NPrintln(alignment, "bg_bright_red")
		BGBrightGreen.NPrintln(alignment, "bg_bright_green")
		BGBrightYellow.NPrintln(alignment, "bg_bright_yellow")
		BGBrightBlue.NPrintln(alignment, "bg_bright_blue")
		BGBrightMagenta.NPrintln(alignment, "bg_bright_magenta")
		BGBrightCyan.NPrintln(alignment, "bg_bright_cyan")
		BGBrightWhite.NPrintln(alignment, "bg_bright_white")

		// mix code
		fmt.Println(">>> mix code")
		Bold.WithStyle(Italic).NPrint(alignment, "bold;italic")
		fmt.Println()
		Bold.WithColor(Red).NPrintf(alignment, "%s", "bold;red")
		fmt.Println()
		Bold.WithBackground(BGWhite).NPrintln(alignment, "bold;bg_white")
		fmt.Print(Red.WithStyle(Bold).NSprint(alignment, "red;bold"))
		fmt.Println()
		fmt.Print(Red.WithBackground(BGWhite).NSprintf(alignment, "%s", "red;bg_white"))
		fmt.Println()
		fmt.Print(BGWhite.WithStyle(Bold).NSprintln(alignment, "bg_white;bold"))
		BGWhite.WithColor(Red).NPrintln(alignment, "bg_white;red")
		Bold.WithColor(Red).WithBackground(BGWhite).NPrintln(alignment, "bold;red;bg_white")
		_, _ = Bold.WithStyle(Italic).NFprint(alignment, os.Stdout, "bold;italic")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Bold.WithColor(Red).NFprintf(alignment, os.Stdout, "%s", "bold;red")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = Bold.WithBackground(BGWhite).NFprintln(alignment, os.Stdout, "bold;bg_white")
	}
}
