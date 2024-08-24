package cmd

import (
	"encoding/hex"
	"github.com/spf13/cobra"
	"strings"
)

var bin2dcb = &cobra.Command{
	Use:   "bin2dcb",
	Short: "Converts between file formats.",
	Long: `
Converts between file formats.
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			input  string
			output string
			err    error
			data   []byte
			result strings.Builder
		)

		if input, err = cmd.Flags().GetString("input"); err != nil {
			return err
		}

		if output, err = cmd.Flags().GetString("output"); err != nil {
			return err
		}

		if data, err = loadBin(input); err != nil {
			return err
		}

		// FIXME make into fixed length lines
		result.WriteString("\n")
		result.WriteString("    DCB ")

		for _, b := range data {
			result.WriteString("$")
			result.WriteString(strings.ToUpper(hex.EncodeToString([]byte{b})))
			result.WriteString(",")
		}

		// save the file
		res := result.String()
		res = strings.Trim(res, ", ")
		//res = "1234,"
		//res = strings.Trim("1234,", ",")
		if err = saveText(output, strings.Trim(res, ",")); err != nil {
			return err
		}

		return nil
	},
}
