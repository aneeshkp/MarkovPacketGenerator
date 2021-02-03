package chain

import "github.com/jmcvetta/randutil"

// Chain represents the main markov chain struct
type Chain struct {
	TransitionMatrix [][]float32
	States           []State
}

// Create generates a new Chain struct
func Create(
	transitionMatrix [][]float32,
	states []State,
) Chain {
	return Chain{
		TransitionMatrix: transitionMatrix,
		States:           states,
	}
}

// Loop runs the markov loop
func (c *Chain) Loop(start int) {

	// initialize current state
	currentStateID := start
	// loop through the transition probabilities
	for {
		currentStateChoice := c.getStateChoice(currentStateID)
		currentStateID = currentStateChoice.Item.(int)
		currentState := c.getState(currentStateID)
		currentState.DoAction()
	}
}

// TransitionProbability returns the probability of transitioning from one state to another
func (c *Chain) TransitionProbability(input, output int) float32 {
	return c.TransitionMatrix[input][output]
}

func (c *Chain) getState(id int) State {
	for _, state := range c.States {
		if state.ID == id {
			return state
		}
	}
	return State{}
}

func (c *Chain) getStateChoice(id int) randutil.Choice {

	transitions := c.TransitionMatrix[id-1]
	choices := []randutil.Choice{}
	weight := 1

	for i, prob := range transitions {
		weight = int(prob * float32(100))
		choices = append(choices, randutil.Choice{
			Weight: weight,
			Item:   i + 1,
		})
	}

	choice, _ := randutil.WeightedChoice(choices)
	return choice
}
