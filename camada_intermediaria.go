package main

import (
	"math/rand"
	"time"
)

type CamadaIntermediaria struct {
	Neuronios []Neuronio
	Peso Peso // Intermediaria -> Saida
	Saida float64
	Erro float64
}

func (c *CamadaIntermediaria) Init(nroNeuroniosIntermediarios int) {
	c.Peso.Init(nroNeuroniosIntermediarios, 10)

	c.Neuronios = make([]Neuronio, nroNeuroniosIntermediarios)

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nroNeuroniosIntermediarios; i++ {
		for j := 0; j < 10; j++ {
			c.Peso.Adicionar(i, j, seed.Float64())
		}
	}
}

func (c *CamadaIntermediaria) SetSaidaNeuronio(i int, s float64) {
	c.Neuronios[i].Saida = s
}

func (c *CamadaIntermediaria) SetErroNeuronio(i int, s float64) {
	c.Neuronios[i].Erro = s
}
