/*
Package toilfork provides simple functionality for managing a group of toilers (i.e., workers) where each toiler in the group is basically the "same" toiler.

In this use case, the toilers are all created from a single toiler creation func or interface that is registered when creating the toilfork group is created.

Usage

To use, create one or more types that implement the toilfork.Toiler interface. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() toilfork.Toiler {
	
		toiler := AwesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}

Then create a toilfork.ToilForker. For example:

	toilForker := toilfork.NewFunc(newAwesomeToiler)

(Note that there is also the toilfork.New() func too, that could be used instead of this.)

Then fork off one of more toilers (i.e., types that implement the toil.Toiler interface)
by calling the toil forker's Fork method. For example:

	const NUM_TOILERS = 7
	for i:=0; i<NUM_TOILERS; i++ {
		toilForker.Fork()
	}

Then you can call the Toil method of the toil forker, and it will cause all the toilers
inside of it to start toiling (by calling each of their Toil methods). For example:

	toilForker.Toil()

Observers

A toiler's Toil method can finish in one of two ways. Either it will return gracefully, or
it will panic().

The toil forker is OK with either.

But also, the toil forker provides the toiler with a convenient way of being notified
of each case.

If a toiler also has a Terminated() method, then the toil forker will call the toiler's
Terminated() method when the toiler's Toil() method has returned gracefully. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) Terminated() {
		//@TODO: Do something with this notification.
	}

If a toiler also has a Recovered() method, then the toil forker will call the toiler's
Recovered() method when the toiler's Toil() method has panic()ed. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) Recovered() {
		//@TODO: Do something with this notification.
	}

And of course, a toiler can take advantage of both of these notifications and have
both a Recovered() and Terminated() method. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) Recovered() {
		//@TODO: Do something with this notification.
	}
	
	func (toiler *awesomeToiler) Terminated() {
		//@TODO: Do something with this notification.
	}

*/
package toilfork
