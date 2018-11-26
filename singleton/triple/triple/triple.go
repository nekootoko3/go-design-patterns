package triple

import "fmt"

type triple struct {
	id int
}

func (t *triple) String() {
	fmt.Printf("Triple %d\n", t.id)
}

var tripleMap map[int]*triple

func GetInstance(id int) *triple {
	if id < 1 || 3 < id {
		return nil
	}

	if tripleMap == nil {
		tripleMap = make(map[int]*triple)
	}

	instance, ok := tripleMap[id]
	if !ok {
		fmt.Printf("create instance: triple %d\n", id)
		instance = &triple{id: id}
		tripleMap[id] = instance
	}

	return instance
}
