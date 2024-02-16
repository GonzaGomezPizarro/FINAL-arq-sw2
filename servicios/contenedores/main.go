package main

import (
	"fmt"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/contenedores/docker"
)

func main() {
	// Prueba de la función GetImages
	fmt.Println("Obteniendo imágenes:")
	images, err := docker.GetImages()
	if err != nil {
		fmt.Printf("Error al obtener imágenes: %v\n", err)
	} else {
		for _, image := range images {
			fmt.Println(image.Name)
		}
	}

	// Prueba de la función GetContenedores
	fmt.Println("\nObteniendo contenedores:")
	contenedores, err := docker.GetContenedores()
	if err != nil {
		fmt.Printf("Error al obtener contenedores: %v\n", err)
	} else {
		for _, contenedor := range contenedores {
			fmt.Printf("ID: %s, Name: %s, Image: %s, Status: %s, InternalPort: %d, ExternalPort: %d\n", contenedor.Id, contenedor.Name, contenedor.Imagen.Name, contenedor.Status, contenedor.InternalPort, contenedor.ExternalPort)
		}
	}

	// Prueba de la función PostContenedor
	fmt.Println("\nCreando nuevo contenedor:")
	newContenedor, err := docker.PostContenedor("servicios-items", "", 0, 0)
	if err != nil {
		fmt.Printf("Error al crear nuevo contenedor: %v\n", err)
	} else {
		fmt.Printf("Contenedor creado: ID: %s, Name: %s, Image: %s, Status: %s, InternalPort: %d, ExternalPort: %d\n", newContenedor.Id, newContenedor.Name, newContenedor.Imagen.Name, newContenedor.Status, newContenedor.InternalPort, newContenedor.ExternalPort)
	}
}
