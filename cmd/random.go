package cmd

import (
	"fmt"

	"github.com/arpitbbhayani/lfsr/lfsr"
	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

var kBit int

func init() {
	rootCmd.AddCommand(randomCmd)
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generates random bits using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		taps := parseTaps(vars.Taps)
		l := lfsr.NewLFSR(vars.BitLength, vars.Seed, taps)
		for {
			fmt.Print(l.NextBit())
		}
	},
}
