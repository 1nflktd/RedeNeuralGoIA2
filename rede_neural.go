package main

import (
	"strings"
	"strconv"
	"math"
//	"log"
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
	r.Treinar(dadosTeste)
}

func (r *RedeNeural) Treinar(dadosTeste string) {
	linhas := strings.Split(dadosTeste, "\n")
	for _, l := range linhas {
		if l != "" {
			//fmt.Printf("%i: %s\n", i, l)
			valores := strings.Split(l, ",")
			if len(valores) == 17 {
				for i := 0; i < 16; i++ {
					v64, err := strconv.ParseFloat(valores[i], 64)
					if err == nil {
						r.CamadaEntrada.AdicionarNeuronio(i, v64)
					}
				}
				r.CamadaSaida.SetSaidaEsperada(valores[16])
				r.CalcularSomatorios()
				r.CalcularErros()
				r.AjustarPesos()
			}
		}
	}
}

func (r *RedeNeural) CalcularSomatorios() {
	for iI, _ := range r.CamadaIntermediaria.Neuronios {
		somatorio := 0.0
		for iE, nE := range r.CamadaEntrada.Neuronios {
			somatorio += r.CamadaEntrada.Peso.Obter(iE, iI) * nE.Saida
		}
		r.CamadaIntermediaria.SetSaidaNeuronio(iI, r.FuncaoAtivacao(somatorio))
	}

	for iS, _ := range r.CamadaSaida.Neuronios {
		somatorio := 0.0
		for iI, nI := range r.CamadaIntermediaria.Neuronios {
			somatorio += r.CamadaIntermediaria.Peso.Obter(iI, iS) * nI.Saida
		}
		r.CamadaSaida.SetSaidaNeuronio(iS, r.FuncaoAtivacao(somatorio))
	}
}

func (r *RedeNeural) FuncaoAtivacao(somatorio float64) float64 {
	if somatorio < 0 {
		return 1 - 1/(1 + math.Exp(somatorio))
	}
	return 1/(1 + math.Exp(-somatorio))
}

func (r *RedeNeural) CalcularErros() {
	for iS, nS := range r.CamadaSaida.Neuronios {
		erro := nS.Saida * (1 - nS.Saida) * (r.CamadaSaida.GetSaidaEsperadaNeuronio(iS) - nS.Saida)
		r.CamadaSaida.SetErroNeuronio(iS, erro)
	}

	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		fatorErro := 0.0
		for iS, nS := range r.CamadaSaida.Neuronios {
			fatorErro += nS.Erro * r.CamadaIntermediaria.Peso.Obter(iI, iS)
		}
		erro := nI.Saida * (1 - nI.Saida) * fatorErro
		r.CamadaIntermediaria.SetErroNeuronio(iI, erro)
	}
}

func (r *RedeNeural) AjustarPesos() {
	for iE, nE := range r.CamadaEntrada.Neuronios {
		for iI, nI := range r.CamadaIntermediaria.Neuronios {
			novoPeso := r.CamadaEntrada.Peso.Obter(iE, iI) + TaxaAprendizagem * nE.Saida * nI.Erro
			r.CamadaEntrada.Peso.Adicionar(iE, iI, novoPeso)
		}
	}

	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		for iS, nS := range r.CamadaSaida.Neuronios {
			novoPeso := r.CamadaIntermediaria.Peso.Obter(iI, iS) + TaxaAprendizagem * nI.Saida * nS.Erro
			r.CamadaIntermediaria.Peso.Adicionar(iI, iS, novoPeso)
		}
	}
}
