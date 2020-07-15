package xterminal

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"testing"
)

func TestInitTerminal(t *testing.T) {
	log.SetFlags(0)

	InitTerminal(os.Stdout)
	log.Printf("\x1b[%dm%s\x1b[0m\n", 31, "test")
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", 31, "test")

	color.ForceOpenColor()
	log.Println(color.Red.Render("test"))
	fmt.Println(color.Red.Render("test"))
}
