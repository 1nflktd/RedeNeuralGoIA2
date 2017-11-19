package main

import (
	"math/rand"
	"time"
)

type CamadaIntermediaria struct {
	Neuronios [13]Neuronio // transformar para variar
	Peso Peso // Intermediaria -> Saida
	Saida float64
	Erro float64
}

func (c *CamadaIntermediaria) Init() {
	c.Peso.Init(13, 10)
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 13; i++ {
		for j := 0; j < 10; j++ {
			c.Peso.Adicionar(i, j, seed.Float64())
		}
	}
}

func (c *CamadaIntermediaria) SetSaidaNeuronio(i int, s float64) {
	c.Neuronios[i].SetSaida(s)
}

func (c *CamadaIntermediaria) SetErroNeuronio(i int, s float64) {
	c.Neuronios[i].SetErro(s)
}
