package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrayToMinGray(t *testing.T) {
	t.Run("test 0 => 0", func(t *testing.T) {
		xExpected, xRes := 0, 0

		xRes = int(grayToMinGray(uint8(0)))
		assert.Equal(t, xExpected, xRes)
	})

	t.Run("test 255 => 8", func(t *testing.T) {
		xExpected, xRes := 7, 0

		xRes = int(grayToMinGray(uint8(255)))
		assert.Equal(t, xExpected, xRes)
	})
}
