package helpers

import "github.com/google/uuid"


// 


//Global function to generate a unique id
func GenerateId(opts ...string) string {
	idGenerator := UuidV4{}
	return idGenerator.GenerateId()
}

//Global function to generate a unique id with a prefix
func GenerateIdWithPrefix(prefix string) string {
	idGenerator := UuidV4{}
	return idGenerator.GenerateIdWithPrefix(prefix)
}


//........................ Interface ...............................
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