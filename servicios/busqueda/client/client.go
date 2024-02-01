package cliente

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/motordebusqueda"
)

// GetQuery realiza una búsqueda en Elasticsearch con la consulta especificada
func GetQuery(query string) (dto.Items, error) {
	// Construye la URL de la consulta
	encodedQuery := url.QueryEscape(query)
	queryURL := fmt.Sprintf("http://localhost:9200/items/_search?q=%s", encodedQuery)

	// Realiza una solicitud HTTP GET directa a la API de Elasticsearch con la consulta
	resp, err := http.Get(queryURL)
	if err != nil {
		return dto.Items{}, fmt.Errorf("Error al realizar la solicitud HTTP a Elasticsearch: %v", err)
	}
	defer resp.Body.Close()

	// Lee el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Items{}, fmt.Errorf("Error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodifica el cuerpo JSON en la estructura RespuestaElasticsearch
	var respuesta dto.RespuestaElasticsearch
	if err := json.Unmarshal(body, &respuesta); err != nil {
		return dto.Items{}, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}

	// Verifica que la respuesta tenga contenido
	if respuesta.Hits.Total.Value == 0 {
		return dto.Items{}, fmt.Errorf("No se encontraron documentos en Elasticsearch para la consulta: %s", query)
	}

	// Construye la lista de items a partir de los hits
	var items dto.Items
	for _, hit := range respuesta.Hits.Hits {
		items = append(items, hit.Source)
	}

	return items, nil
}

// GetAll trae todos los elementos del índice IndexName y los devuelve en dto.Items
func GetAll() (dto.Items, error) {
	// Realiza una solicitud HTTP GET directa a la API de Elasticsearch
	resp, err := http.Get("http://localhost:9200/items/_search?q=*")
	if err != nil {
		return dto.Items{}, fmt.Errorf("Error al realizar la solicitud HTTP a Elasticsearch: %v", err)
	}
	defer resp.Body.Close()

	// Lee el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Items{}, fmt.Errorf("Error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodifica el cuerpo JSON en la estructura RespuestaElasticsearch
	var respuesta dto.RespuestaElasticsearch
	if err := json.Unmarshal(body, &respuesta); err != nil {
		return dto.Items{}, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}

	// Verifica que la respuesta tenga contenido
	if respuesta.Hits.Total.Value == 0 {
		return dto.Items{}, fmt.Errorf("No se encontraron documentos en Elasticsearch")
	}

	// Construye la lista de items a partir de los hits
	var items dto.Items
	for _, hit := range respuesta.Hits.Hits {
		items = append(items, hit.Source)
	}

	return items, nil
}

// GetItemById obtiene el documento que coincide con el ID dado del índice IndexName
func GetItemById(id string) (dto.Item, error) {
	// Realiza una solicitud de Elasticsearch para obtener el documento por ID
	res, err := motordebusqueda.ElasticSearch.Get(motordebusqueda.IndexName, id)
	if err != nil {
		return dto.Item{}, fmt.Errorf("Error en la solicitud de Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := ioutil.ReadAll(res.Body)
		return dto.Item{}, fmt.Errorf("Error en la respuesta de Elasticsearch (código %d): %s", res.StatusCode, body)
	}

	// Lee el cuerpo de la respuesta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return dto.Item{}, fmt.Errorf("Error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodifica el cuerpo JSON en la estructura dto.Item
	var item dto.Item
	if err := json.Unmarshal(body, &item); err != nil {
		return dto.Item{}, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}

	return item, nil
}
