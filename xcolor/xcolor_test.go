package xcolor

import (
	"fmt"
	"github.com/gookit/color"
	"testing"
)

func TestForceColor(t *testing.T) {
	ForceColor()
	color.Red.Println("test")
	fmt.Println(color.Red.Sprint("test"))
}
