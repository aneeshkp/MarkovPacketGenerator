package chain

import "fmt"

// State represents the transition states
type State struct {
	Payload string
	ID      int
}

// DoAction will perform an action for the current state
func (s *State) DoAction() {

	// TODO: add some json action here

	fmt.Printf("TODO: payload printing, id %d, payload %s \n", s.ID, s.Payload)

	return
}
