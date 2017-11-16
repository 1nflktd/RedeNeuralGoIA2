package main

import (
	"strings"
	"strconv"
)

type RedeNeural struct {
	CamadaEntrada CamadaEntrada
	CamadaIntermediaria CamadaIntermediaria
	CamadaSaida CamadaSaida
}

func (r *RedeNeural) Init(dadosTeste string) {
	r.CamadaEntrada.Init()
	r.CamadaIntermediaria.Init()
	//r.CamadaSaida.Init()
	r.IniciarTeste(dadosTeste)
}

func (r *RedeNeural) IniciarTeste(dadosTeste string) {
	linhas := strings.Split(dadosTeste, "\n")
	for _, l := range linhas {
		if l != "" {
			//fmt.Printf("%i: %s\n", i, l)
			valores := strings.Split(l, " ")
			if len(valores) == 17 {
				for i := 0; i < 16; i++ {
					v64, err := strconv.ParseFloat(valores[i], 64)
					if err != nil {
						r.CamadaEntrada.AdicionarNeuronio(i, v64)
					}

				}
				// saida esperada valores[16]
			}
		}
	}
}

func (r *RedeNeural) CalcularSomatorios() {
	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		somatorio := 0.0
		for iE, nE := range r.CamadaEntrada.Neuronios {
			somatorio += r.CamadaEntrada.Peso.Obter(iE, iI) * nE.Saida
		}
		nI.Saida = somatorio
	}

	for iS, nS := range r.CamadaSaida.Neuronios {
		somatorio := 0.0
		for iI, nI := range r.CamadaIntermediaria.Neuronios {
			somatorio += r.CamadaIntermediaria.Peso.Obter(iI, iS) * nI.Saida
		}
		nS.Saida = somatorio
	}
}
