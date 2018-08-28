package object

import (
	"math/rand"
	"testing"
	"time"
)

func set(check []int, o Object) Object {
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
	return o
}
func TestSetGet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var o Object
	check := make([]int, len(Keys))

	o = set(check, o)

	for i := 0; i < len(Keys); i++ {

		value := o.Get(Keys[i])
		if value != check[i] {
			t.Fatalf("expected %d, got %d, on key %d", check[i], value, i)
		}

	}

}

func TestSuperset(t *testing.T) {
	var a Object
	var b Object
	firstObjectValues := make([]int, len(Keys))
	a = set(firstObjectValues, a)
	for i, key := range Keys {
		//set the second one with the exclusive max of the first, so it should always be a subset
		if firstObjectValues[i] > 0 {
			if firstObjectValues[i] == 1 {
				b.Set(key, rand.Intn(2))
			} else {
				b.Set(key, rand.Intn(firstObjectValues[i]))
			}
		}

	}

	if !a.Superset(b) {
		diff := a.RelativeComplement(b)
		for i, key := range diff {
			t.Errorf("A: key %d is %d \n B: key %d is %d \n --- \n", i, a.Get(key), i, b.Get(key))

		}
		t.Fatalf("expected a to be superset of b")
	}

}
