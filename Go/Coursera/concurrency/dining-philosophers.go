package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Scheme: 0 - philosopher, \ - sticks
	   0  |  0
	--		   --
	  0		   0
	   /  0	 \
*/

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	number, mealsCount int
	leftCS, rightCS    *ChopS
}

func (p Philo) Eat(c chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p
		if p.mealsCount < 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Printf("starting to eat: %d\n", p.number)
			p.mealsCount++

			p.rightCS.Unlock()
			p.leftCS.Unlock()
			fmt.Printf("finishing eating: %d\n", p.number)
			wg.Done()
		}
	}
}

func host(c chan *Philo) {
	for {
		if len(c) == 2 {
			fmt.Println()
			<-c
			<-c
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan *Philo, 2)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, 0, CSticks[i], CSticks[(i+1)%5]}
	}

	go host(c)

	wg.Add(15)
	for i := 0; i < 5; i++ {
		go philos[i].Eat(c, &wg)
	}
	wg.Wait()
}
