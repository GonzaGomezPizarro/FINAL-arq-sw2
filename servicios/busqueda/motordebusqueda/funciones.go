package motordebusqueda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
)

// IndexAll obtiene todos los items del servicio de items y los indexa en Elasticsearch
func IndexAll() error {
	// Obtener todos los items del servicio de items

	DELETEItems()

	items, err := getAllItemsFromService()
	if err != nil {
		return err
	}

	if items == nil {
		log.Println("-> Base de datos vacia")
		return nil
	}

	// Indexar los items en Elasticsearch
	for _, item := range items {
		id := item.Id
		err := indexDocument(id, item)
		if err != nil {
			return err
		}
	}

	log.Println("-> Items indexados")

	return nil
}

// getAllItemsFromService obtiene todos los items del servicio de items
func getAllItemsFromService() (dto.Items, error) {
	resp, err := http.Get("http://balanceador_items:8090/items")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		return nil, nil
	}

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

	// Construir la URL para la indexación
	url := fmt.Sprintf("%s/%s/_doc/%s", URL, IndexName, id)

	// Realizar la solicitud HTTP POST para indexar el documento
	resp, err := http.Post(url, "application/json", bytes.NewReader(itemJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Verificar si la respuesta indica un error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error al indexar el documento. código de estado: %d. respuesta: %s", resp.StatusCode, string(body))
	}

	return nil
}

// DELETEItems envía una solicitud DELETE a la URL /items
func DELETEItems() {
	// Definir la URL a la que enviar la solicitud DELETE
	url := URL + "/" + IndexName

	// Crear una solicitud HTTP DELETE
	req, _ := http.NewRequest("DELETE", url, nil)

	// Realizar la solicitud HTTP DELETE
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud DELETE:", err)
		return // Sale de la función si hay un error, pero la ejecución continúa
	}
	defer resp.Body.Close()

	// Comprobar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		fmt.Println("La solicitud DELETE no se completó correctamente. Código de estado:", resp.StatusCode)
		return // Sale de la función si el código de estado no es 200 OK, pero la ejecución continúa
	}

	// Si llega aquí, la solicitud DELETE se realizó correctamente
	fmt.Println("La solicitud DELETE se realizó correctamente.")
}

//----------------------------------------------------------------
// Revisar busca en la base de datos el item y lo elimina o lo agrega en la colección de Elasticsearch.
// Los items en la base de datos solo pueden ser creados o borrados, no se pueden modificar campos...
func Revisar(id string) error {
	// Obtener información del ítem desde el servicio de items
	itemJSON, err := obtenerJSONItem(id)
	if err != nil {
		return err
	}

	if itemJSON == nil {
		// El item no se encuentra en la base de datos
		println()
		errr := deleteDocument(id)
		if errr != nil {
			return errr
		}
	} else {
		// Convertir el JSON a una estructura dto.Item
		var itemsDto dto.Items
		var itemDto dto.Item
		err := json.Unmarshal(itemJSON, &itemsDto)
		if err != nil {
			return err
		}
		itemDto = itemsDto[0] // PING ---------------------------------------------------------

		// Indexar o actualizar el ítem en Elasticsearch
		err = indexDocument(id, itemDto)
		if err != nil {
			return err
		}
	}

	return nil
}

// obtenerJSONItem realiza una solicitud HTTP GET al servicio de items para obtener el JSON del ítem.
func obtenerJSONItem(id string) ([]byte, error) {
	// Construir la URL para obtener el JSON del ítem
	url := "http://balanceador_items:8090/item/" + id
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Verificar el código de estado de la respuesta
	if resp.StatusCode == http.StatusNotFound {
		// El item no se encuentra en la base de datos
		return nil, nil
	} else if resp.StatusCode == http.StatusOK {
		// Se encontró el item, devolvemos el JSON del ítem
		return body, nil
	}

	return nil, fmt.Errorf("error al obtener el json del ítem. código de estado: %d. respuesta: %s", resp.StatusCode, string(body))
}

func deleteDocument(id string) error {
	// Construir la URL para eliminar el documento
	url := fmt.Sprintf("%s/%s/_doc/%s", URL, IndexName, id)

	// Realizar la solicitud HTTP DELETE para eliminar el documento
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Realizar la solicitud HTTP DELETE
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Verificar si la respuesta indica un error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error al eliminar el documento. código de estado: %d. respuesta: %s", resp.StatusCode, string(body))
	}

	return nil
}
