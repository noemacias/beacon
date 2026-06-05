package cmd

import (
	"crypto/tls"
	"embed"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/noemacias/beacon/internal/config"
	"github.com/spf13/cobra"
)

func NewServeCmd(templatesFs embed.FS) *cobra.Command {

	cmd := cobra.Command{
		Use:   `serve`,
		Short: `Start the Beacon web server`,
		Run: func(cmd *cobra.Command, args []string) {
			RunServer(cmd, args, templatesFs)
		},
	}

	return &cmd

}

func RunServer(cmd *cobra.Command, args []string, templatesFS embed.FS) {

	cfgPath, _ := cmd.Flags().GetString("config")

	cfg := config.NewConfig()

	cfg.Load(cfgPath)

	slog.Debug("Server config\n" + cfg.String())

	templatePath := path.Join(
		"templates", cfg.Theme.Name)

	tpl, err := template.New("").
		Funcs(template.FuncMap{
			"initials": config.Initials,
			"year": func() int {
				return time.Now().Year()
			},
		}).
		ParseFS(templatesFS, templatePath)

	if err != nil {
		slog.Error("Failed to register templates", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static",
			http.FileServer(
				http.Dir(cfg.Site.StaticDir),
			),
		),
	)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := tpl.ExecuteTemplate(w, cfg.Theme.Name, cfg); err != nil {
			slog.Error("failed to render template", "error", err)
			http.Error(w, "failed to render page", http.StatusInternalServerError)
		}
	})

	slog.Info("Using template", "path", templatePath)
	slog.Info("Starting Server", "addr", cfg.Server.Addr, "tls", cfg.Server.TLS.Enabled)

	server := http.Server{
		Addr:    cfg.Server.Addr,
		Handler: mux,
	}

	if cfg.Server.TLS.Enabled {

		server.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			},

			PreferServerCipherSuites: true,
		}

		// Redirect to https if if listen address is 443
		if strings.Contains(cfg.Server.Addr, ":443") {

			go func() {
				http.ListenAndServe(":80", http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						target := "https://" + r.Host + r.URL.RequestURI()
						http.Redirect(w, r, target, http.StatusFound)
					},
				))
			}()
		}

		if err := server.ListenAndServeTLS(cfg.Server.TLS.Cert, cfg.Server.TLS.Key); err != nil {
			slog.Error("Failed to start tls server", "error", err)
			os.Exit(1)
		}
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

}
