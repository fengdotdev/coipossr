package helpers

import "github.com/google/uuid"


func GenerateId(opts ...string) string {
	idGenerator := UuidV4{}
	return idGenerator.GenerateId()
}


func GenerateIdWithPrefix(prefix string) string {
	idGenerator := UuidV4{}
	return idGenerator.GenerateIdWithPrefix(prefix)
}

type IdGeneratorInterface interface {
	GenerateId() string
	GenerateIdWithPrefix(prefix string) string
}

type UuidV4 struct{
}

func (id UuidV4)  GenerateId() string {
	uuid := uuid.New()
	return uuid.String()
}

func (id UuidV4)  GenerateIdWithPrefix(prefix string) string {
	uuid := uuid.New()
	return prefix + "-" + uuid.String()
}