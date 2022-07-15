package cache

import "time"

type Data struct {
	value    string
	deadline *time.Time
}

type Cache struct {
	Map map[string]Data
}

func NewCache() Cache {
	return Cache{
		Map: map[string]Data{},
	}
}

func NewData(value string, deadline *time.Time) *Data {
	return &Data{value, deadline}
}

func (cache Cache) Get(key string) (string, bool) {
	cache.cleanup()
	value, err := cache.Map[key]
	return value.value, err
}

func (cache *Cache) Put(key, value string) {
	cache.Map[key] = *NewData(value, nil)
}

func (cache Cache) Keys() []string {
	cache.cleanup()
	keys := make([]string, 0)
	for k := range cache.Map {
		keys = append(keys, k)
	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.Map[key] = *NewData(value, &deadline)
}

func (cache *Cache) cleanup() {
	current_time := time.Now()
	for k, v := range cache.Map {
		if v.deadline != nil && v.deadline.Before(current_time) {
			delete(cache.Map, k)
		}
	}
}
