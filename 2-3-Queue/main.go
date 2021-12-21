package main

import (
	"container/list"
	"fmt"
)

// Queue 구조체 정의
type Queue struct {
	v *list.List
}

// 요소 추가
func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

// 요소를 반환하면서 삭제
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	// 새로운 큐 생성
	queue := NewQueue()

	// 요소 입력
	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()

	// 요소 출력
	for v != nil {
		fmt.Print("%v ->", v)
		v = queue.Pop()
	}
}
