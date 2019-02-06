package pytriples

import "fmt"

// Функция нахождения пифагоровых троек
// В майне можно менять критерий остановки и действия выполняемые с полученными тройками

func MainPytriples() {
	stop := make(chan bool, 1)
	result := make(chan pytriples)
	go getPytriples(stop, result)
	for i := 0; i < 100; i++ {
		p := <-result
		fmt.Printf("\n(%v, %v, %v)", p.x, p.y, p.z)
	}
	stop <- true
}

type pytriples struct {
	x, y, z int
}

func getPytriples(stop chan bool, result chan pytriples) {
	for z := 1; ; z++ {
		for x := 1; x <= z; x++ {
			for y := x; y <= z; y++ {
				select {
				case <- stop:
					return
				default:
					if x*x+y*y == z*z {
						result <- pytriples{x, y, z}
					}
				}


			}
		}
	}
}
