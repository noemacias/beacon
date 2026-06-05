package cmd

import (
	"embed"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func NewTemplateCmd(templatesFS embed.FS) *cobra.Command {
	return &cobra.Command{
		Use:   `templates`,
		Short: `List available themes and templates`,
		Run: func(cmd *cobra.Command, args []string) {
			RunTemplates(templatesFS)
		},
	}
}

func RunTemplates(templatesFS embed.FS) {

	tpls, err := templatesFS.ReadDir("templates")

	if err != nil {
		slog.Error("Failed to list templates", "error", err)
		os.Exit(1)
	}

	fmt.Println("Templates")
	for _, t := range tpls {
		fmt.Println("  ", t.Name())
	}
}
