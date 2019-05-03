package main

import (
	"fmt"
	"math/rand"
	"time"
)

type human struct {
	age  int
	life int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	you := human{
		age:  0,
		life: rand.Intn(89),
	}

	for you.age < you.life {

		fmt.Println(you.age)
		you.age++

	}

}
