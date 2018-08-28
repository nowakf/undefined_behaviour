package world

type trie []nodule

const (
	null = -1
)

func Trie() *trie {
	return nil
}

type bitmap uint64

//here's a bad idea:
//51 + 12 bits = the last 12 bits can be a either three int4s or one int8 and one int4 - marked by the final bit

func (b bitmap) Superset(c bitmap) bool {
	return true
}

type nodule struct {
	fields bitmap
	child  int
	next   int
}

func (t *trie) Insert(root *nodule, requiredFields bitmap) {
	current := root
	for current.next != null {
		if requiredFields.Superset(current.fields) {

			if current.child != null {
				current = &(*t)[current.child]
				continue
			}

			(*t) = append(*t, nodule{requiredFields, -1, -1})
			current.child = len(*t) - 1
			return
		} else if current.fields.Superset(requiredFields) {

			stash := current.fields
			current.fields = requiredFields
			t.Insert(current, stash)
		}

	}
	(*t) = append(*t, nodule{requiredFields, -1, -1})
	current.next = len(*t) - 1

}

func (t trie) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
