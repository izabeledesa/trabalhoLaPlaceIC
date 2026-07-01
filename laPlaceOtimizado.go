package main

import (
	"fmt"
	"math/rand"
)

func imprimeMatriz(mat [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			fmt.Print(mat[contI][contJ], " ")
		}
		fmt.Println()
	}
}

func iniciaMatrizShuffle(mat [][]int, valorInicial int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			mat[contI][contJ] = valorInicial
			valorInicial = valorInicial + 1
		}
	}
	for contI = 0; contI < len(mat)*len(mat); contI++ {
		troca(mat, rand.Intn(len(mat)), rand.Intn(len(mat[0])), rand.Intn(len(mat)), rand.Intn(len(mat[0])))
	}
}

func iniciaMatrizRandomica(mat [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			mat[contI][contJ] = rand.Intn(int(float32(len(mat)) * (1.2)))
		}
	}
}

func criaMatriz(numLinhas, numColunas int) [][]int {
	var matriz [][]int
	var cont int
	matriz = make([][]int, numLinhas)
	for cont = 0; cont < numLinhas; cont++ {
		matriz[cont] = make([]int, numColunas)
	}
	return matriz
}

func troca(mat [][]int, indAi int, indAj int, indBi int, indBj int) {
	var temp int
	temp = mat[indAi][indAj]
	mat[indAi][indAj] = mat[indBi][indBj]
	mat[indBi][indBj] = temp

}

func verificaQuadradaOrdem(mat [][]int) (bool, int) {
	var numLinhas int
	var numColunas int
	var ehQuadrada bool

	numLinhas = len(mat)
	numColunas = len(mat[0])
	ehQuadrada = false
	if numLinhas == numColunas {
		ehQuadrada = true
	}
	return ehQuadrada, numLinhas
}

func copiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int) {
	var contAi, contAj, contBi, contBj, temp, numL, numC int
	numL = len(menor)
	numC = len(menor[0])

	contAi = 0
	for contBi = 0; contBi < numL; contBi++ {
		if contAi == isqn {
			contAi++
		}
		contAj = 0
		for contBj = 0; contBj < numC; contBj++ {
			if contAj == jsqn {
				contAj++
			}
			temp = maior[contAi][contAj]
			menor[contBi][contBj] = temp
			contAj++
		}
		contAi++
	}

	   fmt.Println( )	
	   fmt.Println("Matriz maior")
	   imprimeMatriz(maior)
	   fmt.Println("Matriz menor")
	   imprimeMatriz(menor)
	   //fmt.Println( )

}

func detOrdem1(mat [][]int) int {
	return mat[0][0]
}

func detOrdem2(mat [][]int) int {
	var diagonalP int
	var diagonalI int

	diagonalP = mat[0][0] * mat[1][1]
	diagonalI = mat[1][0] * mat[0][1]

	return (diagonalP - diagonalI)
}

func calculaSinal(indiceL int, indiceC int) int {
	var sinal int
	sinal = -1
	if ((indiceL + indiceC) % 2) == 0 {
		sinal = 1
	}
	return sinal
}

func detOrdemN(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC int
	var matMenor [][]int

	numL = len(mat)
	numC = len(mat[0])

	resposta = 0
	contL = 0

	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = calculaSinal(contL, contC)
		//criando a matriz menor
		matMenor = criaMatriz(numL-1, numC-1)
		copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
		detTemp = determinante(matMenor)
		//fmt.Println("DetTemp ",detTemp)
		//fmt.Println("resposta ",resposta)
		//fmt.Println("cofator ",cofator)
		//fmt.Println("sinal ",sinal)
		resposta = resposta + (cofator * sinal * detTemp)
		//fmt.Println("resposta dps",resposta)
	}

	return resposta
}




// Verifica qual é a linha com mais 0's e quantos 0 ela tem
func melhorLinha(mat [][]int) (int, int) {
    var numI, numJ, contI, contJ int
	var acum, aux int
	var qtds0, melhorI int

	numI = len(mat)
	numJ = len(mat[0])

    melhorI = 0
    qtds0 = 0
	acum = 0
    aux = 0
	for contI = 0; contI < numI; contI++ {
		for contJ = 0; contJ < numJ; contJ++ {
			if mat[contI][contJ] == 0 {
				acum += 1
			}
		}
		if acum > aux {
            melhorI = contI
            qtds0 = acum
		}
        aux = qtds0
        acum = 0
	}

	//fmt.Println("Melhor linha: ", melhorI)
    //fmt.Println("Qtd 0 da melhor linha: ", qtds0)

	return melhorI, qtds0
}

func melhorColuna(mat [][]int) (int, int) {
    var numI, numJ, contI, contJ int
	var acum, aux int
	var qtds0, melhorJ int

	numI = len(mat)
	numJ = len(mat[0])

    melhorJ = 0
    qtds0 = 0
	acum = 0
    aux = 0
	for contJ = 0; contJ < numJ; contJ++ {
		
		for contI = 0; contI < numI; contI++ {
			if mat[contI][contJ] == 0 {
				acum += 1
			}
		}
		if acum > aux {
            melhorJ = contJ
            qtds0 = acum
		}
        aux = qtds0
        acum = 0
	}

    //fmt.Println("Melhor linha: ", melhorJ)
    //fmt.Println("Qtd 0 da melhor linha: ", qtds0)

	return melhorJ, qtds0
}

func linhaOuColuna(mat [][]int) (int, bool) {
    var melhorI, qtd0I int
    var melhorJ, qtd0J int
    var melhorFila int
    var IouJ bool

    melhorI, qtd0I = melhorLinha(mat)
    melhorJ, qtd0J = melhorColuna(mat)

    if(qtd0I >= qtd0J){
        melhorFila = melhorI
        IouJ = true
    } else {
        melhorFila = melhorJ
        IouJ = false
    }

    return melhorFila, IouJ
}

func detOrdemNOtimizado(mat [][]int) int {
	var sinal, cofator, detTemp, resposta int
    var contL, contC, numL, numC, contK, numK int

	var matMenor [][]int

    var melhorFila int
    var IouJ bool
    melhorFila, IouJ = linhaOuColuna(mat)

	numL = len(mat)
	numC = len(mat[0])

    if IouJ == true {
        contL = melhorFila
        contK = contC
        numK = numC
        fmt.Println("Melhor fila: Linha ", melhorFila)
    } else {
        contC = melhorFila
        contK = contL
        numK = numL
        fmt.Println("Melhor fila: Coluna ", melhorFila)
    }
	resposta = 0
	//contL = 0

	for contK = 0; contK < numK; contK++ {
        if IouJ == true {
            cofator = mat[contL][contK]
        } else {
            cofator = mat[contK][contC]
        }
        if cofator == 0 {
            fmt.Println()
            fmt.Println("resposta ",resposta)
            fmt.Println("cofator ",cofator)
            resposta += 0
            fmt.Println("resposta dps",resposta)
        } else {
            sinal = calculaSinal(contL, contC)
		    //criando a matriz menor
		    matMenor = criaMatriz(numL-1, numC-1)
		    copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
		    detTemp = determinante(matMenor)
		    fmt.Println("DetTemp ",detTemp)
		    fmt.Println("resposta ",resposta)
		    fmt.Println("cofator ",cofator)
		    fmt.Println("sinal ",sinal)
		    resposta = resposta + (cofator * sinal * detTemp)
		    fmt.Println("resposta dps",resposta)
        }
		
	}

	return resposta
}




func determinante(mat [][]int) int {
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)
	det = 0
	if ehQuadrada {
		// imprimeMatriz(mat)
		switch ordem {
		case 1:
			// fmt.Println("Ordem 1")
			det = detOrdem1(mat)
		case 2:
			//fmt.Println("Ordem 2")
			det = detOrdem2(mat)
		default:
			//fmt.Println("Ordem ", ordem)
			//det = detOrdemN(mat)
            det = detOrdemNOtimizado(mat)

		}
		//imprimeMatriz(mat)
		//fmt.Println("Det ", det)

	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}

func main() {
	var numLinhas int
	var numColunas int
	var resposta int

	numLinhas = 3
	numColunas = 3

	var matrixA [][]int

	matrixA = criaMatriz(numLinhas, numColunas)
	iniciaMatrizRandomica(matrixA)

	fmt.Println()
	imprimeMatriz(matrixA)
	fmt.Println()
	resposta = determinante(matrixA)
	fmt.Println()
	fmt.Println(resposta)
}
