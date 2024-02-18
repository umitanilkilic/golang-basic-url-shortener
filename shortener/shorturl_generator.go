package shortener

import (
	"github.com/google/uuid"
)

func GenerateID() uint32 {
	return uuid.New().ID()
}
