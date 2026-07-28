package cmd

import (
	"os"
	"time"

	"github.com/kyotomin/epstein/internal/prompt"
	"github.com/kyotomin/epstein/internal/service"
	"github.com/kyotomin/epstein/internal/ui"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt <file>",
	Short: "Encrypt file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		password, err := prompt.ReadAndConfirmPassword()
		if err != nil {
			return err
		}

		start := time.Now()
		input := args[0]
		if err := service.EncryptFile(input, password); err != nil {
			return err
		}

		info, err := os.Stat(input)
		if err != nil {
			return err
		}

		ui.EncryptReport(
			input,
			input+".epst",
			info.Size(),
			time.Since(start),
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
