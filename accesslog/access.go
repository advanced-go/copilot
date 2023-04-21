package accesslog

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Entry struct {
	StartTime   string        `json:"start-time"`
	Duration    time.Duration `json:"duration-ms"`
	Traffic     string        `json:"traffic"`
	RouteName   string        `json:"route-name"`
	Region      string        `json:"region"`
	Zone        string        `json:"zone"`
	SubZone     string        `json:"sub-zone"`
	Service     string        `json:"service"`
	InstanceId  string        `json:"instance-id"`
	Method      string        `json:"method"`
	Url         string        `json:"url"`
	Host        string        `json:"host"`
	Path        string        `json:"path"`
	Protocol    string        `json:"protocol"`
	RequestId   string        `json:"request-id"`
	Forwarded   string        `json:"forwarded"`
	StatusCode  int           `json:"status-code"`
	StatusFlags string        `json:"status-flags"`
	TimeoutMs   int           `json:"timeout-ms"`
	RateLimit   rate.Limit    `json:"rate-limit"`
	RateBurst   int           `json:"rate-burst"`
	Retry       bool          `json:"retry"`
	Proxy       bool          `json:"proxy"`
}

type logStore struct {
	entries []Entry
	mu      sync.Mutex
}

var store *logStore

func init() {
	store = new(logStore)
}

func Get() []Entry {
	return store.entries
}

func Put(entry Entry) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.entries = append(store.entries, entry)
}

func Delete() {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.entries = nil
}
