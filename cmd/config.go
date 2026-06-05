package cmd

import (
	"github.com/noemacias/beacon/internal/config"
	"github.com/spf13/cobra"
)

func NewConfigCmd() *cobra.Command {

	return &cobra.Command{
		Use:   `config`,
		Short: "Print the active configuration",
		Run: func(cmd *cobra.Command, args []string) {
			RunConfig(cmd, args)
		},
	}
}

func RunConfig(cmd *cobra.Command, args []string) {

	cfgPath, _ := cmd.Flags().GetString("config")

	cfg := config.NewConfig()

	cfg.Load(cfgPath)

	cfg.Print()
}
