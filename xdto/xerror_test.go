package xdto

import (
	"fmt"
	"log"
	"testing"
)

func TestErrorDto(t *testing.T) {
	log.Println(BuildBasicErrorDto(fmt.Errorf("test error"), []string{"test"}))
	log.Println()
	log.Println(BuildErrorDto(fmt.Errorf("test error"), nil, 1, true))
}
