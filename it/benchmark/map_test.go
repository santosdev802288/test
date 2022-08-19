//go:build bench
// +build bench

package benchmark__test

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"github.com/mashingan/smapping"
	"github.com/stretchr/testify/assert"
	"siigo.com/kubgo/src/domain/kubgo"
	"testing"
)

func MapGrpcToDomain(b *testing.B) {

	id := cqrs.NewUUID()
	ct := kubgo.NewKubgo(id)

	b.StartTimer()

	grpcModel := &kubgo.Kubgo{
		Id: id,
	}

	mapped := smapping.MapFields(grpcModel)
	smapping.FillStruct(ct, mapped)

	b.StopTimer()

	assert.Equal(b, ct.Id, id)
}

func BenchmarkMapGrpcToDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapGrpcToDomain(b)
	}
}
