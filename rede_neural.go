package main

import (
	"strings"
	"strconv"
	"math"
	"bufio"
	"os"
	"log"
)

type RedeNeural struct {
	CamadaEntrada CamadaEntrada
	CamadaIntermediaria CamadaIntermediaria
	CamadaSaida CamadaSaida
}

func (r *RedeNeural) Init() {
	r.CamadaEntrada.Init()
	r.CamadaIntermediaria.Init()
}

func (r *RedeNeural) Treinar(arq string) int {
	file, errFile := os.Open(arq)
	if errFile != nil {
		log.Fatal(errFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nLinha := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
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

				nLinha++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nLinha
}

func (r *RedeNeural) Testar(arq string) {
	file, errFile := os.Open(arq)
	if errFile != nil {
		log.Fatal(errFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	acertos := 0
	total := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
			valores := strings.Split(l, ",")
			if len(valores) == 17 {
				total++

				for i := 0; i < 16; i++ {
					v64, err := strconv.ParseFloat(valores[i], 64)
					if err == nil {
						r.CamadaEntrada.AdicionarNeuronio(i, v64)
					}
				}
				r.CamadaSaida.SetSaidaEsperada(valores[16])
				r.CalcularSomatorios()

				//log.Printf("%+v\n\n\n", r.CamadaSaida)

				// verificar neuronios ativados se batem com a entrada
				acertou := true
				for i := 0; i < 10; i++ {
					saidaEsperada := r.CamadaSaida.GetSaidaEsperadaNeuronio(i)
					saida := r.CamadaSaida.GetSaidaNeuronio(i)

					if saida < 0.5 {
						saida = 0
					} else {
						saida = 1
					}

					if saidaEsperada != saida {
						acertou = false
						break
					}
				}

				if acertou {
					acertos++
				}
			}
		}
	}

	log.Printf("Total: %d, Acertos: %d, Erros: %d, %% Acertos: %f\n", total, acertos, (total - acertos), ((acertos * 100)/total))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *RedeNeural) CalcularSomatorios() {
	for iI, _ := range r.CamadaIntermediaria.Neuronios {
		somatorio := 0.0
		for iE, nE := range r.CamadaEntrada.Neuronios {
			somatorio += r.CamadaEntrada.Peso.Obter(iE, iI) * nE.GetSaida()
		}
		r.CamadaIntermediaria.SetSaidaNeuronio(iI, r.FuncaoAtivacao(somatorio))
	}

	for iS, _ := range r.CamadaSaida.Neuronios {
		somatorio := 0.0
		for iI, nI := range r.CamadaIntermediaria.Neuronios {
			somatorio += r.CamadaIntermediaria.Peso.Obter(iI, iS) * nI.GetSaida()
		}
		r.CamadaSaida.SetSaidaNeuronio(iS, r.FuncaoAtivacao(somatorio))
	}
}

func (r *RedeNeural) FuncaoAtivacao(somatorio float64) float64 {
	/*
	if somatorio < 0 {
		return 1 - 1/(1 + math.Exp(somatorio))
	}
	*/
	return 1/(1 + math.Exp(-somatorio))
}

func (r *RedeNeural) CalcularErros() {
	for iS, nS := range r.CamadaSaida.Neuronios {
		erro := nS.GetSaida() * (1 - nS.GetSaida()) * (r.CamadaSaida.GetSaidaEsperadaNeuronio(iS) - nS.GetSaida())
		r.CamadaSaida.SetErroNeuronio(iS, erro)
	}

	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		fatorErro := 0.0
		for iS, nS := range r.CamadaSaida.Neuronios {
			fatorErro += nS.GetErro() * r.CamadaIntermediaria.Peso.Obter(iI, iS)
		}
		erro := nI.GetSaida() * (1 - nI.GetSaida()) * fatorErro
		r.CamadaIntermediaria.SetErroNeuronio(iI, erro)
	}
}

func (r *RedeNeural) AjustarPesos() {
	for iE, nE := range r.CamadaEntrada.Neuronios {
		for iI, nI := range r.CamadaIntermediaria.Neuronios {
			novoPeso := r.CamadaEntrada.Peso.Obter(iE, iI) + TaxaAprendizagem * nE.GetSaida() * nI.GetErro()
			r.CamadaEntrada.Peso.Adicionar(iE, iI, novoPeso)
		}
	}

	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		for iS, nS := range r.CamadaSaida.Neuronios {
			novoPeso := r.CamadaIntermediaria.Peso.Obter(iI, iS) + TaxaAprendizagem * nI.GetSaida() * nS.GetErro()
			r.CamadaIntermediaria.Peso.Adicionar(iI, iS, novoPeso)
		}
	}
}
