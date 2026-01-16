package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, value := range log {
		switch value {
		case '\u2757': // ❗
			return "recommendation"
		case '\U0001F50D':
			return "search"
		case '\u2600': // ☀
			return "weather"
		}
	}

	return "default"
}


// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	runes := []rune(log) 
	
	for index , char := range runes {

		if char == oldRune {
			runes[index] = newRune
		}
	}
	return  string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	size := utf8.RuneCountInString(log)

	if(size <= limit){
		return  true
	}

	return  false

}
