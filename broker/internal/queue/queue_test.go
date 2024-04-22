package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleQueue(t *testing.T) {
	q := New[string](5)

	q.Add("t1")
	q.Add("t2")
	q.Add("t3")
	q.Add("t4")
	q.Add("t5")

	expected := []string{"t1", "t2", "t3", "t4", "t5"}
	assert.Equal(t, q.GetAll(), expected)
}

func TestQueueRemoveLastElements(t *testing.T) {
	q := New[string](5)

	q.Add("t1")
	q.Add("t2")
	q.Add("t3")
	q.Add("t4")
	q.Add("t5")
	q.Add("t6")
	q.Add("t7")

	expected := []string{"t3", "t4", "t5", "t6", "t7"}
	assert.Equal(t, q.GetAll(), expected)
}
