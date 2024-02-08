package cacheLocal

import (
	"sync"
	"time"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/dto"
)

// CacheLocal es una estructura que representa una caché local.
type CacheLocal struct {
	data map[string]cacheItem
	mu   sync.RWMutex
}

type cacheItem struct {
	Item      dto.Item
	ExpiresAt time.Time
}

var (
	cacheInstance *CacheLocal
	once          sync.Once
)

// NewCache crea una nueva instancia de caché.
func NewCache() *CacheLocal {
	return &CacheLocal{
		data: make(map[string]cacheItem),
	}
}

// InitCache inicializa la caché y almacena su referencia en una variable global.
func InitCache() {
	once.Do(func() {
		cacheInstance = NewCache()
	})
}

// Set establece un valor en la caché para una clave dada.
// Retorna true si se logró insertar el valor en la caché, de lo contrario, retorna false.
func (c *CacheLocal) Set(item dto.Item) bool {
	key := item.Id
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[key]; !ok {
		c.data[key] = cacheItem{Item: item, ExpiresAt: time.Now().Add(30 * time.Second)}
		return true
	}
	return false
}

// Get devuelve el valor almacenado en la caché para una clave dada.
// Si el ítem ha expirado, se elimina de la caché y se retorna false.
func (c *CacheLocal) Get(key string) (dto.Item, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if item, ok := c.data[key]; ok {
		if time.Now().Before(item.ExpiresAt) {
			return item.Item, true
		} else {
			// El ítem ha expirado, eliminarlo de la caché
			delete(c.data, key)
		}
	}
	return dto.Item{}, false
}
