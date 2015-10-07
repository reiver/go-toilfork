package toilfork


type internalFuncNewer struct {
	newFn func()(Toiler, error)
}


func newFuncNewer(newFn func()(Toiler, error)) Newer {
	funcNewer := internalFuncNewer{
		newFn:newFn,
	}

	return &funcNewer
}


func (newer *internalFuncNewer) New() (Toiler, error) {
	return newer.newFn()
}
