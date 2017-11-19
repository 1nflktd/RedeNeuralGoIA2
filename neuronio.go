package main

type Neuronio struct{
	Saida float64
	Entrada float64
	Erro float64
}

func (n *Neuronio) SetSaida(s float64) {
	n.Saida = s
}

func (n *Neuronio) SetErro(s float64) {
	n.Erro = s
}

func (n *Neuronio) GetSaida() float64 {
	return n.Saida
}

func (n *Neuronio) GetErro() float64 {
	return n.Erro
}
