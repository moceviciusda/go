package twofer

// ShareWith returns a string that represents a sharing message.
// If the name is empty, it defaults to "you".
// The returned string is formatted as "One for {name}, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
