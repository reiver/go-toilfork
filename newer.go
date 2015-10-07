package toilfork


import (
	"github.com/reiver/go-toil"
)


type Newer interface {
	New() (toil.Toiler, error)
}
