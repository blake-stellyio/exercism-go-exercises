
// Package twofer should have a package comment that summarizes what it's about.
package twofer

import "fmt"
func ShareWith(name string) string {
	var placeholder string
    if name == "" {
        placeholder = "you"
    } else {
        placeholder = name
    }
	return fmt.Sprintf("One for %s, one for me.", placeholder)
}
