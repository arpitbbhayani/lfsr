package cmd

import (
	"fmt"

	"github.com/arpitbbhayani/lfsr/lfsr"
	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

var kBit int

func init() {
	randomCmd.Flags().IntVarP(&kBit, "k", "k", 8, "number of bits in randomly generated numbers")
	rootCmd.AddCommand(randomCmd)
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generates random numbers using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		taps := parseTaps(vars.Taps)
		l := lfsr.NewLFSR(vars.Seed, taps)
		for i := 0; i < 10; i++ {
			fmt.Println(l.NextNumber(kBit))
		}
	},
}
