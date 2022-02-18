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
	encryptCmd.Flags().StringVarP(&fpath, "f", "f", "", "path of the file to be encrypted")
	encryptCmd.MarkFlagRequired("f")
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts the provided file using LFSR",
	Run: func(cmd *cobra.Command, args []string) {
		l := lfsr.NewLFSR(
			vars.BitLength,
			vars.Seed,
			parseTaps(vars.Taps),
		)

		data := media.ReadBytes(fpath)
		for i := range data {
			data[i] = data[i] ^ byte(l.NextNumber(8))
		}

		filename := filepath.Base(fpath)
		filedir := filepath.Dir(fpath)
		media.WriteBytes(path.Join(filedir, "enc."+filename), data)
	},
}
