package behavioral

type Cache struct {
	storage map[string]string
	algo EvictionAlgo
	capacity int
	maxCapacity int
}

func initCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:     make(map[string]string),
		algo:        e,
		capacity:    0,
		maxCapacity: 10,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.algo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity > c.maxCapacity {
		c.algo.evict(c)
	}
	c.storage[key] = value
	c.capacity++
}

func (c *Cache) get(key string) string {
	return c.storage[key]
}

func (c *Cache) evict() {
	c.algo.evict(c)
	c.capacity--
}