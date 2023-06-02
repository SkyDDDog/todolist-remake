package rpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
	"todolist-remake/config"
	"todolist-remake/idl/pb/task"
	"todolist-remake/idl/pb/user"
	"todolist-remake/pkg/discovery"
)

var (
	Register   *discovery.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc

	UserClient user.UserServiceClient
	TaskClient task.TaskServiceClient
)

const (
	userSrvName         = "user"
	taskSrvName         = "task"
	experimentalSrvName = "experimental"
)

func Init() {
	Register = discovery.NewResolver([]string{config.Etcd.Addr}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()

	initClient(config.GetService(userSrvName).Name, &UserClient)
	initClient(config.GetService(taskSrvName).Name, &TaskClient)
}

func initClient(serviceName string, client interface{}) {
	conn, err := connectServer(serviceName)

	if err != nil {
		panic(err)
	}

	switch c := client.(type) {
	case *user.UserServiceClient:
		*c = user.NewUserServiceClient(conn)
	case *task.TaskServiceClient:
		*c = task.NewTaskServiceClient(conn)
	default:
		panic("unsupported client types")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)

	// Load balance
	if config.GetService(serviceName).LB {
		log.Printf("load balance enabled for %s\n", serviceName)
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
	}

	return grpc.DialContext(ctx, addr, opts...)
}
