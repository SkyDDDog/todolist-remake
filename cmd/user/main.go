package main

import (
	"flag"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"todolist-remake/cmd/user/dal"
	"todolist-remake/config"
	"todolist-remake/pkg/discovery"
	"todolist-remake/pkg/utils"

	service "todolist-remake/idl/pb/user"
)

var (
	path   *string
	srvnum *int
)

func Init() *discovery.Register {
	// Args
	path = flag.String("config", "./config", "config path")
	srvnum = flag.Int("srvnum", 0, "node number")
	flag.Parse()

	config.Init(*path, srvname, *srvnum)
	dal.Init()

	// etcd
	register := discovery.NewRegister([]string{config.Etcd.Addr}, logrus.New())
	// defer register.Stop()

	node := discovery.Server{
		Name: config.Service.Name,
		Addr: config.Service.Addr,
	}

	if _, err := register.Register(node, 10); err != nil {
		log.Fatalf("register service %s failed, err: %v", node.Name, err)
	}

	return register
}

func main() {
	register := Init()

	lis, err := net.Listen("tcp", config.Service.Addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	service.RegisterUserServiceServer(s, NewUserServiceImpl())
	reflection.Register(s) // Support server reflection

	go utils.ListenSignal(func() {
		register.Stop()
		s.Stop()
	})

	log.Printf("%s listening at %v (node number: %d)\n", srvname, lis.Addr(), *srvnum)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
