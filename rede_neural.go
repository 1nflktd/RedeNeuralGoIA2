package main

import (
	"strings"
	"strconv"
	"math"
	"bufio"
	"os"
	"log"
	"fmt"
)

type RedeNeural struct {
	CamadaEntrada CamadaEntrada
	CamadaIntermediaria CamadaIntermediaria
	CamadaSaida CamadaSaida
	TaxaAprendizagem float64
}

func (r *RedeNeural) Init(taxaAprendizagem float64, nroNeuroniosIntermediarios int) {
	r.CamadaEntrada.Init(nroNeuroniosIntermediarios)
	r.CamadaIntermediaria.Init(nroNeuroniosIntermediarios)
	r.TaxaAprendizagem = taxaAprendizagem
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
	matrizConfusao := [10][10]int{}
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

				// verificar neuronios ativados se batem com a entrada
				saidaNeuronio := [10]float64{}
				max := 0.0
				neuronioAtivado := 0
				for i := 0; i < 10; i++ {
					saidaNeuronio[i] = r.CamadaSaida.GetSaidaNeuronio(i)
					if saidaNeuronio[i] > max {
						max = saidaNeuronio[i]
						neuronioAtivado = i
					}
				}

				saidaNumericaEsperada := r.ObterSaidaNumerica(r.CamadaSaida.SaidaEsperada)
				if saidaNumericaEsperada == neuronioAtivado {
					acertos++
				}

				matrizConfusao[saidaNumericaEsperada][neuronioAtivado]++
			}
		}
	}

	percentualAcertos := (float64(acertos) * 100.0) / float64(total)
	log.Printf("Total: %d, Acertos: %d, Erros: %d, %% Acertos: %f\n", total, acertos, (total - acertos), percentualAcertos)
	log.Printf("Matriz confusao\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d\t", matrizConfusao[i][j])
		}
		fmt.Printf("\n")
	}

	r.CalcularEstatisticas(matrizConfusao)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *RedeNeural) CalcularEstatisticas(matrizConfusao [10][10]int) {
	// vp = 0,0 1,1 2,2
	// vn = E diagonal - vp
	// fp = E coluna  - vp
	// fn = E linha - vp

	diagonal := 0
	coluna := [10]int{}
	linha := [10]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			if j == i {
				diagonal += matrizConfusao[i][j]
			}

			coluna[j] += matrizConfusao[i][j]
			linha[i] += matrizConfusao[i][j]
		}
	}

	for i := 0; i < 10; i++ {
		vp := float64(matrizConfusao[i][i])
		vn := float64(diagonal) - vp
		fp := float64(coluna[i]) - vp
		fn := float64(linha[i]) - vp

		precisao := vp / (vp + fp)
		sensitividade := vp / (vp + fn)
		especificidade := vn / (vn + fp)
		fmt.Printf("classe %d\n", i)
		fmt.Printf("precisao: %f sensitividade: %f especificidade: %f\n", precisao, sensitividade, especificidade)
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
			novoPeso := r.CamadaEntrada.Peso.Obter(iE, iI) + r.TaxaAprendizagem * nE.Saida * nI.Erro
			r.CamadaEntrada.Peso.Adicionar(iE, iI, novoPeso)
		}
	}

	for iI, nI := range r.CamadaIntermediaria.Neuronios {
		for iS, nS := range r.CamadaSaida.Neuronios {
			novoPeso := r.CamadaIntermediaria.Peso.Obter(iI, iS) + r.TaxaAprendizagem * nI.Saida * nS.Erro
			r.CamadaIntermediaria.Peso.Adicionar(iI, iS, novoPeso)
		}
	}
}

func (r *RedeNeural) ObterSaidaNumerica(saida [10]float64) int {
	for i := 0; i < 10; i++ {
		if saida[i] == 1 {
			return i
		}
	}
	return 0
}
