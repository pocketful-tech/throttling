package helpers

import (
	"sync"
	"time"
)

type APIThrottler struct {
	RequestsPerSecond int
	LastRequestTime   time.Time
	Mu                sync.Mutex
	ClientName        string
	ClientIp          string
	VendorName        string
}

func (t *APIThrottler) Throttle() {
	t.Mu.Lock()
	defer t.Mu.Unlock()

	currentTime := time.Now()
	timeSinceLastRequest := currentTime.Sub(t.LastRequestTime)
	desiredInterval := time.Second / time.Duration(t.RequestsPerSecond)

	if timeSinceLastRequest < desiredInterval {
		sleepDuration := desiredInterval - timeSinceLastRequest
		time.Sleep(sleepDuration)
	}

	t.LastRequestTime = time.Now()
}
