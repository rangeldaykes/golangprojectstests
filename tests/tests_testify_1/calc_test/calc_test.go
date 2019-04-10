package calc_test

import (
	"testing"
	"tests_testify_1/calc"

	"github.com/stretchr/testify/assert"
)

// TestCalculate_1 is tes
func TestCalculate_1(t *testing.T) {

	if calc.Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculate_2(t *testing.T) {

	assert.Equal(t, calc.Calculate(2), 4)
}

func TestCalculate_Table_1(t *testing.T) {

	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(calc.Calculate(test.input), test.expected)
	}
}
