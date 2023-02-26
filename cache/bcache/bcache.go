package bcache

import "sync"

type BCache struct {
}

type bucket struct {
	l sync.Mutex
}
