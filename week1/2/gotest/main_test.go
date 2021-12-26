package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func add(a, b int) int {

// 	return a + b
// }

func TestAdd(t *testing.T) {
	var res int
	t.Log("starting ...")
	res = add(1, 2)
	assert.Equal(t, res, 3)
}
