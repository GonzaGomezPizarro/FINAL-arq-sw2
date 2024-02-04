package dto

// RespuestaElasticsearch representa la estructura de la respuesta de Elasticsearch
type RespuestaElasticsearch struct {
	Took int  `json:"took"`
	Hits Hits `json:"hits"`
}

// Hits representa la estructura de la sección "hits" en la respuesta de Elasticsearch
type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

// Total representa la estructura de la subsección "total" en la respuesta de Elasticsearch
type Total struct {
	Value int `json:"value"`
}

// Hit representa la estructura de un documento individual en la respuesta de Elasticsearch
type Hit struct {
	Source Item `json:"_source"`
}

// RespuestaCompletaElasticsearch representa la estructura completa de la respuesta de Elasticsearch
type RespuestaCompletaElasticsearch struct {
	Index       string `json:"_index"`
	Type        string `json:"_type"`
	ID          string `json:"_id"`
	Version     int    `json:"_version"`
	SeqNo       int    `json:"_seq_no"`
	PrimaryTerm int    `json:"_primary_term"`
	Found       bool   `json:"found"`
	Source      Item   `json:"_source"`
}
