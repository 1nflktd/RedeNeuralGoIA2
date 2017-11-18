package main

type Peso struct {
	Pesos [][]float64
}

func (p *Peso) Init(n1, n2 int) {
	p.Pesos = make([][]float64, n1)
	for i := range p.Pesos {
	    p.Pesos[i] = make([]float64, n2)
	}
}

func (p *Peso) Obter(n1, n2 int) float64 {
	return p.Pesos[n1][n2]
}

func (p *Peso) Adicionar(n1, n2 int, valor float64) {
	p.Pesos[n1][n2] = valor
}