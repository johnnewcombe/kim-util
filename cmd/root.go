package cmd

import (
	_ "embed"
	"fmt"
	"github.com/johnnewcombe/telstar-util/globals"
	"github.com/spf13/cobra"
	"os"
)

const (
	k_inputFile    = "Source file to be processed."
	k_outputFile   = "Resultant file following conversion."
	k_startAddress = "The load address of the file in hexadecimal format."
)

func init() {

	rootCmd.AddCommand(bin2paper)
	rootCmd.AddCommand(bin2keys)
	rootCmd.AddCommand(bin2dcb)
	rootCmd.AddCommand(version)

	bin2paper.PersistentFlags().StringP("start-address", "a", "0000", k_startAddress)
	bin2paper.PersistentFlags().StringP("input", "i", "a.in", k_inputFile)
	bin2paper.PersistentFlags().StringP("output", "o", "a.out", k_outputFile)

	bin2keys.PersistentFlags().StringP("start-address", "a", "0000", k_startAddress)
	bin2keys.PersistentFlags().StringP("input", "i", "a.in", k_inputFile)
	bin2keys.PersistentFlags().StringP("output", "o", "a.out", k_outputFile)

	bin2dcb.PersistentFlags().StringP("input", "i", "a.in", k_inputFile)
	bin2dcb.PersistentFlags().StringP("output", "o", "a.out", k_outputFile)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error()+".")
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "kim-util",
	Short: "Utility program for converting KIM-1 filetypes. (c) John Newcombe 2024. Version: " + globals.Version,
	Long: `
Utility program for converting KIM-1 filetypes. (c) John Newcombe 2024.`,
}
