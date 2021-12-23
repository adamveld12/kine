package consul

import (
	"github.com/hashicorp/consul/api"
)

func New(ctx context.Context, driverName, dataSourceName string) (server.Backend, error) {

}


func NewVariant(ctx context.Context, driverName, dataSourceName string) (server.Backend, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}


}


type consulServerBackend struct {
	client *api.Client

}
