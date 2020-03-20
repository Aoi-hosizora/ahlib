package xnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderLatency(t *testing.T) {
	assert.Equal(t, RenderLatency(-0.1), "0.0000ns")
	assert.Equal(t, RenderLatency(0), "0.0000ns")
	assert.Equal(t, RenderLatency(999), "999.0000ns")
	assert.Equal(t, RenderLatency(10000), "10.0000µs")
	assert.Equal(t, RenderLatency(1000000), "1.0000ms")
	assert.Equal(t, RenderLatency(10000000000), "10.0000s")
	assert.Equal(t, RenderLatency(59000000000), "59.0000s")
	assert.Equal(t, RenderLatency(60000000000), "1.0000min")
}

func TestRenderByte(t *testing.T) {
	assert.Equal(t, RenderByte(-5), "0B")
	assert.Equal(t, RenderByte(0), "0B")
	assert.Equal(t, RenderByte(1023), "1023B")
	assert.Equal(t, RenderByte(1024), "1.00KB")
	assert.Equal(t, RenderByte(1536), "1.50KB")
	assert.Equal(t, RenderByte(2048), "2.00KB")
	assert.Equal(t, RenderByte(1024*1024), "1.00MB")
	assert.Equal(t, RenderByte(2.51*1024*1024), "2.51MB")
}
