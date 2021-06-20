package ramdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecord_NewRecord(t *testing.T) {
	tests := []struct {
		test           string
		key            string
		keyColumn      string
		data           interface{}
		expectedRecord *Record
		expectedError  string
	}{
		{
			test:      "it should error if json serialization fails",
			key:       "test-record",
			keyColumn: "test-column",
			data: map[string]interface{}{
				"error": make(chan int),
			},
			expectedError: "json: unsupported type: chan int",
		},
		{
			test:      "it should return a valid record when successful",
			key:       "test-record",
			keyColumn: "test-column",
			data:      struct{}{},
			expectedRecord: &Record{
				serialized: []byte("{}"),
				key:        "test-record",
				keyColumn:  "test-column",
				id:         0x267fc212f178ef79,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			r, err := NewRecord(tc.key, tc.keyColumn, tc.data)

			assert.Equal(t, tc.expectedRecord, r)
			if tc.expectedError != "" {
				assert.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}

func TestRecord_keyHash(t *testing.T) {
	tests := []struct {
		test   string
		input  string
		output uint64
	}{
		{
			test:   "it should hash the input deterministically",
			input:  "test string",
			output: 0xd5579c46dfcc7f18,
		},
		{
			test:   "it should produce an entirely different value with a small change to input",
			input:  "test string.",
			output: 0x84083c0b244440c0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			result := keyHash(tc.input)

			assert.Equal(t, tc.output, result)
		})
	}
}

func TestRecord_Deserialize(t *testing.T) {
	tests := []struct {
		test          string
		into          struct{ Key string }
		expectedInto  struct{ Key string }
		expectedError error
	}{
		{
			test: "it should unmarshal record data into into",
			into: struct {
				Key string
			}{},
			expectedInto: struct {
				Key string
			}{
				Key: "test string",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			rec, err := NewRecord("", "", struct{ Key string }{Key: "test string"})
			if err != nil {
				t.Error(err)
			}

			into := tc.into
			err = rec.Deserialize(&into)

			assert.Equal(t, tc.expectedInto, into)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
