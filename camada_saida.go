package main

import (
	"strings"
	"strconv"
)

type CamadaSaida struct {
	Neuronios [10]Neuronio
	SaidaEsperada [10]float64
}

func (c *CamadaSaida) SetSaidaEsperada(s string) {
	if len(s) == 10 {
		aBits := strings.Split(s, "")
		for i, bit := range aBits {
			v64, err := strconv.ParseFloat(bit, 64)
			if err == nil {
				c.SaidaEsperada[i] = v64
			}
		}
	}
	// else error
}

func (c *CamadaSaida) GetSaidaEsperadaNeuronio(i int) float64 {
	if (i >= 0 && i <= 9) {
		return c.SaidaEsperada[i]
	}
	// else error
	return 0.0
}

func (c *CamadaSaida) GetSaidaNeuronio(i int) float64 {
	if (i >= 0 && i <= 9) {
		return c.Neuronios[i].GetSaida()
	}
	// else error
	return 0.0
}

func (c *CamadaSaida) SetSaidaNeuronio(i int, s float64) {
	c.Neuronios[i].SetSaida(s)
}

func (c *CamadaSaida) SetErroNeuronio(i int, s float64) {
	c.Neuronios[i].SetErro(s)
}
