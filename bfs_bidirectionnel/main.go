package main

import (
	"fmt"
	"sync"
)

// declaration des constantes
const (
	Search   = 1
	Null     = 0
	Diameter = 4
	NonDef   = -1
	nbNodes  = 12
	nbChan   = 22
	Fils     = 2
)

// ///////////// //
// TD 3 EXO 2 //
// ///////////// //
func node(id int, in []chan int, out []chan int, racine bool) {
	messIn := make([]int, len(in))
	messOut := make([]int, len(out))
	visite := false
	parent := -1
	tabFils := make([]int, 0)
	nbFils := 0
	for i := range out {
		if racine {
			messOut[i] = Search
			visite = true
		} else {
			messOut[i] = Null
		}
	}

	for i := 0; i < Diameter; i++ {
		communication(in, messIn, out, messOut)

		if !visite {
			for j := range in {
				if messIn[j] == Search {
					parent = j
					visite = true
					messOut[parent] = Fils
					for j := range out {
						if j != parent {
							messOut[j] = Search
						}
					}
				}
			}
		} else {
			for j := range out {
				messOut[j] = Null
			}
		}
		for j := range in {
			//on regarde dans messIn si on a recu fils, on le met dans le tableau
			if messIn[j] == Fils { //si mon tableau de fils est plein inutile de remettre
				tabFils = append(tabFils, j)
				nbFils++
			}
		}
	}
	fmt.Println("id = ", id, "tabFils = ", tabFils)
}

func receive(messIn []int, in []chan int, wg *sync.WaitGroup) {
	wg.Add(len(in))
	for i := 0; i < len(in); i++ {
		go func(index int) {
			messIn[index] = <-in[index]
			wg.Done()
		}(i)
	}
}

func send(sendOut []int, out []chan int, wg *sync.WaitGroup) {
	wg.Add(len(out))
	for i := 0; i < len(out); i++ {
		go func(index int) {
			out[index] <- sendOut[index]
			wg.Done()
		}(i)
	}
}

func communication(in []chan int, messIn []int, out []chan int, messOut []int) {
	var wg sync.WaitGroup
	send(messOut, out, &wg)
	receive(messIn, in, &wg)
	wg.Wait()
}

func main() {
	var tabChan [nbChan]chan int
	for i := range tabChan {
		tabChan[i] = make(chan int)
	}
	var (
		comIn  [nbNodes][]chan int
		comOut [nbNodes][]chan int
		wg     sync.WaitGroup
	)

	comIn[0] = []chan int{tabChan[5], tabChan[18]}
	comOut[0] = []chan int{tabChan[16], tabChan[6]}

	comIn[1] = []chan int{tabChan[15], tabChan[14]}
	comOut[1] = []chan int{tabChan[0], tabChan[7]}

	comIn[2] = []chan int{tabChan[17], tabChan[16], tabChan[0]}
	comOut[2] = []chan int{tabChan[1], tabChan[5], tabChan[15]}

	comIn[3] = []chan int{tabChan[13], tabChan[7]}
	comOut[3] = []chan int{tabChan[8], tabChan[14]}

	comIn[4] = []chan int{tabChan[6]}
	comOut[4] = []chan int{tabChan[18]}

	comIn[5] = []chan int{tabChan[10]}
	comOut[5] = []chan int{tabChan[12]}

	comIn[6] = []chan int{tabChan[8], tabChan[12], tabChan[11]}
	comOut[6] = []chan int{tabChan[13], tabChan[10], tabChan[9]}

	comIn[7] = []chan int{tabChan[9]}
	comOut[7] = []chan int{tabChan[11]}

	comIn[8] = []chan int{tabChan[2]}
	comOut[8] = []chan int{tabChan[21]}

	comIn[9] = []chan int{tabChan[1], tabChan[21], tabChan[20], tabChan[19]}
	comOut[9] = []chan int{tabChan[17], tabChan[2], tabChan[3], tabChan[4]}

	comIn[10] = []chan int{tabChan[3]}
	comOut[10] = []chan int{tabChan[20]}

	comIn[11] = []chan int{tabChan[4]}
	comOut[11] = []chan int{tabChan[19]}

	wg.Add(nbNodes)

	for i := 0; i < nbNodes; i++ {
		go func(index int) {
			node(index, comIn[index], comOut[index], index == 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
