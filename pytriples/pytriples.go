package pytriples

import "fmt"

// Функция нахождения пифагоровых троек
// В майне можно менять критерий остановки и действия выполняемые с полученными тройками

func MainPytriples() {
	stop := make(chan bool, 1)
	result := make(chan pytriples, 5)
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

// при написании использовалась статья https://habr.com/ru/company/jugru/blog/438260/
func getPytriples(stop chan bool, result chan pytriples) {
	x := 1
	y := 1
	z := 1
	for {
		select {
		case <-stop:
			return
		default:
			for {
				if y <= z {
					y++
				} else {
					if x <= z {
						x++
					} else {
						x = 1
						z++
					}
					y = x
				}
				if x*x+y*y == z*z {
					result <- pytriples{x, y, z}
					break
				}
			}
		}
	}
}