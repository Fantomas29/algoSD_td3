package main

import (
	"fmt"
	"sync"
)

const (
	Search   = 1
	Null     = 0
	Diameter = 5
	NonDef   = -1
	nbNodes  = 11
	nbChan   = 15
)

// ///////////// //
// TD 3 EXO 1 //
// ///////////// //

func node(id int, in []chan int, out []chan int, racine bool) {
	messIn := make([]int, len(in))
	messOut := make([]int, len(out))
	visite := false
	parent := -1

	for i := range out {
		if racine {
			messOut[i] = Search
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
					for j := range out {
						messOut[j] = Search
					}
				}
			}

		} else {
			for j := range out {
				messOut[j] = Null
			}
		}
		if parent >= 0 {
			fmt.Println("id =Mon pere est : ", parent) //max 1 parce que un noeud reçoit au max de 2 autres noeuds
		}
		if id == 7 {
			if visite {
				fmt.Println("tour = ", i, " j'ai recu")
			}
		}
	}
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

	comOut[0] = []chan int{tabChan[0], tabChan[1], tabChan[2]}

	comIn[1] = []chan int{tabChan[0]}
	comOut[1] = []chan int{tabChan[4], tabChan[6], tabChan[5]}

	comIn[2] = []chan int{tabChan[4]}
	comOut[2] = []chan int{tabChan[3]}

	comIn[3] = []chan int{tabChan[3], tabChan[2]}
	comOut[3] = []chan int{tabChan[13]}

	comIn[4] = []chan int{tabChan[11], tabChan[1]}
	comOut[4] = []chan int{tabChan[10]}

	comIn[5] = []chan int{tabChan[13]}
	comOut[5] = []chan int{tabChan[14], tabChan[12]}

	comIn[6] = []chan int{tabChan[14]}
	comOut[6] = []chan int{tabChan[11]}

	comIn[7] = []chan int{tabChan[12], tabChan[10], tabChan[9]}

	comIn[8] = []chan int{tabChan[8], tabChan[6]}
	comOut[8] = []chan int{tabChan[9]}

	comIn[9] = []chan int{tabChan[5]}
	comOut[9] = []chan int{tabChan[7]}

	comIn[10] = []chan int{tabChan[7]}
	comOut[10] = []chan int{tabChan[8]}

	wg.Add(nbNodes)

	for i := 0; i < nbNodes; i++ {
		go func(index int) {
			node(index, comIn[index], comOut[index], index == 0)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
