package carbon

/**
* Create a new CarbonInterval instance from specific values.
* This is an alias for the constructor that allows better fluent
* syntax as it allows you to do CarbonInterval::create(1)->fn() rather than
* (new CarbonInterval(1))->fn().
 */
//func Create() {
//}

/**
* Create a CarbonInterval instance from a DateInterval one.  Can not instance
* DateInterval objects created from DateTime::diff() as you can't externally
* set the $days field.
 */
//func Instance() {
//}

// Get the translator instance in use
//func GetTranslator() {
//}

/*
* Set the translator instance to use
 */
//func SetTranslator() {
//}

// Get the current translator locale
//func GetLocale() {
//}

// Set the current translator locale
//func SetLocale() {
//}

// Allow setting of weeks and days to be cumulative.
func WeeksAndDays() {
}

// Get the current interval in a human readable format in the current locale.
func ForHumans() {
}

// Add the passed interval to the current instance
func Add() {
}

// Get the interval_spec string
func Spec() {
}
