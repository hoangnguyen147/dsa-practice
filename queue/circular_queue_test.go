package queue

import (
	"testing"
)

// MyCircularQueue represents a circular queue data structure
// This is the interface for LeetCode Problem 622: Design Circular Queue
type MyCircularQueue interface {
	EnQueue(value int) bool
	DeQueue() bool
	Front() int
	Rear() int
	IsEmpty() bool
	IsFull() bool
}

func Constructor(k int) MyCircularQueue {
	q := NewQueue(k)
	return &q
}

// TestBasicOperations tests the basic functionality of the circular queue
func TestBasicOperations(t *testing.T) {
	// Test case from LeetCode example
	q := Constructor(3)

	// Test initial state
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.IsFull() {
		t.Error("New queue should not be full")
	}
	if q.Front() != -1 {
		t.Error("Front of empty queue should return -1")
	}
	if q.Rear() != -1 {
		t.Error("Rear of empty queue should return -1")
	}

	// Test enqueue operations
	if !q.EnQueue(1) {
		t.Error("EnQueue(1) should succeed")
	}
	if !q.EnQueue(2) {
		t.Error("EnQueue(2) should succeed")
	}
	if !q.EnQueue(3) {
		t.Error("EnQueue(3) should succeed")
	}

	// Queue should now be full
	if !q.IsFull() {
		t.Error("Queue should be full after adding 3 elements to capacity 3")
	}
	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	// Test that enqueue fails when full
	if q.EnQueue(4) {
		t.Error("EnQueue(4) should fail when queue is full")
	}

	// Test front and rear
	if q.Front() != 1 {
		t.Errorf("Front should be 1, got %d", q.Front())
	}
	if q.Rear() != 3 {
		t.Errorf("Rear should be 3, got %d", q.Rear())
	}

	// Test dequeue operations
	if !q.DeQueue() {
		t.Error("DeQueue should succeed")
	}
	if q.Front() != 2 {
		t.Errorf("Front should be 2 after dequeue, got %d", q.Front())
	}
	if q.Rear() != 3 {
		t.Errorf("Rear should still be 3, got %d", q.Rear())
	}

	// Test enqueue after dequeue
	if !q.EnQueue(4) {
		t.Error("EnQueue(4) should succeed after dequeue")
	}
	if q.Rear() != 4 {
		t.Errorf("Rear should be 4, got %d", q.Rear())
	}
}

// TestEmptyQueueOperations tests operations on empty queue
func TestEmptyQueueOperations(t *testing.T) {
	q := Constructor(2)

	// Test dequeue on empty queue
	if q.DeQueue() {
		t.Error("DeQueue should fail on empty queue")
	}

	// Test front and rear on empty queue
	if q.Front() != -1 {
		t.Error("Front should return -1 for empty queue")
	}
	if q.Rear() != -1 {
		t.Error("Rear should return -1 for empty queue")
	}
}

// TestSingleElementQueue tests queue with capacity 1
func TestSingleElementQueue(t *testing.T) {
	q := Constructor(1)

	// Add one element
	if !q.EnQueue(5) {
		t.Error("EnQueue should succeed")
	}
	if !q.IsFull() {
		t.Error("Queue should be full")
	}
	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}
	if q.Front() != 5 {
		t.Errorf("Front should be 5, got %d", q.Front())
	}
	if q.Rear() != 5 {
		t.Errorf("Rear should be 5, got %d", q.Rear())
	}

	// Try to add another element (should fail)
	if q.EnQueue(6) {
		t.Error("EnQueue should fail when queue is full")
	}

	// Remove the element
	if !q.DeQueue() {
		t.Error("DeQueue should succeed")
	}
	if !q.IsEmpty() {
		t.Error("Queue should be empty after removing only element")
	}
	if q.IsFull() {
		t.Error("Queue should not be full after removing element")
	}
}

// TestCircularBehavior tests the circular nature of the queue
func TestCircularBehavior(t *testing.T) {
	q := Constructor(3)

	// Fill the queue
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	// Remove all elements
	q.DeQueue() // Remove 1
	q.DeQueue() // Remove 2
	q.DeQueue() // Remove 3

	// Queue should be empty
	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}

	// Add elements again to test circular behavior
	if !q.EnQueue(4) {
		t.Error("EnQueue(4) should succeed")
	}
	if !q.EnQueue(5) {
		t.Error("EnQueue(5) should succeed")
	}
	if !q.EnQueue(6) {
		t.Error("EnQueue(6) should succeed")
	}

	// Test values
	if q.Front() != 4 {
		t.Errorf("Front should be 4, got %d", q.Front())
	}
	if q.Rear() != 6 {
		t.Errorf("Rear should be 6, got %d", q.Rear())
	}
}

// TestMixedOperations tests a sequence of mixed operations
func TestMixedOperations(t *testing.T) {
	q := Constructor(4)

	// Sequence: enqueue 1, 2, dequeue, enqueue 3, 4, 5, dequeue, enqueue 6
	operations := []struct {
		op       string
		value    int
		expected interface{}
	}{
		{"enqueue", 1, true},
		{"enqueue", 2, true},
		{"front", 0, 1},
		{"rear", 0, 2},
		{"dequeue", 0, true},
		{"front", 0, 2},
		{"enqueue", 3, true},
		{"enqueue", 4, true},
		// {"enqueue", 5, true}, // Should fail - queue full
		{"rear", 0, 4},
		{"dequeue", 0, true},
		{"front", 0, 3},
		{"enqueue", 5, true},
		{"rear", 0, 5},
	}

	for i, op := range operations {
		switch op.op {
		case "enqueue":
			result := q.EnQueue(op.value)
			if result != op.expected {
				t.Errorf("Operation %d: EnQueue(%d) = %v, expected %v", i, op.value, result, op.expected)
			}
		case "dequeue":
			result := q.DeQueue()
			if result != op.expected {
				t.Errorf("Operation %d: DeQueue() = %v, expected %v", i, result, op.expected)
			}
		case "front":
			result := q.Front()
			if result != op.expected {
				t.Errorf("Operation %d: Front() = %v, expected %v", i, result, op.expected)
			}
		case "rear":
			result := q.Rear()
			if result != op.expected {
				t.Errorf("Operation %d: Rear() = %v, expected %v", i, result, op.expected)
			}
		}
	}
}

// TestLeetCodeExample tests the exact example from LeetCode
func TestLeetCodeExample(t *testing.T) {
	// ["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
	// [[3],[1],[2],[3],[4],[],[],[],[4],[]]
	// Expected: [null,true,true,true,false,3,true,true,true,4]

	q := Constructor(3)

	// enQueue(1) -> true
	if !q.EnQueue(1) {
		t.Error("enQueue(1) should return true")
	}

	// enQueue(2) -> true
	if !q.EnQueue(2) {
		t.Error("enQueue(2) should return true")
	}

	// enQueue(3) -> true
	if !q.EnQueue(3) {
		t.Error("enQueue(3) should return true")
	}

	// enQueue(4) -> false (queue is full)
	if q.EnQueue(4) {
		t.Error("enQueue(4) should return false when queue is full")
	}

	// Rear() -> 3
	if q.Rear() != 3 {
		t.Errorf("Rear() should return 3, got %d", q.Rear())
	}

	// isFull() -> true
	if !q.IsFull() {
		t.Error("isFull() should return true")
	}

	// deQueue() -> true
	if !q.DeQueue() {
		t.Error("deQueue() should return true")
	}

	// enQueue(4) -> true (now there's space)
	if !q.EnQueue(4) {
		t.Error("enQueue(4) should return true after dequeue")
	}

	// Rear() -> 4
	if q.Rear() != 4 {
		t.Errorf("Rear() should return 4, got %d", q.Rear())
	}
}
