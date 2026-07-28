package cmd

import (
	"os"
	"time"

	"github.com/kyotomin/epstein/internal/prompt"
	"github.com/kyotomin/epstein/internal/service"
	"github.com/kyotomin/epstein/internal/ui"
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

		start := time.Now()

		if err := service.DecryptFile(args[0], password); err != nil {
			return err
		}

		info, err := os.Stat(args[0])
		if err != nil {
			return err
		}

		ui.DecryptReport(
			args[0],
			time.Since(start),
			info.Size(),
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
