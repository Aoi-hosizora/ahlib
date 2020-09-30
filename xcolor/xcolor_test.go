package xcolor

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestForceColor(t *testing.T) {
	ForceColor()
	fmt.Println("test")
	fmt.Printf(FullColorTpl+"\n", 32, "test")
	fmt.Printf(FullColorTpl+"\n", 33, "test")
}

func TestDisableColor(t *testing.T) {
	EnableColor()
	Red.Println("test red")
	fmt.Println(Red.Sprint("test red"))
	Yellow.Println("test yellow")
	fmt.Println(Yellow.Sprint("test yellow"))

	DisableColor()
	Red.Println("test red (disable)")
	fmt.Println(Red.Sprint("test red (disable)"))
	Yellow.Println("test yellow (disable)")
	fmt.Println(Yellow.Sprint("test yellow (disable)"))
}

func TestColor(t *testing.T) {
	xtesting.Equal(t, Black.String(), "30")
	xtesting.Equal(t, Red.String(), "31")
	xtesting.Equal(t, Green.String(), "32")
	xtesting.Equal(t, Yellow.String(), "33")
	xtesting.Equal(t, Blue.String(), "34")
	xtesting.Equal(t, Magenta.String(), "35")
	xtesting.Equal(t, Cyan.String(), "36")
	xtesting.Equal(t, White.String(), "37")
	xtesting.Equal(t, Default.String(), "39")

	xtesting.Equal(t, Black.Code(), "30")
	xtesting.Equal(t, Red.Code(), "31")
	xtesting.Equal(t, Green.Code(), "32")
	xtesting.Equal(t, Yellow.Code(), "33")
	xtesting.Equal(t, Blue.Code(), "34")
	xtesting.Equal(t, Magenta.Code(), "35")
	xtesting.Equal(t, Cyan.Code(), "36")
	xtesting.Equal(t, White.Code(), "37")
	xtesting.Equal(t, Default.Code(), "39")
}

func TestPrint(t *testing.T) {
	for _, color := range []Color{
		Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Default,
	} {
		color.Print("test\n")
		color.Println("test")
		color.Printf("%s\n", "test")

		fmt.Print(color.Sprint("test\n"))
		fmt.Print(color.Sprintln("test"))
		fmt.Print(color.Sprintf("%s\n", "test"))
	}
}
