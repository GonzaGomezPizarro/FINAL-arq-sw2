package motordebusqueda

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

const IndexName = "items"

var ElasticSearch *elasticsearch.Client

func init() {
	// Inicia la conexión con Elasticsearch
	if err := StartSearchEngine(); err != nil {
		fmt.Printf("Error al iniciar el motor de búsqueda: %s\n", err)
	}
}

// StartSearchEngine inicia la conexión con Elasticsearch
func StartSearchEngine() error {
	// Configurar la conexión a Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	var err error
	ElasticSearch, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("Error al crear el cliente de Elasticsearch: %s", err)
	}

	return nil
}
