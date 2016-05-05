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
app:
  static:
`

func init() {
	var err error
	cfg, err = config.ParseYamlFile(configFile())
	if err != nil {
		fmt.Print("config parse error or not exists, use default")
		cfg, _ = config.ParseYaml(defaultSetting)
		Save()
	}
}

func configFile() (config string) {
	config = os.Getenv("CAPTAIN_CONFIG")
	if config == "" {
		config = "config.yml"
	}
	return
}

func Save() {
	yml, _ := config.RenderYaml(cfg.Root)
	d := []byte(yml)
	ioutil.WriteFile(configFile(), d, 0644)
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

func GetStaticPath() (path string) {
	path = os.Getenv("CAPTAIN_STATIC")
	if path != "" {
		return
	}

	path, _ = Get("app.static")
	if path != "" {
		return
	}
	path, _ = os.Getwd()
	return
}
