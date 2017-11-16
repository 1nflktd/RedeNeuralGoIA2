package main

type CamadaIntermediaria struct {
	Neuronios [13]Neuronio // transformar para variar
	Peso Peso // Intermediaria -> Saida
	Saida int
}

func (c *CamadaIntermediaria) Init() {
	c.Peso.Init(13, 10)
}
