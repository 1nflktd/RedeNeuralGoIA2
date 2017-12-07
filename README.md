# RedeNeuralGoIA2

Windows:

	Para rodar basta clicar no arquivo RedeNeuralGoIA2.exe. O programa executará com as seguintes configurações padrões:
  
      - taxaAprendizagem = 0.5
	    - iteracoesTreinamento = 1000
	    - nroNeuroniosIntermediarios = 10
	    - arqTreino = "pendigits.tes"
	    - arqTeste = "pendigits.tra"

	Para mudá-las, basta entrar pelo console (cmd) e rodar o executável com as opções, ex:

		C:\go\src\RedeNeuralGoIA2\RedeNeuralGoIA2.exe -taxaAprendizagem 0.4 -iteracoesTreinamento 500 -nroNeuroniosIntermediarios 20 -arqTreino pendigits2.tes -arqTeste pendigits2.tra

    Para compilar, é necessário instalar o Go (https://golang.org/doc/install?download=go1.9.windows-amd64.msi).
	Depois, rodar (na pasta onde se encontra os arquivos):

		C:\go\src\PredadorPresaIA\go build
