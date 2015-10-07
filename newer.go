package toilfork


type Newer interface {
	New() (Toiler, error)
}
