package main

import "fmt"

type Aceleravel interface {
	Acelera()
}

type Moto struct {
	Modelo string
}

func (m *Moto) Acelera() {
	fmt.Printf("vrummmmm\n")
}

type Carro struct {
	Modelo string
}

func (c *Carro) Acelera() {
	fmt.Printf("acelerando um %s...\n", c.Modelo)
}

func aceleraQualquerCoisa(items []Aceleravel) {
	for _, aceleravel := range items {
		aceleravel.Acelera()
	}
}

func main() {
	c := Carro{Modelo: "Gol"}
	m := Moto{Modelo: "Harley"}
	aceleraQualquerCoisa([]Aceleravel{&c, &m})
}
