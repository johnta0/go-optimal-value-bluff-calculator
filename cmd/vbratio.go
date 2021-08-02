package cmd

import "github.com/spf13/cobra"


var vbratioCmd *cobra.Command

// Returns bluff ratio when value is 1
func calcRiverRatio(betSize float32) float32 {
	return 1 / (1 + 2 * betSize)
}

func calcTurnRatio(betSize float32, riverBluffRatio float32) float32 {
	return riverBluffRatio + 1 / (1 + betSize) * (1 + riverBluffRatio)
}

func calcFlopRatio(betSize float32, turnBluffRatio float32) float32 {
	return calcTurnRatio(betSize, turnBluffRatio)
}
