package dto

type Imagen struct {
	Name string `json:"name"`
}

type Imagenes []Imagen

type Contenedor struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	Imagen       Imagen `json:"imagen"`
	Status       string `json:"status"`
	InternalPort int    `json:"internal_port"`
	ExternalPort int    `json:"external_port"`
}

type Contenedores []Contenedor
