package main

import (
	"log"
)

const TaxaAprendizagem = 0.5

func main() {
	// ler o arquivo de teste
	redeNeural := RedeNeural{}
	redeNeural.Init()

	log.Printf("Iniciou treinamento...")

	nLinhas := 0

	for i := 0; i < 500; i++ {
		nLinhas = redeNeural.Treinar("pendigits.tes")

		 log.Printf("Treinamento %d. Leu %d linhas\n", i, nLinhas)
	}

	log.Printf("Terminou treinamento.")
}