package main

import (
	"fmt"
	"math"
)

func main() {
	// start_block_reward := 50
	reward_interval := 210000

	current_reward := 50 * (math.Pow10(8))
	total := 0.0
	for current_reward > 0 {
		total += float64(reward_interval) * current_reward
		current_reward /= 2
	}

	fmt.Println("Total BTC to ever be created: ", total, "Satoshi")
}
