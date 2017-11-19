package main

import (
	"log"
)

const TaxaAprendizagem = 0.5

func main() {
	// ler o arquivo de teste
	redeNeural := RedeNeural{}
	redeNeural.Init()

	log.Printf("Iniciou treinamento...\n")

	for i := 0; i < 1; i++ {
		//_ = redeNeural.Treinar("pendigits.tes")
		_ = redeNeural.Treinar("pendigits.tes2")

		//log.Printf("Treinamento %d. Leu %d linhas\n", i, nLinhas)
	}

	log.Printf("Terminou treinamento.\n")

	/*
	log.Printf("Iniciou teste...\n")

	redeNeural.Testar("pendigits.tra")

	log.Printf("Terminou teste.\n")
	*/
}