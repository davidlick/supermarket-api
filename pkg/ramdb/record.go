package ramdb

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"

	"github.com/google/btree"
)

type Record struct {
	serialized []byte
	keyColumn  string
	key        string
	id         uint64
}

// NewRecord returns a pointer to a Record populated with data, key, and a hash of the key used for ordering in the tree.
func NewRecord(key, keyColumn string, data interface{}) (*Record, error) {
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var r Record
	r.serialized = serialized
	r.keyColumn = keyColumn
	r.key = key
	r.id = keyHash(key)
	return &r, nil
}

func keyHash(s string) uint64 {
	h := sha256.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)

	return binary.BigEndian.Uint64(sum)
}

// Deserialize unmarshals the serialized data into `into`.
func (r *Record) Deserialize(into interface{}) error {
	err := json.Unmarshal(r.serialized, &into)
	return err
}

// Less is used to order items and for looking up Records in the tree.
func (r *Record) Less(than btree.Item) bool {
	re := than.(*Record)
	return r.id < re.id
}
