package object

import (
	"math/rand"
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var o Object

	check := make([]int, len(Keys))

	for i, key := range Keys {

		switch key.Length() {
		case _bool:
			check[i] = rand.Intn(2)
			o.Set(key, check[i])
		case _nibble:
			check[i] = rand.Intn(16)
			o.Set(key, check[i])
		case _uint8:
			check[i] = rand.Intn(256)
			o.Set(key, check[i])
		}
	}
	for i := 0; i < len(Keys); i++ {

		value := o.Get(Keys[i])
		if value != check[i] {
			t.Fatalf("expected %d, got %d, on key %d", check[i], value, i)
		}

	}

}
