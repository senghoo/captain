package settings

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olebedev/config"
)

var cfg *config.Config
var defaultSetting = `
db:
  name:
  user:
  password:
  host:
  port:
`

func init() {
	var err error
	cfg, err = config.ParseYamlFile("config.yml")
	if err != nil {
		fmt.Print("config parse error or not exists, use default")
		cfg, _ = config.ParseYaml(defaultSetting)
		Save()
	}
}

func Save() {
	yml, _ := config.RenderYaml(cfg.Root)
	d := []byte(yml)
	ioutil.WriteFile("config.yml", d, 0644)
}

func Get(path string) (string, error) {
	return cfg.String(path)
}

func GetOrDefault(path, d string) string {
	v, err := cfg.String(path)
	if err != nil || v == "" {
		return d
	}
	return v
}

func Set(path, val string) error {
	return cfg.Set(path, val)
}

func GetStaticPath() string {
	staticPath, _ := Get("application.dir")
	if staticPath == "" {
		staticPath, _ = os.Getwd()
	}
	return staticPath
}
