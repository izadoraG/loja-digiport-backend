package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	t.Run("Should sum two numbers", func(t *testing.T) {
		// arrange
		num1 := 4
		num2 := 8

		resultadoEsperado := 12

		// act
		resultado := soma(num1, num2)

		// assert
		assert.Equal(t, resultadoEsperado, resultado)
	})
}

func TestValidaNome(t *testing.T) {
	t.Run("Should return an error when name is empty", func(t *testing.T) {
		// arrange <=> given
		name := ""

		expectedErrorMessage := "nome não preenchido"

		// act <=> when
		err := validaNome(name)

		// assert <=> then
		assert.Error(t, err)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})

	t.Run("Should not return an error when name is filled", func(t *testing.T) {
		// arrange
		name := "product"

		// act
		err := validaNome(name)

		// assert
		assert.Nil(t, err)
	})
}
