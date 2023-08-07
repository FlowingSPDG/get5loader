package uuid

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() string
}

type uuidGenerator struct {
}

func NewUUIDGenerator() UUIDGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) Generate() string {
	id := uuid.New()
	return id.String()
}
