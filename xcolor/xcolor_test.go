package xcolor

import (
	"fmt"
	"testing"
)

func TestColorCode(t *testing.T) {
	fmt.Printf("%-8s: %s ===\n", "Black", Black.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Red", Red.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Green", Green.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Yellow", Yellow.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Purple", Purple.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Magenta", Magenta.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Teal", Teal.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "White", White.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Info", Info.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Warn", Warn.Paint("abc"))
	fmt.Printf("%-8s: %s ===\n", "Fata", Fata.Paint("abc"))

	fmt.Println()

	fmt.Printf("%-8s: %s ===\n", "Black", Black.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Red", Red.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Green", Green.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Yellow", Yellow.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Purple", Purple.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Magenta", Magenta.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Teal", Teal.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "White", White.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Info", Info.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Warn", Warn.PaintAlign(10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Fata", Fata.PaintAlign(10, "abc"))

	fmt.Println()

	fmt.Printf("%-8s: %s ===\n", "Black", Black.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Red", Red.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Green", Green.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Yellow", Yellow.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Purple", Purple.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Magenta", Magenta.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Teal", Teal.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "White", White.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Info", Info.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Warn", Warn.PaintAlign(-10, "abc"))
	fmt.Printf("%-8s: %s ===\n", "Fata", Fata.PaintAlign(-10, "abc"))
}
