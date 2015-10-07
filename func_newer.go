package toilfork


import (
	"github.com/reiver/go-toil"
)

type internalFuncNewer struct {
	newFn func()(toil.Toiler, error)
}


func newFuncNewer(newFn func()(toil.Toiler, error)) Newer {
	funcNewer := internalFuncNewer{
		newFn:newFn,
	}

	return &funcNewer
}


func (newer *internalFuncNewer) New() (toil.Toiler, error) {
	return newer.newFn()
}
