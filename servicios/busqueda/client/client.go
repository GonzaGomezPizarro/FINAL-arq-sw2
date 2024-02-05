package cliente

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
)

// GetQuery realiza una búsqueda en Elasticsearch con la consulta especificada
func GetQuery(query string) (dto.Items, error) {
	// Construye la URL de la consulta
	encodedQuery := url.QueryEscape(query)
	queryURL := fmt.Sprintf("http://localhost:9200/items/_search?q=%s*", encodedQuery)

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
	url := "http://localhost:9200/items/_doc/" + id

	resp, err := http.Get(url)
	if err != nil {
		return dto.Item{}, fmt.Errorf("Error al realizar la solicitud HTTP a Elasticsearch: %v", err)
	}
	defer resp.Body.Close()

	// Verifica el código de respuesta
	if resp.StatusCode != http.StatusOK {
		return dto.Item{}, fmt.Errorf("Error al obtener el documento de Elasticsearch. Código de estado: %d", resp.StatusCode)
	}

	// Lee el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Item{}, fmt.Errorf("Error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodifica el cuerpo JSON en la estructura RespuestaCompletaElasticsearch
	var respuestaCompleta dto.RespuestaCompletaElasticsearch
	if err := json.Unmarshal(body, &respuestaCompleta); err != nil {
		return dto.Item{}, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}

	// Devuelve directamente el campo Source de RespuestaCompletaElasticsearch
	return respuestaCompleta.Source, nil
}

func GetItemByUserId(userId int) (dto.Items, error) {
	// Construir la URL de la consulta
	queryURL := fmt.Sprintf("http://localhost:9200/items/_search?q=userId:%d", userId)

	// Realizar una solicitud HTTP GET directa a la API de Elasticsearch con la consulta
	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, fmt.Errorf("Error al realizar la solicitud HTTP a Elasticsearch: %v", err)
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodificar el cuerpo JSON en la estructura RespuestaElasticsearch
	var respuesta map[string]interface{}
	if err := json.Unmarshal(body, &respuesta); err != nil {
		return nil, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}

	// Verificar que la respuesta tenga contenido
	hits, ok := respuesta["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("No se encontraron documentos en Elasticsearch para la consulta")
	}

	// Construir la lista de items a partir de los hits
	var items dto.Items
	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		var item dto.Item

		// Utilizar json.Unmarshal para convertir automáticamente otros campos
		itemData, err := json.Marshal(source)
		if err != nil {
			return nil, fmt.Errorf("Error al convertir el item: %v", err)
		}
		if err := json.Unmarshal(itemData, &item); err != nil {
			return nil, fmt.Errorf("Error al convertir el item: %v", err)
		}

		items = append(items, item)
	}

	return items, nil
}
