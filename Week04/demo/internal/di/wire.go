// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"Go-000/Week04/demo/internal/dao"
	"Go-000/Week04/demo/internal/service"
	"Go-000/Week04/demo/internal/server/grpc"
	"Go-000/Week04/demo/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
