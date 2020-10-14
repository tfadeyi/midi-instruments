package util

import (
	"reflect"
	"testing"
)

func TestGetKeyboardNote(t *testing.T) {
	testCases := []struct{
		key uint8
		expected map[string]uint8
	}{
		{
			key: 12,
			expected: map[string]uint8{
				"note:b#": 12,
				"note:c": 12,
			},
		},
		{
			key: 1,
			expected: map[string]uint8{
				"note:c#": 1,
				"note:db": 1,
			},
		},
		{
			key: 14,
			expected: map[string]uint8{
				"note:d": 14,
			},
		},
		{
			key: 3,
			expected: map[string]uint8{
				"note:d#": 3,
				"note:eb": 3,
			},
		},
		{
			key: 4,
			expected: map[string]uint8{
				"note:e": 4,
				"note:fb": 4,
			},
		},
		{
			key: 17,
			expected: map[string]uint8{
				"note:f": 17,
				"note:e#": 17,
			},
		},
		{
			key: 18,
			expected: map[string]uint8{
				"note:f#": 18,
				"note:gb": 18,
			},
		},
		{
			key: 7,
			expected: map[string]uint8{
				"note:g": 7,
			},
		},
		{
			key: 8,
			expected: map[string]uint8{
				"note:g#": 8,
				"note:ab": 8,
			},
		},
		{
			key: 9,
			expected: map[string]uint8{
				"note:a": 9,
			},
		},
		{
			key: 22,
			expected: map[string]uint8{
				"note:a#": 22,
				"note:bb": 22,
			},
		},
		{
			key: 23,
			expected: map[string]uint8{
				"note:b": 23,
				"note:cb": 23,
			},
		},
	}

	for _, element := range testCases{
		got := GetKeyboardNote(element.key)
		if !reflect.DeepEqual(got,element.expected){
			t.Fatalf("unexpected result got:%+v want:%+v", got, element.expected)
		}
	}
}
