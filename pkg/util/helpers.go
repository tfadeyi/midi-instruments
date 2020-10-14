package util

// GetKeyboardNote takes the keyboard key as input and returns
// a map containing the corresponding notes with their int value
func GetKeyboardNote(key uint8) map[string]uint8 {
	note := key
	switch note % 12 {
	case 0:
		notes := make(map[string]uint8, 2)
		notes["note:c"] = key
		notes["note:b#"] = key
		return notes
	case 1:
		notes := make(map[string]uint8, 2)
		notes["note:c#"] = key
		notes["note:db"] = key
		return notes
	case 2:
		notes := make(map[string]uint8, 1)
		notes["note:d"] = key
		return notes
	case 3:
		notes := make(map[string]uint8, 2)
		notes["note:d#"] = key
		notes["note:eb"] = key
		return notes
	case 4:
		notes := make(map[string]uint8, 2)
		notes["note:e"] = key
		notes["note:fb"] = key
		return notes
	case 5:
		notes := make(map[string]uint8, 2)
		notes["note:f"] = key
		notes["note:e#"] = key
		return notes
	case 6:
		notes := make(map[string]uint8, 2)
		notes["note:f#"] = key
		notes["note:gb"] = key
		return notes
	case 7:
		notes := make(map[string]uint8, 1)
		notes["note:g"] = key
		return notes
	case 8:
		notes := make(map[string]uint8, 2)
		notes["note:g#"] = key
		notes["note:ab"] = key
		return notes
	case 9:
		notes := make(map[string]uint8, 1)
		notes["note:a"] = key
		return notes
	case 10:
		notes := make(map[string]uint8, 2)
		notes["note:a#"] = key
		notes["note:bb"] = key
		return notes
	case 11:
		notes := make(map[string]uint8, 2)
		notes["note:b"] = key
		notes["note:cb"] = key
		return notes
	}
	return make(map[string]uint8, 1)
}
