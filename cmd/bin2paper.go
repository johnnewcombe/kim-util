package cmd

import (
	"encoding/hex"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

const (
	recordStart byte = 0x3B
	dataLen     byte = 0x18
	XOFF        byte = 0x13
)

var bin2paper = &cobra.Command{
	Use:   "bin2paper",
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
			//chksum       int16
			//row          int16
			startAddress string
		)

		// FIXME get from command line
		//startAddress = "0000"

		if input, err = cmd.Flags().GetString("input"); err != nil {
			return err
		}

		if output, err = cmd.Flags().GetString("output"); err != nil {
			return err
		}

		if startAddress, err = cmd.Flags().GetString("start-address"); err != nil {
			return err
		}

		addr, err := strconv.ParseUint(startAddress, 16, 16)
		if err != nil {
			return err
		}

		if data, err = loadBin(input); err != nil {
			return err
		}

		records := getRecordData(data, uint16(addr))

		for _, record := range records {

			result.WriteByte(recordStart)
			result.WriteString(strings.ToUpper(hex.EncodeToString(record)))
			result.WriteString("\r\n")

			// TODO: determine if last romm should also have 6 nulls
			//  NULLS added to last row for now
			//if i < len(records)-1 {
			for j := 0; j < 6; j++ {
				result.WriteByte(0)
			}
			//}
		}

		// all done so add the XOFF
		result.WriteByte(XOFF)

		// save the file
		if err = saveText(output, result.String()); err != nil {
			return err
		}

		return nil
	},
}

// getRecordData returns the records without the initial ";" and without the checksum
func getRecordData(data []byte, startAddress uint16) [][]byte {

	var (
		result = make([][]byte, 0)
		row    = make([]byte, 0)
	)
	// split data
	for i, b := range data {

		if i%int(dataLen) == 0 {

			if len(row) > 0 {
				row[0] = byte(len(row) - 3)
				row = append(row, int16ToBytes(calcChechsum(row), false)...)
				result = append(result, row)
			}
			row = make([]byte, 0)

			row = append(row, dataLen)
			row = append(row, int16ToBytes(startAddress, false)...)
			startAddress = startAddress + uint16(dataLen)

		}

		row = append(row, b)

	}

	// if we have a row with stuff in it it means that the row isn't at its max size
	// so we need to add it to the result and update byte[0] (data length)
	if len(row) > 0 {
		row[0] = byte(len(row) - 3)
		row = append(row, int16ToBytes(calcChechsum(row), false)...)
		result = append(result, row)
	}

	// add the last row
	row = make([]byte, 0)

	// last row starts with zero
	row = append(row, 0)

	// add the number of records and checksum
	row = append(row, int16ToBytes(uint16(len(result)), false)...)
	row = append(row, int16ToBytes(calcChechsum(row), false)...)

	result = append(result, row)

	return result
}

func calcChechsum(data []byte) uint16 {

	var result uint16

	for _, b := range data {
		result += uint16(b)
	}

	return result
}
