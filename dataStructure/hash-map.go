package dataStructure

import "fmt"

type KeyValueStruct struct {
	key   string
	value string
}

type HashMapStruct struct {
	items []KeyValueStruct
	size  int
}

func NewHashMap(initialSize int) *HashMapStruct {
	return &HashMapStruct{
		items: make([]KeyValueStruct, initialSize),
		size:  initialSize,
	}
}

func (h *HashMapStruct) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = 31*hash + int(char)
	}
	return hash % h.size
}

func (h *HashMapStruct) set(key string, value string) {
	index := h.hash(key)

	if h.items[index].key == "" {
		h.items[index] = KeyValueStruct{key: key, value: value}
		return
	} else {
		h.items[index].value = value
		return
	}
}

func (h HashMapStruct) get(key string) (KeyValueStruct, bool) {
	index := h.hash(key)

	if h.items[index].key == "" {
		return KeyValueStruct{}, false
	} else {
		return h.items[index], true
	}
}

func (h *HashMapStruct) remove(key string) {
	index := h.hash(key)

	if h.items[index].key == "" {
		return
	} else {
		h.items[index].key = ""
		h.items[index].value = ""
		return
	}
}

func HashMapMain() {
	hashMap := NewHashMap(10)

	hashMap.set("Thalles", "1")
	hashMap.set("Ian", "2")
	hashMap.set("Angelo", "3")

	fmt.Println(hashMap.get("Thalles"))
	fmt.Println(hashMap.get("Ian"))
	fmt.Println(hashMap.get("Angelo"))
	hashMap.remove("Angelo")
	fmt.Println(hashMap.get("Angelo"))
}
