package cmd

import (
	"embed"
	"log/slog"

	"github.com/spf13/cobra"
)

func Execute(templatesFS embed.FS) {

	cmd := cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")

			if debug {

				slog.SetLogLoggerLevel(slog.LevelDebug)
			}
		},
	}

	cmd.PersistentFlags().String("config", "config.yaml", "Config Path")
	cmd.PersistentFlags().Bool("debug", false, "Enable Log Debug")

	cmd.AddCommand(NewServeCmd(templatesFS))
	cmd.AddCommand(NewTemplateCmd(templatesFS))
	cmd.AddCommand(NewConfigCmd())
	cmd.Execute()
}
