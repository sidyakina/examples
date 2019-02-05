package pytriples

import "fmt"

// Функция нахождения пифагоровых троек
// В майне можно менять критерий остановки и действия выполняемые с полученными тройками

func MainPytriples() {
	do := make(chan bool, 1)
	result := make(chan pytriples)
	go getPytriples(do, result)
	for i := 0; i < 100; i++ {
		do <- true
		p := <-result
		fmt.Printf("\n(%v, %v, %v)", p.x, p.y, p.z)
	}
	do <- false
}

type pytriples struct {
	x int
	y int
	z int
}

func getPytriples(do chan bool, result chan pytriples) {
	for z := 1; ; z++ {
		for x := 1; x <= z; x++ {
			for y := x; y <= z; y++ {
				switch <-do {
				case true:
					if x*x+y*y == z*z {
						p := pytriples{x, y, z}
						result <- p
					} else {
						do <- true
					}
				case false:
					fmt.Println("exit")
					return
				}
			}
		}
	}
}
