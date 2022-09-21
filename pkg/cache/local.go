package cache

import (
	"sync"
	"time"
	"web-service/pkg/log"

	"github.com/dgraph-io/ristretto"
)

// Default cache item cost
const _defaultCacheItemCost = 1

// Local Cache Configuration Struct
type localCacheConfig struct {
	MaxCost     int64
	NumCounters int64
	BufferItems int64
	Metrics     bool
}

// Local Cache Configuration Variable
var localCacheCfg localCacheConfig

type LocalCacheStore struct {
	mutex sync.RWMutex
	store *ristretto.Cache
}

// Local Cache Variable
var LocalCache *LocalCacheStore

// Local Cache Connect Function
func localCacheConnect() *LocalCacheStore {
	// Initialize Connection
	cache, err := ristretto.NewCache(&ristretto.Config{
		MaxCost:     localCacheCfg.MaxCost,
		NumCounters: localCacheCfg.NumCounters,
		BufferItems: localCacheCfg.BufferItems,
		Metrics:     localCacheCfg.Metrics,
	})
	if err != nil {
		log.Println(log.LogLevelFatal, "local cache init failed:", err.Error())
	}

	// Return Connection
	return &LocalCacheStore{
		store: cache,
	}
}

// Get method to retrieve the value of a key. If not present, returns false.
func (localCacheStore *LocalCacheStore) Get(key string) (any, bool) {
	return localCacheStore.store.Get(key)
}

// SetByKey method to set cache by given key with time to live
func (localCacheStore *LocalCacheStore) SetByKey(key string, value any, timeToLive time.Duration) {
	localCacheStore.mutex.Lock()
	defer localCacheStore.mutex.Unlock()

	localCacheStore.store.SetWithTTL(key, value, _defaultCacheItemCost, timeToLive)
	localCacheStore.store.Wait()
}

// InvalidateByKey  method to delete a key from cahce.
func (localCacheStore *LocalCacheStore) InvalidateByKey(key string) {
	localCacheStore.mutex.Lock()
	defer localCacheStore.mutex.Unlock()

	localCacheStore.store.Del(key)
	localCacheStore.store.Wait()
}

// SetByTags method to set cache by given tags with time to live
func (localCacheStore *LocalCacheStore) SetByTags(key string, value any, timeToLive time.Duration, tags []string) {
	localCacheStore.mutex.Lock()
	defer localCacheStore.mutex.Unlock()

	for _, tag := range tags {
		tagSet := NewTagSet()
		v, found := localCacheStore.store.Get(tag)
		if found {
			tagSet = v.(*TagSet)
		}
		tagSet.Add(key)
		localCacheStore.store.Set(tag, tagSet, _defaultCacheItemCost)
	}

	localCacheStore.store.SetWithTTL(key, value, _defaultCacheItemCost, timeToLive)
	localCacheStore.store.Wait()
}

// InvalidateByTags method to invalidate cache with given tags.
func (localCacheStore *LocalCacheStore) InvalidateByTags(tags []string) {
	localCacheStore.mutex.Lock()
	defer localCacheStore.mutex.Unlock()

	keys := make([]string, 0)
	for _, tag := range tags {
		tagSet := NewTagSet()
		v, found := localCacheStore.store.Get(tag)
		if found {
			tagSet = v.(*TagSet)
		}
		keys = append(keys, tagSet.Members()...)
		keys = append(keys, tag)
	}

	for _, k := range keys {
		localCacheStore.store.Del(k)
	}
	localCacheStore.store.Wait()
}

func (localCacheStore *LocalCacheStore) MetricsString() string {
	return localCacheStore.store.Metrics.String()
}

// Close method clear and then close the cache store.
func (localCacheStore *LocalCacheStore) Close() {
	localCacheStore.store.Wait()
	localCacheStore.store.Clear()
	localCacheStore.store.Close()
}
