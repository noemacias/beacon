package config

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"go.yaml.in/yaml/v3"
)

type TLS struct {
	Enabled bool
	Cert    string
	Key     string
}

type Server struct {
	Addr string
	TLS  TLS
}

type Avatar struct {
	Url  string
	Name string
}
type Site struct {
	Title       string
	SubTitle    string
	Description string
	Avatar      Avatar
	Initials    string
	StaticDir   string `yaml:"static_dir"`
}

type Theme struct {
	Name string
}

type Links struct {
	Name        string
	Icon        string
	Url         string
	Description string
}

type LinkList []Links

type Config struct {
	Server Server
	Site   Site
	Theme  Theme
	Links  LinkList
}

func Initials(name string) string {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return ""
	}

	var initials strings.Builder

	for _, part := range parts {
		r := []rune(part)
		if len(r) > 0 {
			initials.WriteRune(unicode.ToUpper(r[0]))
		}
	}

	return initials.String()
}

func (c *Config) Load(configPath string) {

	data, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return
	}
}

func (c *Config) Print() {

	fmt.Println(c.String())
}

func (c *Config) String() string {
	out, err := yaml.Marshal(c)

	if err != nil {
		return ""
	}

	return string(out)
}

func NewConfig() *Config {
	return &Config{
		Server: Server{
			Addr: ":8080",
		},
		Theme: Theme{
			Name: "dark.html",
		},
		Site: Site{
			StaticDir: "./static",
		},
		Links: LinkList{},
	}
}
