package main

import (
	"fmt"
)

const TaxaAprendizagem = 0.5

type Neuronio struct{
	Saida float64
	Entrada float64
}

func main() {
	// ler o arquivo de teste
	dadosTeste := `
0 89 27 100 42 75 29 45 15 15 37 0 69 2 100 6 2
0 57 31 68 72 90 100 100 76 75 50 51 28 25 16 0 1
0 100 7 92 5 68 19 45 86 34 100 45 74 23 67 0 4
0 67 49 83 100 100 81 80 60 60 40 40 33 20 47 0 1
100 100 88 99 49 74 17 47 0 16 37 0 73 16 20 20 6
0 100 3 72 26 35 85 35 100 71 73 97 65 49 66 0 4
0 39 2 62 11 5 63 0 100 43 89 99 36 100 0 57 0
13 89 12 50 72 38 56 0 4 17 0 61 32 94 100 100 5
`
	redeNeural := RedeNeural{}
	redeNeural.Init(dadosTeste)

	fmt.Printf("Rodou\n")
}