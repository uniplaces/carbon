package lang

import (
	"errors"
	"io/ioutil"
)

func LoadLocaleText(l string) ([]byte, error) {
	lText, err := ioutil.ReadFile("./lang/" + l + ".json")
	if err != nil {
		return nil, errors.New("not able to read the lang file:" + err.Error())
	}

	return lText, nil
}
