package api

type Keyboard interface {
	// RegisterNote registers a callback function triggered whenever the note is played on the keyboard.
	RegisterNote(note string, callback func(buffer []string) error) error
	// RegisterScopedNote registers a callback function under a different scope than the "generic" one.
	// This is triggered only is the whenever the note is played on the keyboard
	RegisterScopedNote(scope string, note string, callback func(buffer []string) error) error
	// RegisterChord registers a callback function triggered whenever the chord is played on the keyboard.
	RegisterChord(note string, callback func(buffer []string) error) error
	// RegisterScopedChord registers a callback function under a different scope than the "generic" one.
	RegisterScopedChord(scope string, note string, callback func(buffer []string) error) error
	// StartListening tells the Keyboard instance to listen for events
	StartListening() error
	// Close closes all open streams present in the Keyboard
	Close() error
}
