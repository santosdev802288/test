package kubgo

import "dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"

type IKubgoFinder interface {
	GetAll() chan KubgosResponse
	Get(id uuid.UUID) chan *KubgoResponse
}
