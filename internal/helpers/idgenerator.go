package helpers

import "github.com/google/uuid"


func GenerateId(opts ...string) string {
	idGenerator := UuidV4{}
	return idGenerator.GenerateId()
}


type IdGeneratorInterface interface {
	GenerateId() string
}

type UuidV4 struct{
}

func (id UuidV4)  GenerateId() string {
	uuid := uuid.New()
	return uuid.String()
}