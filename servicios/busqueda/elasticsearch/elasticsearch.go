package elasticsearch

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// Indexall obtiene todos los items del servicio de items y los indexa en Elasticsearch
func Indexall() error {
	// Obtener todos los items del servicio de items
	itemsJSON, err := getAllItemsJSONFromService()
	if err != nil {
		return err
	}

	// Conectar a Elasticsearch
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}

	// Indexar el JSON en Elasticsearch
	indexDocument(es, "items", itemsJSON)

	return nil
}

// getAllItemsJSONFromService obtiene el JSON de todos los items del servicio de items
func getAllItemsJSONFromService() ([]byte, error) {
	resp, err := http.Get("http://localhost:8091/items")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// indexDocument indexa un documento JSON en Elasticsearch
func indexDocument(es *elasticsearch.Client, indexName string, documentJSON []byte) error {
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: "your_document_id", // Puedes proporcionar tu propio ID o dejar que Elasticsearch genere uno automáticamente.
		Body:       bytes.NewReader(documentJSON),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error al indexar el documento: %s", res.String())
	}

	return nil
}

// Actualizar busca en la base de datos el item y lo actualiza en la colección de Elasticsearch.
func Actualizar(id string) error {
	// Obtener información del ítem desde el servicio de items
	itemJSON, err := obtenerJSONItem(id)
	if err != nil {
		return err
	}

	// Actualizar el ítem en Elasticsearch
	err = actualizarItemEnElasticsearch(itemJSON)
	if err != nil {
		return err
	}

	return nil
}

// obtenerJSONItem realiza una solicitud HTTP GET al servicio de items para obtener el JSON del ítem.
func obtenerJSONItem(id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8091/item/%s", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// ActualizarItemEnElasticsearch actualiza el ítem en Elasticsearch.
func ActualizarItemEnElasticsearch(itemJSON []byte) error {
	// Configurar la conexión a Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	// Decodificar el JSON del ítem
	var item map[string]interface{}
	if err := json.Umarshal(itemJSON, &item); err != nil {
		return err
	}

	// Obtener el ID del ítem
	id, ok := item["id"].(string)
	if !ok {
		return fmt.Errrf("No se pudo obtener el ID del ítem")
	}

	// Preparar la solicitd de índice
	req := elasticsearch.IndexRequest{
		Index:      "tu-indice", // Ajustar segúntu índice
		DocumentID: id,          // Utilizar el ID como Document ID
		Body:       bytes.NewReader(itemJSON),
		efresh:     "true",
	}

	// Enviar la solicitud a Elasticsearch
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Verificar la respuesta de Elasticsearch
	if res.IsError() {
		return fmt.Errorf("Error al actulizar el ítem en Elasticsearch: %s", res.String())
	}

	return nil

}

// GetQuery realiza una búsqueda en Solr con laconsulta especificada
func GetQuery(query string) (dto.Items, error) {

	return dto.Items{}, nil
}

// GetAll realiza una consulta para obtener todos los elementos indexados en Solr
func GetAll() (dto.Items, error) {

	return dto.Items{}, nil
}

// GetItemById realiza una consulta para obtener un elemento específico por su ID desde Solr
func GetItemById(id string) (dto.Item, error) {

	return dto.Item{}, nil
}
