package queue

import (
	"encoding/json"
)

type Queue[T any] struct {
	data []T
	max  int
}

func New[T any](max int) *Queue[T] {
	return &Queue[T]{
		max: max,
	}
}

func (q *Queue[T]) Add(value T) {
	if len(q.data) < q.max {
		q.data = append(q.data, value)
		return
	}

	q.data = append(q.data[1:], value)
}

func (q *Queue[T]) AddAll(values []T) {
	for _, value := range values {
		q.Add(value)
	}
}

func (q *Queue[T]) GetAll() []T {
	return q.data
}

func (q *Queue[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.data)
}

func (q *Queue[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &q.data)
}
