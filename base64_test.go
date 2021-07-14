package xbase64

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRawBase64StdEncode(t *testing.T) {
	base64Bytes := RawStdEncode([]byte("离离原上草，一岁一枯荣。野火烧不尽，春风吹又生。远芳侵古道，晴翠接荒城。又送王孙去，萋萋满别情。"))
	assert.Equal(t,base64Bytes, []byte("56a756a75Y6f5LiK6I2J77yM5LiA5bKB5LiA5p6v6I2j44CC6YeO54Gr54On5LiN5bC977yM5pil6aOO5ZC55Y+I55Sf44CC6L+c6Iqz5L615Y+k6YGT77yM5pm057+g5o6l6I2S5Z+O44CC5Y+I6YCB546L5a2Z5Y6777yM6JCL6JCL5ruh5Yir5oOF44CC"))
}
func TestRawBase64StdDecode(t *testing.T) {
	base64Bytes, err := RawStdDecode([]byte("56a756a75Y6f5LiK6I2J77yM5LiA5bKB5LiA5p6v6I2j44CC6YeO54Gr54On5LiN5bC977yM5pil6aOO5ZC55Y+I55Sf44CC6L+c6Iqz5L615Y+k6YGT77yM5pm057+g5o6l6I2S5Z+O44CC5Y+I6YCB546L5a2Z5Y6777yM6JCL6JCL5ruh5Yir5oOF44CC"))
	assert.NoError(t, err)
	assert.Equal(t, base64Bytes, []byte("离离原上草，一岁一枯荣。野火烧不尽，春风吹又生。远芳侵古道，晴翠接荒城。又送王孙去，萋萋满别情。"))
}
func TestRawBase64URLEncode(t *testing.T) {
	base64Bytes := RawURLEncode([]byte("离离原上草，一岁一枯荣。野火烧不尽，春风吹又生。远芳侵古道，晴翠接荒城。又送王孙去，萋萋满别情。"))
	assert.Equal(t,base64Bytes, []byte("56a756a75Y6f5LiK6I2J77yM5LiA5bKB5LiA5p6v6I2j44CC6YeO54Gr54On5LiN5bC977yM5pil6aOO5ZC55Y-I55Sf44CC6L-c6Iqz5L615Y-k6YGT77yM5pm057-g5o6l6I2S5Z-O44CC5Y-I6YCB546L5a2Z5Y6777yM6JCL6JCL5ruh5Yir5oOF44CC"))
}
func TestRawBase64URLDecode(t *testing.T) {
	base64Bytes, err := RawURLDecode([]byte("56a756a75Y6f5LiK6I2J77yM5LiA5bKB5LiA5p6v6I2j44CC6YeO54Gr54On5LiN5bC977yM5pil6aOO5ZC55Y-I55Sf44CC6L-c6Iqz5L615Y-k6YGT77yM5pm057-g5o6l6I2S5Z-O44CC5Y-I6YCB546L5a2Z5Y6777yM6JCL6JCL5ruh5Yir5oOF44CC"))
	assert.NoError(t, err)
	assert.Equal(t, base64Bytes, []byte("离离原上草，一岁一枯荣。野火烧不尽，春风吹又生。远芳侵古道，晴翠接荒城。又送王孙去，萋萋满别情。"))
}