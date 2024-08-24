package cmd

import (
	"encoding/hex"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

const (
	SPC  byte = 0x20
	STOP byte = 0x2e
)

var bin2keys = &cobra.Command{
	Use:   "bin2keys",
	Short: "Converts between file formats.",
	Long: `
Converts between file formats.
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			input        string
			output       string
			err          error
			data         []byte
			result       strings.Builder
			startAddress string
		)

		if input, err = cmd.Flags().GetString("input"); err != nil {
			return err
		}

		if output, err = cmd.Flags().GetString("output"); err != nil {
			return err
		}

		if startAddress, err = cmd.Flags().GetString("start-address"); err != nil {
			return err
		}

		_, err = strconv.ParseUint(startAddress, 16, 16)
		if err != nil {
			return err
		}

		if data, err = loadBin(input); err != nil {
			return err
		}

		result.WriteString(startAddress)
		result.WriteByte(SPC)

		for _, b := range data {
			result.WriteString(strings.ToUpper(hex.EncodeToString([]byte{b})))
			result.WriteByte(STOP)
		}

		// save the file
		if err = saveText(output, result.String()); err != nil {
			return err
		}

		return nil
	},
}
