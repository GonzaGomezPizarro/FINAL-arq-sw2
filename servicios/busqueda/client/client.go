package cliente

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"strings"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/motordebusqueda"
)

// GetQuery realiza una búsqueda en Solr con laconsulta especificada
func GetQuery(query string) (dto.Items, error) {

	return dto.Items{}, nil
}

// GetAll trae todos los elementos del índice IndexName y los devuelve en dto.Items
func GetAll() (dto.Items, error) {
	// Realiza una solicitud de búsqueda para obtener todos los documentos
	var buf strings.Builder
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return dto.Items{}, fmt.Errorf("Error al codificar la consulta JSON: %v", err)
	}

	res, err := motordebusqueda.ElasticSearch.Search(
		motordebusqueda.ElasticSearch.Search.WithContext(context.Background()),
		motordebusqueda.ElasticSearch.Search.WithIndex(motordebusqueda.IndexName),
		motordebusqueda.ElasticSearch.Search.WithBody(strings.NewReader(buf.String())),
		motordebusqueda.ElasticSearch.Search.WithTrackTotalHits(true),
		motordebusqueda.ElasticSearch.Search.WithPretty(),
	)
	if err != nil {
		return dto.Items{}, fmt.Errorf("Error en la solicitud de búsqueda: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := ioutil.ReadAll(res.Body)
		return dto.Items{}, fmt.Errorf("Error en la respuesta de Elasticsearch (código %d): %s", res.StatusCode, body)
	}

	// Lee el cuerpo de la respuesta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return dto.Items{}, fmt.Errorf("Error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodifica el cuerpo JSON en la estructura dto.Items
	var items dto.Items
	if err := json.Unmarshal(body, &items); err != nil {
		return dto.Items{}, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
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
