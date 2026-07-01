package main
  
import (
    "fmt"
    "time"
)

func fatorialI(num int) int{
	var cont int
	var resposta int
	resposta = 1
	for cont= 2;cont<num; cont++{
		resposta = resposta * cont
	}
	
	return resposta
}

func fatorialR(num int) int{
	var resposta int
	
        if num < 2 {
        	resposta = 1	
        } else {
        	resposta = resposta * fatorialR(num-1)
        }
        
	return resposta
}

func main(){

    var contOrdem int
    var numRepeticoes int
    var contRepeticoes int     
    var inicio time.Time
 	var fim time.Time
 	
        // parametros do experimento
	var ordens []int
	//ordens = []int{3, 5, 7, 9, 11}
	ordens = []int{10, 100, 1000, 10000, 100000}
	numRepeticoes = 3
	//estruturas para processamento do tempo dos experimentos
	var tempoBaseline []int64
	var tempoHacked []int64
        tempoBaseline = make([]int64, len(ordens))
        tempoHacked   = make([]int64, len(ordens))  
    var tempoExperimento int64
    
    // para o exemplo usando matriz
  
    //fim da inicalizacao da matriz
      
    for contOrdem = 0; contOrdem < len(ordens); contOrdem++ {
       
        for contRepeticoes = 0; contRepeticoes < numRepeticoes; contRepeticoes++{		

            // para cada repeticao, iniciar uma matriz rand diferente!
	       		
	       	//medir o tempo do baseline com a matriz
			inicio = time.Now()
			fatorialR(ordens[contOrdem])
			fim = time.Now()
			tempoExperimento = fim.UnixNano() - inicio.UnixNano()
			//fmt.Println(tempoExperimento)
			tempoBaseline[contOrdem] = tempoBaseline[contOrdem] + tempoExperimento
			
			
			//medir o tempo da hacked com a mesma matriz anterior
			inicio = time.Now()
			fatorialI(ordens[contOrdem])
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
