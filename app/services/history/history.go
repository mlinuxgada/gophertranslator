package history

// History - 
type History map[string]string

var hist = History{}

// Get - 
func Get() History {
	return hist
}

// Add - 
func Add(en string, gTranslated string) error {

	hist[en] = gTranslated

	return nil
}
