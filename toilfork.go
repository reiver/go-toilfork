package toilfork


import (
	"errors"

	"github.com/reiver/go-toil"
)


var (
	errNilToiler = errors.New("The received value for the toiler was nil.")
)


// ToilForker is an interface that wraps the Fork, Len and Toil methods.
type ToilForker interface {
	toil.Toiler

	// Len returns the number of toilers registered with this toil fork group.
	Len() int

	// Fork forks off another toiler, using the registered Newer.
	Fork() error
}


type internalToilFork struct {
	newer Newer
	group toil.Group
}


// New returns an initialized ToilForker.
func New(newer Newer) ToilForker {
	group := toil.NewGroup()

	toilForker := internalToilFork{
		newer: newer,
		group:group,
	}

	return &toilForker
}


// NewFunc returns an initialized ToilForker.
func NewFunc(fn func()(toil.Toiler, error)) ToilForker {
	newer := newFuncNewer(fn)

	return New(newer)
}


func (toilForker *internalToilFork) Len() int {
	return toilForker.group.Len()
}


func (toilForker *internalToilFork) Fork() error {
	toiler, err := toilForker.newer.New()
	if nil != err {
		return err
	}
	if nil == toiler {
//@TODO: Return better error.
		return errNilToiler
	}

	toilForker.group.Register(toiler)

	return nil
}


func (toilForker *internalToilFork) Toil() {
	toilForker.group.Toil()
}
