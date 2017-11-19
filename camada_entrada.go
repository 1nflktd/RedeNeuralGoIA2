package main

import (
	"math/rand"
	"time"
)

const LimiteInferior = 0
const LimiteSuperior = 100

type CamadaEntrada struct {
	Neuronios [16]Neuronio
	Peso Peso // Entrada -> Intermediaria
}

func (c *CamadaEntrada) Init(nroNeuroniosIntermediarios int) {
	c.Peso.Init(16, nroNeuroniosIntermediarios)
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// inicializar pesos
	for j := 0; j < 16; j++ {
		for i := 0; i < nroNeuroniosIntermediarios; i++ {
			c.Peso.Adicionar(j, i, seed.Float64())
		}
	}
}

func (c *CamadaEntrada) AdicionarNeuronio(index int, valor float64) {
	// normalizar
	vNormalizado := (valor - LimiteInferior) / (LimiteSuperior - LimiteInferior)

	c.Neuronios[index] = Neuronio{ Entrada: vNormalizado, Saida: vNormalizado }
}
