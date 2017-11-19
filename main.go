package main

import (
	"log"
	"flag"
)

func main() {
	taxaAprendizagem := flag.Float64("taxaAprendizagem", 0.5, "Taxa de aprendizagem. Default: 0.5")
	iteracoesTreinamento := flag.Int("iteracoesTreinamento", 1000, "Iteracoes de treinamento. Default: 1000")
	nroNeuroniosIntermediarios := flag.Int("nroNeuroniosIntermediarios", 10, "Iteracoes de treinamento. Default: 10")
	arqTreino := flag.String("arqTreino", "pendigits.tes", "Iteracoes de treinamento. Default: pendigits.tes")
	arqTeste := flag.String("arqTeste", "pendigits.tra", "Iteracoes de treinamento. Default: pendigits.tra")

	flag.Parse()

	// ler o arquivo de teste
	redeNeural := RedeNeural{}
	redeNeural.Init(*taxaAprendizagem, *nroNeuroniosIntermediarios)

	log.Printf("Iniciou treinamento...\n")

	for i := 0; i < *iteracoesTreinamento; i++ {
		_ = redeNeural.Treinar(*arqTreino)
		//log.Printf("Treinamento %d. Leu %d linhas\n", i, nLinhas)
	}

	log.Printf("Terminou treinamento.\n")

	log.Printf("Iniciou teste...\n")

	redeNeural.Testar(*arqTeste)

	log.Printf("Terminou teste.\n")
}