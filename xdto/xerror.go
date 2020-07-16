package xdto

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xruntime"
	"time"
)

// An error response model for fiber and gin
// Request: Need to dump gin.Request or fiber.Fasthttp
// Filename... need for runtime stack, need to provide skip
type ErrorDto struct {
	Time    string   `json:"time"`
	Type    string   `json:"type"`
	Detail  string   `json:"detail"`
	Request []string `json:"request"`

	Filename  string `json:"filename,omitempty"`
	Funcname  string `json:"funcname,omitempty"`
	LineIndex int    `json:"line_index,omitempty"`
	Line      string `json:"line,omitempty"`
}

// Build a basic dto (only include time, type, detail, request)
// noinspection GoUnusedExportedFunction
func BuildBasicErrorDto(err interface{}, requests []string) *ErrorDto {
	return BuildErrorDto(err, requests, -1, false)
}

// Build a complete dto (also include runtime parameters)
// noinspection GoUnusedExportedFunction
func BuildErrorDto(err interface{}, requests []string, skip int, print bool) *ErrorDto {
	now := time.Now().Format(time.RFC3339)
	errType := fmt.Sprintf("%T", err)
	errDetail := fmt.Sprintf("%v", err)
	if e, ok := err.(error); ok {
		errDetail = e.Error()
	}
	if requests == nil {
		requests = []string{}
	}

	dto := &ErrorDto{
		Time:    now,
		Type:    errType,
		Detail:  errDetail,
		Request: requests,
	}

	if skip >= 0 {
		var stacks []*xruntime.Stack
		stacks, dto.Filename, dto.Funcname, dto.LineIndex, dto.Line = xruntime.GetStackWithInfo(skip)
		if print {
			fmt.Println()
			xruntime.PrintStacksRed(stacks)
			fmt.Println()
		}
	}

	return dto
}
