package synx

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var myInt = NewInt(1)
var wg sync.WaitGroup

func startReaderLocking(id int, wg *sync.WaitGroup) {
	fmt.Printf("[Reader %d] Started\n", id)
	counter := 0
	for counter < 100 {
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Microsecond)
		myInt.SetValue(counter)
		time.Sleep(time.Duration(r) * time.Microsecond)
		myInt.GetValue()
		counter++
	}
	fmt.Printf("[Reader %d] Ended\n", id)
	wg.Done()
}

func TestRace(t *testing.T) {
	noOfReaders := 100
	wg.Add(noOfReaders)
	for i := 0; i < noOfReaders; i++ {
		go startReaderLocking(i, &wg)
	}
	wg.Wait()
}
