package internal

import "sync"

type State struct {
	timestamp int
	port      string
	mu        sync.Mutex
}

func NewState(port string) *State {
	return &State{
		timestamp: 0,
		port:      port,
	}
}

func (st *State) incrementTimestamp() {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.timestamp++
}
