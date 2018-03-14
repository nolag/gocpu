package processor

import (
	"container/list"
)

// InteruptAction represents an action to happen later
type InteruptAction struct {
	Action      func()
	Delay       uint64
	Description interface{}
}

// Interuptable is a processor that can have interupts.
// Each InteruptAction will have Delay decremented by 1 after a each step and will run when they hit zero.
// Actions must only be written to (with the exception of test cases for this class).
// Use NewInteruptable to create an instance
type Interuptable struct {
	Processor
	Actions                chan *InteruptAction
	delayedInteruptActions *list.List
}

// NewInteruptable creates a new Interuptable, with a queue size of queueSize for unacknowledge actions.
func NewInteruptable(cpu Processor, queueSize int) *Interuptable {
	queue := make(chan *InteruptAction, queueSize)
	return &Interuptable{cpu, queue, list.New()}
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (cpu *Interuptable) Step() error {

	err := cpu.Processor.Step()

	if err != nil {
		return err
	}

ActionLoop:
	for {
		select {
		case action := <-cpu.Actions:
			cpu.delayedInteruptActions.PushBack(action)
		default:
			break ActionLoop
		}
	}

	toRemove := make([]*list.Element, cpu.delayedInteruptActions.Len())
	itemsToRemove := 0

	for e := cpu.delayedInteruptActions.Front(); e != nil; e = e.Next() {
		action := e.Value.(*InteruptAction)
		if action.Delay == 0 {
			action.Action()
			toRemove[itemsToRemove] = e
			itemsToRemove++
		} else {
			action.Delay--
		}
	}

	for itemToRemove := 0; itemToRemove < itemsToRemove; itemToRemove++ {
		cpu.delayedInteruptActions.Remove(toRemove[itemToRemove])
	}

	return nil
}
