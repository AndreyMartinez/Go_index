package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
Manejo de pruebas con testify a mis helloWorld
**/
func TestToLearnGo(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(add(1, 2), 3, "Probando funciones del hola mundo")
	assert.Equal(nameReturn("hola", "mundo"), "holamundo", "Retornar el hola mundo ")
	assert.NotEqual(varReturnValue(), "holas", "Operación para validar valores diferentes")
}

/**
Manejo de pruebas con testify a mis helloWorld en forma de tabla
**/
func TestForTable(t *testing.T) {
	assert := assert.New(t)

	//Valores de test para función add
	var testValues = []struct {
		valueOne int
		valueTwo int
		result   int
	}{
		{1, 2, 3},
		{1, 3, 4},
		{2, 3, 5},
	}

	for _, test := range testValues {
		assert.Equal(add(test.valueOne, test.valueTwo), test.result)
	}

	//valores de test para name Return

	var nameTestVlues = []struct {
		valueOne string
		valueTwo string
		result   string
	}{
		{"test", "unico", "testunico"},
		{"raphael", "martinez", "raphaelmartinez"},
	}

	for _, nameTest := range nameTestVlues {
		assert.Equal(nameReturn(nameTest.valueOne, nameTest.valueTwo), nameTest.result)
	}

}
