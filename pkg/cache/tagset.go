package cache

import "sync"

// New TagSet.
func NewTagSet() *TagSet {
	return &TagSet{
		data: make(map[string]struct{}),
	}
}

// TagSet is our struct that acts as a set data structure
// with string as membertagSet.
type TagSet struct {
	mutex sync.RWMutex
	data  map[string]struct{}
}

// Add method to add a member to the TagSet.
func (tagSet *TagSet) Add(member string) {
	tagSet.mutex.Lock()
	defer tagSet.mutex.Unlock()

	tagSet.data[member] = struct{}{}
}

// Remove method to remove a member from the TagSet.
func (tagSet *TagSet) Remove(member string) {
	tagSet.mutex.Lock()
	defer tagSet.mutex.Unlock()

	delete(tagSet.data, member)
}

// IsMember method to check if a member is present in the TagSet.
func (tagSet *TagSet) IsMember(member string) bool {
	tagSet.mutex.RLock()
	defer tagSet.mutex.RUnlock()

	_, found := tagSet.data[member]
	return found
}

// Members method to retrieve all members of the TagSet.
func (tagSet *TagSet) Members() []string {
	tagSet.mutex.RLock()
	defer tagSet.mutex.RUnlock()

	keys := make([]string, 0)
	for k := range tagSet.data {
		keys = append(keys, k)
	}
	return keys
}

// Size method to get the cardinality of the TagSet.
func (tagSet *TagSet) Size() int {
	tagSet.mutex.RLock()
	defer tagSet.mutex.RUnlock()

	return len(tagSet.data)
}

// Clear method to remove all members from the TagSet.
func (tagSet *TagSet) Clear() {
	tagSet.mutex.Lock()
	defer tagSet.mutex.Unlock()

	tagSet.data = make(map[string]struct{})
}
