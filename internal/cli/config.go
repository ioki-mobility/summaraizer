package cli

import (
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type config struct {
	k *koanf.Koanf
}

func newConfig() *config {
	k := koanf.New(".")
	k.Load(file.Provider("config.json"), json.Parser())
	k.Load(file.Provider(os.ExpandEnv("$HOME/.config/summaraizer/config.json")), json.Parser())
	k.Load(file.Provider(os.ExpandEnv("$HOME/.summaraizer/config.json")), json.Parser())
	k.Load(env.Provider("SUMMARAIZER_", ".", func(s string) string {
		envWithoutPrefix := strings.TrimPrefix(s, "SUMMARAIZER_")
		return strings.Replace(strings.ToLower(envWithoutPrefix), "_", ".", -1)
	}), nil)
	return &config{
		k: k,
	}
}

func (c *config) GetString(path string) string {
	return c.k.String(path)
}
