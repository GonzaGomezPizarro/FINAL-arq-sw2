package dto

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
)

type Item struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Country     string   `json:"country"`
	State       string   `json:"state"`
	City        string   `json:"city"`
	Address     string   `json:"address"`
	Photos      []string `json:"photos"`
	Price       int      `json:"price"`
	Bedrooms    int      `json:"bedrooms"`
	Bathrooms   int      `json:"bathrooms"`
	Mts2        int      `json:"mts2"`
	UserId      int      `json:"userId"`
}

type Items []Item

func (items Items) ConvertirImagenes() {
	for i := range items {
		items[i].ConvertirImagenes()
	}
}

func (item *Item) ConvertirImagenes() {
	var imagenesBase64 []string

	for _, rutaImagen := range item.Photos {
		imagenBase64, err := CargarImagen(rutaImagen)
		if err != nil {
			log.Printf("Error al cargar la imagen desde la ruta %s: %s", rutaImagen, err.Error())
			continue
		}
		imagenesBase64 = append(imagenesBase64, imagenBase64)
	}

	item.Photos = imagenesBase64
}

func CargarImagen(ruta string) (string, error) {
	// Abrir el archivo de imagen
	archivo, err := os.Open(ruta)
	if err != nil {
		return "", err
	}
	defer archivo.Close()

	// Leer el contenido del archivo
	datos, err := ioutil.ReadAll(archivo)
	if err != nil {
		return "", err
	}

	// Codificar los datos de la imagen en Base64
	imagenBase64 := base64.StdEncoding.EncodeToString(datos)

	return imagenBase64, nil
}
