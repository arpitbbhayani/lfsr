package cmd

import (
	"path"
	"path/filepath"

	"github.com/arpitbbhayani/lfsr/lfsr"
	"github.com/arpitbbhayani/lfsr/media"
	"github.com/arpitbbhayani/lfsr/vars"
	"github.com/spf13/cobra"
)

func init() {
	decryptCmd.Flags().StringVarP(&fpath, "f", "f", "", "path of the file to be decrypted")
	decryptCmd.MarkFlagRequired("f")
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts the provided file using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		l := lfsr.NewLFSR(
			vars.Seed,
			parseTaps(vars.Taps),
		)

		data := media.ReadBytes(fpath)
		for i := range data {
			data[i] = data[i] ^ byte(l.NextNumber(8))
		}

		filename := filepath.Base(fpath)
		filedir := filepath.Dir(fpath)
		media.WriteBytes(path.Join(filedir, "dec."+filename), data)
	},
}
