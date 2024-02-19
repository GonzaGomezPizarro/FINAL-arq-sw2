package motordebusqueda

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	IndexName = "items"
	URL       = "http://elasticsearch:9200"
)

func init() {
	// Inicia la conexión con Elasticsearch
	if err := Check(); err != nil {
		fmt.Printf("Error al iniciar el motor de búsqueda: %s\n", err)
	}

}

// StartSearchEngine inicia la conexión con Elasticsearch
func Check() error {

	err := checkConection()
	if err != nil {
		log.Println(" -> No se pudo establecer conexion con el motor de búsqueda")
		return err
	}

	fmt.Println("Elasticsearch healthy.")
	return nil
}

func checkConection() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}

	var er error = fmt.Errorf(strconv.Itoa(resp.StatusCode))
	return er
}
