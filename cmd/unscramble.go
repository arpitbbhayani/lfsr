package cmd

import (
	"path"
	"path/filepath"

	"github.com/arpitbbhayani/lfsr/media"
	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

func init() {
	unscrambleCmd.Flags().StringVarP(&fpath, "f", "f", "", "path of the file to be unscrambled")
	unscrambleCmd.MarkFlagRequired("f")
	rootCmd.AddCommand(unscrambleCmd)
}

var unscrambleCmd = &cobra.Command{
	Use:   "unscramble",
	Short: "unscrambles the provided file using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		inData := media.ReadBytes(fpath)

		// unscramble
		outData := unscramble(
			inData, vars.BitLength,
			vars.Seed, parseTaps(vars.Taps))

		// create a new file with the scrambled data
		filename := filepath.Base(fpath)
		filedir := filepath.Dir(fpath)

		media.WriteBytes(path.Join(filedir, "dec."+filename), outData)
	},
}

func unscramble(in []byte, n uint32, seed uint32, taps []uint32) []byte {
	return scramble(in, n, seed, taps)
}
