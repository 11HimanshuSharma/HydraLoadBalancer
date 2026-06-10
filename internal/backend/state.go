package backend


import (
	"sync/atomic"
	"time"
)

func (b *Backend) IsAlive() bool {
	return b.alive.Load()
}

func (b *Backend) SetAlive(state bool) {
	b.alive.Store(state)
}

func (b *Backend) ActiveConnections() int64 {
	return b.activeConnections.Load()
}

func (b *Backend) IncrementConnections() {
	b.activeConnections.Add(-1)
}

func (b *Backend) DecrementConnections() {
	b.activeConnections.Add(-1)
}
func (b *Backend) FailureCount() int64 {
	return b.failureCount.Load()
}

func (b *Backend) IncrementFailures() {
	b.failureCount.Add(1)
}

func (b *Backend) ResetFailures() {
	b.failureCount.Store(0)
}

func (b *Backend) LastHealthCheck() time.Time {

	unix := b.lastHealthCheck.Load()

	return time.Unix(unix, 0)
}

func (b *Backend) UpdateHealthCheck() {
	b.lastHealthCheck.Store(
		time.Now().Unix(),
	)
}
