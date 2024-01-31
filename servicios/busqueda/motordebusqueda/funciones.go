package motordebusqueda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// IndexAll obtiene todos los items del servicio de items y los indexa en Elasticsearch
func IndexAll() error {
	// Obtener todos los items del servicio de items
	items, err := getAllItemsFromService()
	if err != nil {
		return err
	}

	// Indexar los items en Elasticsearch
	for _, item := range items {
		id := item.Id
		indexDocument(id, item)
	}

	return nil
}

// getAllItemsFromService obtiene todos los items del servicio de items
func getAllItemsFromService() (dto.Items, error) {
	resp, err := http.Get("http://localhost:8091/items")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var items dto.Items

	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// indexDocument indexa un documento JSON en Elasticsearch
func indexDocument(id string, item dto.Item) error {
	// Convertir el item a JSON
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      IndexName,
		DocumentID: id,
		Body:       bytes.NewReader(itemJSON),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), ElasticSearch)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error al indexar el documento: %s", res.String())
	}

	return nil
}

//----------------------------------------------------------------
// Actualizar busca en la base de datos el item y lo actualiza en la colección de Elasticsearch.
func Actualizar(id string) error {
	// Obtener información del ítem desde el servicio de items
	itemJSON, err := obtenerJSONItem(id)
	if err != nil {
		return err
	}

	// Actualizar el ítem en Elasticsearch
	err = ActualizarItemEnElasticsearch(itemJSON)
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
	// Decodificar el JSON del ítem
	var item map[string]interface{}
	if err := json.Unmarshal(itemJSON, &item); err != nil {
		return err
	}

	// Obtener el ID del ítem
	id, ok := item["id"].(string)
	if !ok {
		return fmt.Errorf("No se pudo obtener el ID del ítem")
	}

	// Preparar la solicitud de actualización
	req := esapi.UpdateRequest{
		Index:      IndexName,
		DocumentID: id,
		Body:       strings.NewReader(fmt.Sprintf(`{"doc": %s}`, itemJSON)),
		Refresh:    "true",
	}

	// Enviar la solicitud a Elasticsearch usando la conexión global
	res, err := req.Do(context.Background(), ElasticSearch)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Verificar la respuesta de Elasticsearch
	if res.IsError() {
		return fmt.Errorf("Error al actualizar el ítem en Elasticsearch: %s", res.String())
	}

	return nil
}
