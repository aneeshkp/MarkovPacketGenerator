package main

import "github.com/sjakati98/MarkovPacketGenerator/pkg/chain"

func main() {

	states := []chain.State{
		{
			Payload: "a",
			ID:      1,
		},
		{
			Payload: "b",
			ID:      2,
		},
		{
			Payload: "c",
			ID:      3,
		},
	}

	transition := [][]float32{
		{
			0, 0.1, 0.9,
		},
		{
			0.1, 0, 0.9,
		},
		{
			0.5, 0.5, 0,
		},
	}

	c := chain.Create(transition, states)

	c.Loop(2)

}
