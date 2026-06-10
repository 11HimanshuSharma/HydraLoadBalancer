package backend


import (
	"errors"
	"sync"
)

type Pool struct {
	mu sync.RWMutex

	backends []*Backend

	index map[string]*Backend

}

func NewPool() *Pool {
	return &Pool {
		backends: make([]*Backend, 0),
		index: make(map[string]*Backend),
	}
}

func (p *Pool) AddBackend(b *Backend) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key := b.URL.String()
	if _, exists := p.index[key]; exists {
		return errors.New("backend already exists")
	}
	p.backends = append(p.backends, b)

	p.index[key] = b
	return nil
}

func (p *Pool) RemoveBackend(url string) error {

	p.mu.Lock()
	defer p.mu.Unlock()


	backend, exists := p.index[url]

	if !exists {
		return errors.New("backend not found")
	}

	delete(p.index, url)

	for i, b := range p.backends {
		if b == backend {

			p.backends = append(
				p.backends[:i],
				p.backends[i+1:]...,
			)
			break
		}
	}
	return nil
}


func (p *Pool) GetBackends(url string) (*Backend, bool) {
	p.mu.RLock()

	defer p.mu.RUnlock()

	b, ok := p.index[url]

	return b, ok
}

func (p *Pool) Backends() []*Backend {

	p.mu.RLock()
	defer p.mu.RUnlock()

	copySlice := make([]*Backend, len(p.backends),)

	copy(copySlice, p.backends)
	return copySlice
}


func (p *Pool) HealthyBackends() []*Backend {

	p.mu.RLock()
	defer p.mu.RUnlock()

	result := make(
		[]*Backend,
		0, 
		len(p.backends),
	)

	for _, b := range p.backends {
		if b.IsAlive(){
			result = append(result, b)
		}
	}
	return result
}

func (p *Pool) Size() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.backends)
}

func (p *Pool) HasHealthyBackend() bool {

	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, b := range p.backends {

		if b.IsAlive() {
			return true
		}
	}
	return false
}