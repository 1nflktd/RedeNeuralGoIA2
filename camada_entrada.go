package main

import (
	"math/rand"
	"time"
)

type CamadaEntrada struct {
	Neuronios [16]Neuronio
	Peso Peso // Entrada -> Intermediaria
	SaidaEsperada int
	Seed *rand.Rand
}

func (c *CamadaEntrada) Init() {
	c.Peso.Init(16, 13)
	c.Seed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (c *CamadaEntrada) AdicionarNeuronio(index int, valor float64) {
	c.Neuronios[index] = Neuronio{ Entrada: valor, Saida: valor }

	if !(c.Peso.Obter(index, 0) > 0) {
		// inicializar peso
		for i := 0; i < 13; i++ {
			c.Peso.Adicionar(index, i, c.Seed.Float64())
		}
	}
}