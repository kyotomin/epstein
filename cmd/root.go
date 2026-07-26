package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "epst",
	Short: "Secure file ecnryption utility",
	Long:  "Encrypt and decrypt files using modern cryptography",
}

func Execute() error {
	return rootCmd.Execute()
}
