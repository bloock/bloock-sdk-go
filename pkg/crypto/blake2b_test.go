package crypto

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlake2bHash(t *testing.T) {
	//b2b.Sum256([]byte("Hello, world!"))
	//hash, _ := b2b.New256([]byte("Hello, world!"))
	hash := Blake2b().Hash([]byte("Hello, world!"))

	hx := hex.EncodeToString(hash)
	assert.Equal(t, "b5da441cfe72ae042ef4d2b17742907f675de4da57462d4c3609c2e2ed755970", hx)
}
