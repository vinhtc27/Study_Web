package cache

import (
	"web-service/pkg/server"
)

// Initialize Function in Cache Package
func init() {
	// Local Cache Configuration Value
	localCacheCfg.NumCounters = server.Config.GetInt64("LOCAL_CACHE_NUM_COUNTERS")
	localCacheCfg.MaxCost = server.Config.GetInt64("LOCAL_CACHE_MAX_COST")
	localCacheCfg.BufferItems = server.Config.GetInt64("LOCAL_CACHE_BUFFER_ITEMS")
	localCacheCfg.Metrics = server.Config.GetBool("LOCAL_CACHE_METRICS")

	if localCacheCfg.MaxCost != 0 && localCacheCfg.BufferItems != 0 && localCacheCfg.NumCounters != 0 {
		// Do Redis Cache Connection
		LocalCache = localCacheConnect()
	}
}
