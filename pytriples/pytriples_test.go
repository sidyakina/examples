package pytriples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type variants struct {
	pEx pytriples
	numberIterations int
	name string
}

func Test_getPytriples(t *testing.T) {
	tests := []variants{
		{pytriples{3, 4, 5}, 1, "iteration №1"},
		{pytriples{6, 8, 10}, 2, "iteration №2"},
		{pytriples{5, 12, 13}, 3, "iteration №3"},
		{pytriples{9, 12, 15}, 4, "iteration №4"},
		{pytriples{8, 15, 17}, 5, "iteration №5"},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			do := make(chan bool, 1)
			result := make(chan pytriples)
			p := pytriples{}
			go getPytriples(do, result)
			for i := 0; i < test.numberIterations; i++ {
				do <- true
				p = <-result
			}
			do <- false
			assert.Equal(t, test.pEx, p)
		})
	}
}


// Benchmark_getPytriples-4   	       5	 226612960 ns/op	     435 B/op	       1 allocs/op
func Benchmark_getPytriples(b *testing.B) {
	b.ReportAllocs()
	do := make(chan bool, 1)
	result := make(chan pytriples)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go getPytriples(do, result)
		for i := 0; i < 100; i++ {
			do <- true
			<-result
		}
		do <- false
	}
}