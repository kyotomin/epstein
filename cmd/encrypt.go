package cmd

import (
	"os"
	"time"

	"github.com/kyotomin/epstein/internal/prompt"
	"github.com/kyotomin/epstein/internal/ui"
	"github.com/spf13/cobra"
)

var ecnryptCmd = &cobra.Command{
	Use:   "encrypt <file>",
	Short: "Encrypt file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		password, err := prompt.ReadAndConfirmPassword()
		if err != nil {
			return err
		}

		// report init
		start := time.Now()
		input := args[0]
		info, err := os.Stat(input)
		if err != nil {
			return err
		}
		// report init

		ui.EncryptReport(
			input,
			input+".epst",
			info.Size(),
			time.Since(start),
		)

		_ = password

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ecnryptCmd)
}
