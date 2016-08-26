package carbon

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/uniplaces/carbon/lang"
)

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

func (t *Translator) GetLocale() string {
	return t.locale
}

func (t *Translator) SetLocale(l string) error {
	err := t.AssertValidLocale(l)
	if err != nil {
		return err
	}

	err = t.LoadResource(l)
	if err != nil {
		return err
	}

	t.locale = l

	return nil
}

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

func (t *Translator) LoadResource(l string) error {
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

func (t *Translator) ChooseUnit(unit string, count int64) (string, error) {
	s := strings.Split(t.resources[unit], "|")
	if count > 1 {
		return strings.Replace(s[1], ":count", strconv.FormatInt(int64(count), 10), 1), nil
	}

	return s[0], nil
}

func (t *Translator) ChooseTrans(transID, time string) string {
	return strings.Replace(t.resources[transID], ":time", time, 1)
}

/*
func (t *Translator) Choose(message, number, locale string) (string, error) {
	$parts = explode('|', $message);
	$explicitRules = array();
	$standardRules = array();
	foreach ($parts as $part) {
	$part = trim($part);

	if (preg_match('/^(?P<interval>'.Interval::getIntervalRegexp().')\s*(?P<message>.*?)$/xs', $part, $matches)) {
	$explicitRules[$matches['interval']] = $matches['message'];
	} elseif (preg_match('/^\w+\:\s*(.*?)$/', $part, $matches)) {
	$standardRules[] = $matches[1];
	} else {
	$standardRules[] = $part;
	}
	}

	// try to match an explicit rule, then fallback to the standard ones
	foreach ($explicitRules as $interval => $m) {
	if (Interval::test($number, $interval)) {
	return $m;
	}
	}

	$position = PluralizationRules::get($number, $locale);

	if (!isset($standardRules[$position])) {
	// when there's exactly one rule given, and that rule is a standard
	// rule, use this rule
	if (1 === count($parts) && isset($standardRules[0])) {
	return $standardRules[0];
	}

	throw new \InvalidArgumentException(sprintf('Unable to choose a translation for "%s" with locale "%s" for value "%d". Double check that this translation has the correct plural options (e.g. "There is one apple|There are %%count%% apples").', $message, $locale, $number));
	}

	return $standardRules[$position];
}
*/
