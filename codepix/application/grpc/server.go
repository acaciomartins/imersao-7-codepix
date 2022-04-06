package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/acaciomartins/imersao-7-codepix/codepix-go/application/grpc/pb"
	"github.com/acaciomartins/imersao-7-codepix/codepix-go/infrastructure/repository"
	"github.com/codeedu/imersao/codepix-go/application/usecase"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}

func NewPixGrpcService(pixUseCase invalid type) {
	panic("unimplemented")
}
