package xcolor

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestForceColor(t *testing.T) {
	EnableColor()
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

	xtesting.Equal(t, Black.Code(), uint8(30))
	xtesting.Equal(t, Red.Code(), uint8(31))
	xtesting.Equal(t, Green.Code(), uint8(32))
	xtesting.Equal(t, Yellow.Code(), uint8(33))
	xtesting.Equal(t, Blue.Code(), uint8(34))
	xtesting.Equal(t, Magenta.Code(), uint8(35))
	xtesting.Equal(t, Cyan.Code(), uint8(36))
	xtesting.Equal(t, White.Code(), uint8(37))
	xtesting.Equal(t, Default.Code(), uint8(39))

	xtesting.Equal(t, Black.Len(), 9)
	xtesting.Equal(t, Red.Len(), 9)
	xtesting.Equal(t, Green.Len(), 9)
	xtesting.Equal(t, Yellow.Len(), 9)
	xtesting.Equal(t, Blue.Len(), 9)
	xtesting.Equal(t, Magenta.Len(), 9)
	xtesting.Equal(t, Cyan.Len(), 9)
	xtesting.Equal(t, White.Len(), 9)
	xtesting.Equal(t, Default.Len(), 9)
}

func TestPrint(t *testing.T) {
	EnableColor()
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
