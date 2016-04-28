package settings

import (
	"fmt"
	"io/ioutil"

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

func Set(path, val string) error {
	return cfg.Set(path, val)
}
