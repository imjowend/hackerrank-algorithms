package main

import "fmt"

type Kangaroo struct {
	initialPosition int32
	jumpLength      int32
}

func main() {
	kangaroo1 := Kangaroo{initialPosition: 0, jumpLength: 3}
	kangaroo2 := Kangaroo{initialPosition: 4, jumpLength: 2}

	fmt.Println(areSynchronized(kangaroo1.initialPosition, kangaroo1.jumpLength, kangaroo2.initialPosition, kangaroo2.jumpLength))
}

func areSynchronized(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if (x1 > x2 && v1 > v2) || (x2 > x1 && v2 > v1) {
		return "NO"
	}

	// Mientras El de MENOR Posicion Inicial esté por atrás del de MAYOR Posición Inicial
	// SUMAR a las posiciones la Longitud del Salto y Evaluar que sean IGUALES
	// si son Iguales, devolver YES, si el de MENOR Posicion Inicial pasa al de MAYOR Posición Inicial
	// Terminar el Loop indicando que NO
	var lowerInitialPosition int32
	var upperInitialPosition int32
	var lowerJumpLength int32
	var upperJumpLength int32
	if x1 >= x2 {
		lowerInitialPosition = x2
		lowerJumpLength = v2
		upperInitialPosition = x1
		upperJumpLength = v1
	} else {
		lowerInitialPosition = x1
		lowerJumpLength = v1
		upperInitialPosition = x2
		upperJumpLength = v2
	}
	fmt.Println(lowerInitialPosition, upperInitialPosition)
	var lowerKangaroo int64
	var upperKangaroo int64
	for true {

	}
	return "YES"
}
