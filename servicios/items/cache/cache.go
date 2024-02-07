package cache

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	memcacheClient *memcache.Client
)

func init() {
	// Configura la conexión a Memcached
	memcacheClient = memcache.New("memcached:11211")

	// Verificar la conexión a Memcached
	err := memcacheClient.Ping()
	if err != nil {
		log.Fatal("Error conectando a Memcached:", err)
	}

	log.Println("Connected to Memcached")
}

// SetToCache almacena datos en Memcached
func SetToCache(key string, value string) error {
	return memcacheClient.Set(&memcache.Item{
		Key:   key,
		Value: []byte(value),
		//Expiration: 600, // 10 minutos, default
	})
}

// GetFromCache obtiene datos de Memcached
func GetFromCache(key string) (string, error) {
	item, err := memcacheClient.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

// DeleteFromCache borra datos de Memcached
func DeleteFromCache(key string) error {
	return memcacheClient.Delete(key)
}
