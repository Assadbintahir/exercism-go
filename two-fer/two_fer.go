// Package 2-fer is a problem on the exercism track.
package twofer

import "strings"

/*
If the given name is "Alice", the result should be "One for Alice, one for me."
If no name is given, the result should be "One for you, one for me."
*/
func ShareWith(name string) string {
	result := [] string {"One for ", "X", ", one for me."}
	if name == "" {
		result[1] = "you"
	} else {
		result[1] = name
	}

	return strings.Join(result, "")
}
