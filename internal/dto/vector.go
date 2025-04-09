package dto

// Vector is a struct that represents a vector in the database.
type Vector struct {
	ID         string    `json:"id"`
	VectorData []float64 `json:"vector_data"`
}

type VectorWithMetadata struct {
	Vector
	Metadata map[string]interface{} `json:"metadata"`
}
