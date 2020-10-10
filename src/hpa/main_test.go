package main

import (
	"testing"
)

func TestMsgCodeEducation(t *testing.T) {
	resultado := msgCodeEducation()
	esperado := "Code.education Rocks!"

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestCalcRaiz(t *testing.T) {
	resultado := calcRaiz()
	esperado := 2.499962871007386e+11

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}
