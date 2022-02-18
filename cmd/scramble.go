package cmd

import (
	"path"
	"path/filepath"

	"github.com/arpitbbhayani/lfsr/lfsr"
	"github.com/arpitbbhayani/lfsr/media"
	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

var fpath string

func init() {
	scrambleCmd.Flags().StringVarP(&fpath, "f", "f", "", "path of the file to be scrambled")
	scrambleCmd.MarkFlagRequired("f")
	rootCmd.AddCommand(scrambleCmd)
}

var scrambleCmd = &cobra.Command{
	Use:   "scramble",
	Short: "Scrambles the provided file using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		inData := media.ReadBytes(fpath)

		// scramble
		outData := scramble(
			inData, vars.BitLength,
			vars.Seed, parseTaps(vars.Taps))

		// create a new file with the scrambled data
		filename := filepath.Base(fpath)
		filedir := filepath.Dir(fpath)

		media.WriteBytes(path.Join(filedir, "enc."+filename), outData)
	},
}

func scramble(in []byte, n uint32, seed uint32, taps []uint32) []byte {
	l := lfsr.NewLFSR(n, seed, taps)
	out := make([]byte, len(in))
	for i := range in {
		out[i] = in[i] ^ byte(l.NextNumber(8))
	}
	return out
}
