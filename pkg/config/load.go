package config

import (
	"os"
	"todo/internal/utils/must"

	"gopkg.in/yaml.v3"
)

func Load(path string) Config {
	data := must.Eval(os.ReadFile(path))
	var cfg Config
	must.Panic(yaml.Unmarshal(data, &cfg))
	return cfg
}
