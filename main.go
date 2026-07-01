package main
  
import (
    "fmt"
    "time"
)

func main(){

    var contOrdem int
    var numRepeticoes int
    var contRepeticoes int     
    var inicio time.Time
 	var fim time.Time
 	
        // parametros do experimento
	var ordens []int
	ordens = []int{3, 5, 7, 9, 11}
	//ordens = []int{10, 100, 1000, 10000, 100000}
	numRepeticoes = 3
	//estruturas para processamento do tempo dos experimentos
	var tempoBaseline []int64
	var tempoHacked []int64
        tempoBaseline = make([]int64, len(ordens))
        tempoHacked   = make([]int64, len(ordens))  
    var tempoExperimento int64
    
    // para o exemplo usando matriz

    var numLinhas, numColunas int
    var matrizBase [][]int
  
    //fim da inicalizacao da matriz
      
    for contOrdem = 0; contOrdem < len(ordens); contOrdem++ {
        
        numLinhas = ordens[contOrdem]
        numColunas = ordens[contOrdem]

        matrizBase = criaMatriz(numLinhas, numColunas)
       
        for contRepeticoes = 0; contRepeticoes < numRepeticoes; contRepeticoes++{		

            // para cada repeticao, iniciar uma matriz rand diferente!

            iniciaMatrizRandomica(matrizBase)
	       		
	       	//medir o tempo do baseline com a matriz
			inicio = time.Now()
            determinante(matrizBase)
			fim = time.Now()
			tempoExperimento = fim.UnixNano() - inicio.UnixNano()
			//fmt.Println(tempoExperimento)
			tempoBaseline[contOrdem] = tempoBaseline[contOrdem] + tempoExperimento
			
			
			//medir o tempo da hacked com a mesma matriz anterior
			inicio = time.Now()
            determinanteOtimizado(matrizBase)
			fim = time.Now()
			tempoExperimento = fim.UnixNano() - inicio.UnixNano()
			//fmt.Println(tempoExperimento)
			tempoHacked[contOrdem] = tempoHacked[contOrdem] +  tempoExperimento
		}
		//fmt.Print()
		//fmt.Println(tempoBaseline)
		//fmt.Println(tempoHacked)
		
	}
	
	for contOrdem = 0; contOrdem < len(ordens); contOrdem++ {
		tempoBaseline[contOrdem] = tempoBaseline[contOrdem]/int64(numRepeticoes)
		tempoHacked[contOrdem] = tempoHacked[contOrdem]/int64(numRepeticoes)
	}
	fmt.Println("Tempo no Total")
	fmt.Println(tempoBaseline)
	fmt.Println(tempoHacked)
}
