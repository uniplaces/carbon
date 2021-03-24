package lang

import (
	"errors"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadLocaleText will load the translations from the locale json files according to the locale
func LoadLocaleText(l string) ([]byte, error) {
	lText, err := ioutil.ReadFile("./lang/" + l + ".json")
	if os.IsNotExist(err) {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		lPath := filepath.Join(gopath, "src", "github.com", "uniplaces", "carbon", "lang", l + ".json")
		lText, err = ioutil.ReadFile(lPath)
	}

	if err != nil {
		return nil, errors.New("not able to read the lang file:" + err.Error())
	}

	return lText, nil
}
