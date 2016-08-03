package carbon

/**
* Create a new CarbonInterval instance from specific values.
* This is an alias for the constructor that allows better fluent
* syntax as it allows you to do CarbonInterval::create(1)->fn() rather than
* (new CarbonInterval(1))->fn().
 */
func create() {
}

/**
* Create a CarbonInterval instance from a DateInterval one.  Can not instance
* DateInterval objects created from DateTime::diff() as you can't externally
* set the $days field.
 */
func instance() {
}

// Get the translator instance in use
func getTranslator() {
}

/*
* Set the translator instance to use
 */
func setTranslator() {
}

// Get the current translator locale
func getLocale() {
}

// Set the current translator locale
func setLocale() {
}

// Allow setting of weeks and days to be cumulative.
func weeksAndDays() {
}

// Allow fluent calls on the setters... CarbonInterval::years(3)->months(5)->day().
func __call() {
}

// Get the current interval in a human readable format in the current locale.
func forHumans() {
}

// Format the instance as a string using the forHumans() function.
func __toString() {
}

// Add the passed interval to the current instance
func add() {
}

// Get the interval_spec string
func spec() {
}
