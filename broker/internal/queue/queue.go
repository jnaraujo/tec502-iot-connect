// Estrutura de dado para armazenar os valores de um tipo genérico em uma fila de tamanho máximo.
//
// Exemplo de uso:
//
//	q := queue.New[int](3)
//	q.Add(1)
//	q.Add(2)
//	q.Add(3)
//	q.Add(4)
//	fmt.Println(q.GetAll()) // Output: [2 3 4]
package queue

import (
	"encoding/json"
)

type Queue[T any] struct {
	data []T
	max  int
}

// Cria uma nova fila com o tamanho máximo especificado.
func New[T any](max int) *Queue[T] {
	return &Queue[T]{
		max: max,
	}
}

// Adiciona um valor na fila.
func (q *Queue[T]) Add(value T) {
	if len(q.data) < q.max {
		q.data = append(q.data, value)
		return
	}

	q.data = append(q.data[1:], value)
}

// Adiciona todos os valores na fila.
func (q *Queue[T]) AddAll(values []T) {
	for _, value := range values {
		q.Add(value)
	}
}

// Retorna todos os valores da fila.
func (q *Queue[T]) GetAll() []T {
	return q.data
}

// Usado para serializar a fila em JSON.
func (q *Queue[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.data)
}

// Usado para deserializar a fila em JSON.
func (q *Queue[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &q.data)
}
