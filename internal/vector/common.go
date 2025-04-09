package vector

import "github.com/Phlosy/myVectorDB/internal/dto"

type DistanceResult struct {
	Distance float64
	Vector   *dto.Vector
}
