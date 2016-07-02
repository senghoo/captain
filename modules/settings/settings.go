package settings

import (
	"os"
	"path"
)

func Get(key string) string {
	return os.Getenv(key)
}

func GetOrDefault(key, d string) string {
	s := Get(key)
	if s == "" {
		return d
	}
	return s
}

func GetStaticPath() (path string) {
	path = os.Getenv("CAPTAIN_STATIC")
	if path != "" {
		return
	}

	path, _ = os.Getwd()
	return
}

func GetWorkspacePath() (p string) {
	p = os.Getenv("CAPTAIN_WORKSPACE")
	if p != "" {
		return
	}

	p, _ = os.Getwd()
	p = path.Join(p, "workspace")

	return
}

func SiteURL() string {
	return GetOrDefault("SITE_URL", "")
}

func CsrfKey() string {
	return GetOrDefault("CSRF_KEY", "set it to random string")
}
