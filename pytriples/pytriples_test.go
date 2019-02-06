package pytriples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type variants struct {
	pEx              pytriples
	numberIterations int
	name             string
}

func Test_getPytriples(t *testing.T) {
	tests := []variants{
		{pytriples{3, 4, 5}, 1, "iteration №1"},
		{pytriples{6, 8, 10}, 2, "iteration №2"},
		{pytriples{5, 12, 13}, 3, "iteration №3"},
		{pytriples{9, 12, 15}, 4, "iteration №4"},
		{pytriples{8, 15, 17}, 5, "iteration №5"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stop := make(chan bool, 1)
			result := make(chan pytriples)
			p := pytriples{}
			go getPytriples(stop, result)
			for i := 0; i < test.numberIterations; i++ {
				p = <-result
			}
			stop <- true
			assert.Equal(t, test.pEx, p)
		})
	}
}

//goos: windows
//goarch: amd64
//pkg: github.com/sidyakina/Examples/pytriples
//Benchmark_getPytriples-4   	      50	  22741302 ns/op	     358 B/op	       2 allocs/op
func Benchmark_getPytriples(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
			stop := make(chan bool, 1)
			result := make(chan pytriples, 5)
			go getPytriples(stop, result)
			for i := 0; i < 100; i++ {
				<-result
			}
			stop <- true
	}
}
