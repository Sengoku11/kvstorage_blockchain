package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TextStorageKeyPrefix is the prefix to retrieve all TextStorage
	TextStorageKeyPrefix = "TextStorage/value/"
)

// TextStorageKey returns the store key to retrieve a TextStorage from the index fields
func TextStorageKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
