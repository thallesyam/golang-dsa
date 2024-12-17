package dataStructure

type QueueStruct struct {
	items []string
	size  int
}

func (q *QueueStruct) push(item string) {
	q.items = append(q.items, item)
	q.size++
}

func (q *QueueStruct) pop() string {
	if len(q.items) == 0 {
		return ""
	}

	el := q.items[0]
	q.items = q.items[1:]
	q.size--
	return el
}

func Queue() *QueueStruct {
	return &QueueStruct{}
}
