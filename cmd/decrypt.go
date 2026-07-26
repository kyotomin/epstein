package cmd

import (
	"fmt"

	"github.com/kyotomin/epstein/internal/prompt"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt <file>",
	Short: "Decrypt file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		password, err := prompt.ReadPassword("Password: ")
		if err != nil {
			return err
		}

		fmt.Println("Decrypt:", args[0])
		fmt.Println("Password:", password)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
