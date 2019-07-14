package job

const (
	// QueueBuffer is the buffer size that all job queues are using.
	QueueBuffer = 10
)

// Queue represents a JobQueue as specified in 8.4.
type Queue struct {
	q chan PendingJob
}

// NewQueue returns a new Queue.
func NewQueue() *Queue {
	q := new(Queue)
	q.q = make(chan PendingJob, QueueBuffer)
	return q
}

// Enqueue adds a new PendingJob at the very end of the queue.
func (q *Queue) Enqueue(j PendingJob) {
	q.q <- j
}

// Dequeue removes the first element from the queue and returns it.
func (q *Queue) Dequeue() (PendingJob, bool) {
	j, ok := <-q.q
	return j, ok
}
