package main

import (
	"strings"
	"strconv"
	"math"
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
			valores := strings.Split(l, ",")
			if len(valores) == 17 {
				for i := 0; i < 16; i++ {
					v64, err := strconv.ParseFloat(valores[i], 64)
					if err != nil {
						r.CamadaEntrada.AdicionarNeuronio(i, v64)
					}
				}
				r.CamadaSaida.SetSaidaEsperada(valores[16])
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
		nI.Saida = r.FuncaoAtivacao(somatorio)
	}

	for iS, nS := range r.CamadaSaida.Neuronios {
		somatorio := 0.0
		for iI, nI := range r.CamadaIntermediaria.Neuronios {
			somatorio += r.CamadaIntermediaria.Peso.Obter(iI, iS) * nI.Saida
		}
		nS.Saida = r.FuncaoAtivacao(somatorio)
	}
}

func (r *RedeNeural) FuncaoAtivacao(somatorio float64) float64 {
	return 1/(1 + math.Exp(-somatorio))
}

func (r *RedeNeural) CalcularErros() {
	for nI, nS := range r.CamadaSaida.Neuronios {
		nS.Erro = nS.Saida * (1 - nS.Saida) * (r.CamadaSaida.GetSaidaEsperadaNeuronio(nI) - nS.Saida)
	}

	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		fatorErro := 0.0
		for iS, nS := range r.CamadaSaida.Neuronios {
			fatorErro += nS.Erro * r.CamadaIntermediaria.Peso.Obter(iI, iS)
		}
		nI.Erro = nI.Saida * (1 - nI.Saida) * fatorErro
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
