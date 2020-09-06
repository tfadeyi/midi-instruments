package akai

import (
	"fmt"
	"github.com/tfadeyi/midi-instruments/api"
	"github.com/tfadeyi/midi-instruments/internal"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"
	"go.uber.org/zap"
)

var _ api.Keyboard = &AkaiLpk25{}

const genericScope string = "generic"

type AkaiLpk25 struct {
	Name      string
	Logger    *zap.SugaredLogger
	Driver    *driver.Driver
	In        midi.In
	Listeners map[string]map[string]func(buf []string) error
	Buffer    []string
}

func NewLpk25() *AkaiLpk25 {
	portName := "LPK25:LPK25 MIDI 1 24:0"

	l, err := zap.NewProduction()
	if err != nil {
		panic("Failure to initialize the zap logger")
	}
	logger := l.Sugar()
	logger.Named("AKAI-LPK25")

	d, err := driver.New()
	if err != nil{
		logger.Fatalf("Error initializing keyboard instance: %s",err)
	}
	drv := d

	in, err := findInstrumentPort(portName,drv)
	if err != nil{
		logger.Fatalf("Error initializing keyboard instance: %s",err)
	}

	err = in.Open()
	if err != nil{
		logger.Fatalf("Error initializing keyboard instance: %s",err)
	}

	return &AkaiLpk25{
		Name:portName,
		Logger:logger,
		Driver:drv,
		In:in,
		Listeners:make(map[string]map[string]func(buf []string)error,15),
		Buffer: []string{},
	}
}

func (k *AkaiLpk25) RegisterNote(note string, callback func(buffer []string) error) error {
	if _, ok := k.Listeners[genericScope]; !ok{
		k.Listeners[genericScope] = make(map[string]func(buf []string)error,24)
	}

	if _, ok := k.Listeners[genericScope][note]; !ok{
		k.Listeners[genericScope][note] = callback
	}
	return nil
}

func (k *AkaiLpk25) RegisterScopedNote(scope string, note string, callback func(buffer []string) error) error {
	if _, ok := k.Listeners[scope]; !ok{
		k.Listeners[scope] = make(map[string]func(buf []string)error,24)
	}

	if _, ok := k.Listeners[scope][note]; !ok{
		k.Listeners[scope][note] = callback
	}
	return nil
}

func (*AkaiLpk25) RegisterChord(note string, callback func(buffer []string) error) error {
	panic("implement me")
}

func (*AkaiLpk25) RegisterScopedChord(scope string, note string, callback func(buffer []string) error) error {
	panic("implement me")
}

func (k *AkaiLpk25) StartListening() error {
	rd := reader.New(
		reader.NoLogger(),
		reader.NoteOn(func(pos *reader.Position, ch uint8, key uint8, vel uint8){
			notes := internal.GetKeyboardNote(key)
			for note, _ := range notes{
				// Append event note to keyboard buffer
				k.Buffer = append(k.Buffer, note)
				//Check for listeners have the event note in the general scope
				if _, ok := k.Listeners["generic"][note]; ok{
					err := k.Listeners["generic"][note](k.Buffer)
					if err != nil{
						k.Logger.Infof("note callback function error: %s",err)
					}
					break
				}
				//Check secondary scopes for note that match the event

			}
		}))
	rd.ListenTo(k.In)
	return nil
}

func (k *AkaiLpk25) Close() error {
	err := k.Driver.Close()
	if err != nil {
		return err
	}
	err = k.In.Close()
	if err != nil {
		return err
	}
	err = k.Logger.Sync()
	return err
}

func findInstrumentPort(name string, drv *driver.Driver) (midi.In, error) {
	ins, err := drv.Ins()
	if err != nil{
		return nil, err
	}
	for _, in := range ins{
		if in.String() == name{
			return in, nil
		}
	}
	return nil, fmt.Errorf("akai keyboard is not connected")
}
