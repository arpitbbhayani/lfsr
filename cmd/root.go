package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lfsr",
	Short: "LFSR based encryption and decryption",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.PersistentFlags().Uint32VarP(&vars.BitLength, "bits", "n", 32, "number of bits in the LFSR")
	rootCmd.PersistentFlags().Uint32VarP(&vars.Seed, "seed", "s", 32109412, "seed integer")
	rootCmd.PersistentFlags().StringVarP(&vars.Taps, "taps", "t", "1", "comma separated tap positions (starts with 0)")
	rootCmd.PersistentFlags().IntVar(&vars.RandomCount, "count", 10, "number of random bits/numbers to be generated")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseTaps(tapsStr string) []uint32 {
	tokens := strings.Split(tapsStr, ",")
	taps := make([]uint32, len(tokens))
	for i, token := range tokens {
		if n, err := strconv.Atoi(token); err == nil {
			taps[i] = uint32(n)
		}
	}
	return taps
}
