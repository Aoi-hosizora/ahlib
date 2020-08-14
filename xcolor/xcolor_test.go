package xcolor

import (
	"fmt"
	"log"
	"testing"
)

func TestForceColor(t *testing.T) {
	ForceColor()
	fmt.Println("test")
	fmt.Printf(FullColorTpl, 32, "test")
	fmt.Println()
	fmt.Printf(FullColorTpl, 33, "test")
	fmt.Println()
}

func TestColor(t *testing.T) {
	Black.Println("test")
	Red.Println("test")
	Green.Println("test")
	Yellow.Println("test")
	Blue.Println("test")
	Magenta.Println("test")
	Cyan.Println("test")
	White.Println("test")
	Default.Println("test")

	log.Println(Black.Sprintf("test"))
	log.Println(Red.Sprintf("test"))
	log.Println(Green.Sprintf("test"))
	log.Println(Yellow.Sprintf("test"))
	log.Println(Blue.Sprintf("test"))
	log.Println(Magenta.Sprintf("test"))
	log.Println(Cyan.Sprintf("test"))
	log.Println(White.Sprintf("test"))
	log.Println(Default.Sprintf("test"))

	fmt.Println(Black.Sprintf("test"))
	fmt.Println(Red.Sprintf("test"))
	fmt.Println(Green.Sprintf("test"))
	fmt.Println(Yellow.Sprintf("test"))
	fmt.Println(Blue.Sprintf("test"))
	fmt.Println(Magenta.Sprintf("test"))
	fmt.Println(Cyan.Sprintf("test"))
	fmt.Println(White.Sprintf("test"))
	fmt.Println(Default.Sprintf("test"))
}

func TestDisableColor(t *testing.T) {
	DisableColor()
	Red.Println("test")
	Yellow.Println("test")

	EnableColor()
	Red.Println("test")
	Yellow.Println("test")
}
