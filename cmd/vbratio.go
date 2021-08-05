package cmd

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
)


var calcvbratioCmd *cobra.Command
func runCalcVbratio(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		fmt.Println("You need 3 inputs: Flop, Turn, River betsizes.")
		os.Exit(1)
	}
	bSizes := [3]string{args[0], args[1], args[2]}
	bf, bt, br := splitToFloats(bSizes)
	rr := calcRiverRatio(br)
	rt := calcTurnRatio(bt, rr)
	rf := calcFlopRatio(bf, rt)

	fmt.Println(rr)
	fmt.Println(rt)
	fmt.Println(rf)

	var v, b float64
	v, b = bluffRatioToValueBluffRatio(rf)
	fmt.Printf("Flop: %f : %f \n", v, b)
	v, b = bluffRatioToValueBluffRatio(rt)
	fmt.Printf("Turn: %f : %f \n", v, b)
	v, b = bluffRatioToValueBluffRatio(rr)
	fmt.Printf("River: %f : %f \n", v, b)
}

func splitToFloats(ss [3]string) (float64, float64, float64) {
	var floats [3]float64
	for i, s := range ss {
		// var err error
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		floats[i] = f
	}
	return floats[0], floats[1], floats[2]
}

// Returns bluff ratio when value is 1
func calcRiverRatio(betSize float64) float64 {
	return float64(betSize) / (float64(1) + float64(betSize))
}

func calcTurnRatio(betSize float64, riverBluffRatio float64) float64 {
	return riverBluffRatio + float64(betSize) / (1 + float64(betSize)) * (float64(1) + float64(riverBluffRatio))
}

func calcFlopRatio(betSize float64, turnBluffRatio float64) float64 {
	return calcTurnRatio(betSize, turnBluffRatio)
}

// Returns value : bluff ratio in percentage
func bluffRatioToValueBluffRatio(r float64) (float64, float64) {
	b := r / (float64(1) + r) * 100
	v := 100 - b
	return v, b
}

func init() {
	calcvbratioCmd = &cobra.Command{
		Use: "vbratio",
		Short: "calculate value bluff ratio",
		Run: runCalcVbratio,
	}

	rootCmd.AddCommand(calcvbratioCmd)
}
