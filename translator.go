package carbon

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/uniplaces/carbon/lang"
)

// Translator helps to translate time based on locale
type Translator struct {
	locale    string
	cacheDir  string
	resources map[string]string
}

// NewTranslator returns a initialized instance of Translator
func NewTranslator() *Translator {
	return &Translator{
		resources: make(map[string]string),
	}
}

// GetLocale will return locale of a Translator
func (t *Translator) GetLocale() string {
	return t.locale
}

// SetLocale will set locale on a Translator
func (t *Translator) SetLocale(l string) error {
	err := t.AssertValidLocale(l)
	if err != nil {
		return err
	}

	err = t.loadResource(l)
	if err != nil {
		return err
	}

	t.locale = l

	return nil
}

// AssertValidLocale checks if the locale is valid or not
func (t *Translator) AssertValidLocale(l string) error {
	matched, err := regexp.MatchString("^(?:[a-z]{2}|[a-z]{2}(([_-]{1})([a-zA-Z]{2}){1,2}))$", l)
	if err != nil {
		return errors.New("unable to match locale code : " + err.Error())
	}
	if !matched {
		return errors.New("invalid locale code : " + l)
	}

	return nil
}

// loadResource loads the translations according to the locale
func (t *Translator) loadResource(l string) error {
	ltext, err := lang.LoadLocaleText(l)
	if err != nil {
		return err
	}

	err = json.Unmarshal(ltext, &t.resources)
	if err != nil {
		return errors.New("unable to unmarshall locale data : " + err.Error())
	}

	return nil
}

// chooseUnit will choose unit of translations according to the count
func (t *Translator) chooseUnit(unit string, count int64) (string, error) {
	s := strings.Split(t.resources[unit], "|")
	if count > 1 {
		return strings.Replace(s[1], ":count", strconv.FormatInt(int64(count), 10), 1), nil
	}

	return s[0], nil
}

// chooseTrans will choose the word to make a diffForHumans statement
func (t *Translator) chooseTrans(transID, time string) string {
	return strings.Replace(t.resources[transID], ":time", time, 1)
}
