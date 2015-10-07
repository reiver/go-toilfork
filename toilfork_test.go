package toilfork


import (
	"testing"

	"github.com/reiver/go-toil/toiltest"

	"math/rand"
	"sync"
	"time"
)



func TestNew(t *testing.T) {

	newFn := func() (Toiler, error) {
		return nil, nil
	}

	newer := newFuncNewer(newFn)

	toilForker := New(newer)
	if nil == toilForker {
		t.Errorf("After calling New(), expected returned value not to be nil, but instead was: %v", toilForker)
	}
}


func TestNewFunc(t *testing.T) {

	newFn := func() (Toiler, error) {
		return nil, nil
	}

	toilForker := NewFunc(newFn)
	if nil == toilForker {
		t.Errorf("After calling New(), expected returned value not to be nil, but instead was: %v", toilForker)
	}
}


func TestLen(t *testing.T) {

	newFn := func() (Toiler, error) {
		return nil, nil
	}

	toilForker := NewFunc(newFn)

	length := toilForker.Len()

	if expected, actual := 0, length; expected != actual {
		t.Errorf("Expected the number of registered toilers to be %d, but actually was %d.", expected, actual)
		return
	}
}


func TestFork(t *testing.T) {

	toiler := toiltest.NewRecorder()

	newFn := func() (Toiler, error) {
		return toiler, nil
	}

	toilForker := NewFunc(newFn)

	const NUM_FORK_TESTS = 20
	for testNumber:=0; testNumber<NUM_FORK_TESTS; testNumber++ {

		toilForker.Fork()

		length := toilForker.Len()

		if expected, actual := 1+testNumber, length; expected != actual {
		t.Errorf("For test #%d, after fork, expected the number of toilers to be %d, but actually was %d.", testNumber, expected, actual)
			continue
		}
	}
}



func TestToil(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )


	// Do tests.
	const NUM_TOIL_TESTS = 20
	for testNumber:=0; testNumber<NUM_TOIL_TESTS; testNumber++ {

		numberOfTimesToToil := randomness.Intn(44)

		var waitGroup sync.WaitGroup
		waitGroup.Add(numberOfTimesToToil)

		toiler := toiltest.NewRecorder()
		toiler.ToilFunc(func(){
			waitGroup.Done()
		})


		newFn := func() (Toiler, error) {
			return toiler, nil
		}

		toilForker := NewFunc(newFn)


		for i:=0; i<numberOfTimesToToil; i++ {
			toilForker.Fork()
		}


		go toilForker.Toil()


		waitGroup.Wait() // Make sure all the calls on the Toil() method are done before continuing.


		if expected, actual := numberOfTimesToToil, toiler.NumToiling(); expected != actual {
			t.Errorf("For test #%d with, expected the number of toiling toilers to be %d, but actually was %d.", testNumber, expected, actual)
			continue
		}
	}
}
