package cmd

import (
	"fmt"

	"github.com/arpitbbhayani/lfsr/lfsr"
	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(randomBitsCmd)

	randomNumbersCmd.PersistentFlags().IntVarP(&vars.KBit, "k", "k", 8, "width of the random number (in bits)")
	rootCmd.AddCommand(randomNumbersCmd)
}

var randomBitsCmd = &cobra.Command{
	Use:   "random-bits",
	Short: "Generates random bits using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		taps := parseTaps(vars.Taps)
		l := lfsr.NewLFSR(vars.BitLength, vars.Seed, taps)
		for i := 0; i < vars.RandomCount; i++ {
			fmt.Println(l.NextBit())
		}
	},
}

var randomNumbersCmd = &cobra.Command{
	Use:   "random-numbers",
	Short: "Generates random numbers using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		taps := parseTaps(vars.Taps)
		l := lfsr.NewLFSR(vars.BitLength, vars.Seed, taps)
		for i := 0; i < vars.RandomCount; i++ {
			fmt.Println(l.NextNumber(vars.KBit))
		}
	},
}
