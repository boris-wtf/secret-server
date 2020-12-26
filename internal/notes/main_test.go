package notes

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const maxNoteLength = 256

func randomString(length int) string {
	bytea := make([]byte, length, length)
	_, _ = rand.Read(bytea)
	return string(bytea)
}

func TestRAMRegistry_Create(t *testing.T) {
	type testCase struct {
		name string
		text string
		err  error
	}
	testCases := []testCase{
		{
			name: "ok",
			text: randomString(maxNoteLength - 1),
			err:  nil,
		},
		{
			name: "no text",
			text: "",
			err:  ErrNoText,
		},
		{
			name: "max length inclusion",
			text: randomString(maxNoteLength),
			err:  nil,
		}, {
			name: "too long",
			text: randomString(maxNoteLength + 1),
			err:  ErrTooLong,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reg := NewRAMRegistry(maxNoteLength)

			id, err := reg.Create(tc.text)
			assert.Equal(t, err, tc.err)
			if err != nil {
				assert.Empty(t, id)
			} else {
				assert.NotEmpty(t, id)
			}
		})
	}
}

func TestRAMRegistry_Delete(t *testing.T) {
	type testCase struct {
		name string
		test func(t *testing.T)
	}
	testCases := []testCase{
		{
			name: "ok",
			test: func(t *testing.T) {
				reg := NewRAMRegistry(maxNoteLength)

				givenText := randomString(maxNoteLength - 1)
				id, err := reg.Create(givenText)
				assert.Nil(t, err)

				repliedText, err := reg.Delete(id)
				assert.Nil(t, err)
				assert.Equal(t, givenText, repliedText)
			},
		},
		{
			name: "invalid id",
			test: func(t *testing.T) {
				reg := NewRAMRegistry(maxNoteLength)

				text, err := reg.Delete("invalid id value" /* uuid expected */)
				assert.Empty(t, text)
				assert.Equal(t, err, ErrInvalidID)
			},
		},
		{
			name: "not found",
			test: func(t *testing.T) {
				reg := NewRAMRegistry(maxNoteLength)

				givenText := randomString(maxNoteLength - 1)
				id, err := reg.Create(givenText)
				assert.Nil(t, err)

				repliedText, err := reg.Delete(id)
				assert.Nil(t, err)
				assert.Equal(t, givenText, repliedText)

				repliedText, err = reg.Delete(id)
				assert.Empty(t, repliedText)
				assert.Equal(t, err, ErrNotFound)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}
