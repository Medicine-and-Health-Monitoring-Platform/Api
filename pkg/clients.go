package pkg

import (
	"Api_Gateway/config"
	pb "Api_Gateway/genproto/health_analytics"
	pbu "Api_Gateway/genproto/users"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(cfg *config.Config) pbu.AuthServiceClient {
	conn, err := grpc.NewClient(cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbu.NewAuthServiceClient(conn)
}

func NewAdminClient(cfg *config.Config) pbu.AdminClient {
	conn, err := grpc.NewClient(cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbu.NewAdminClient(conn)
}

func NewHealthClient(cfg *config.Config) pb.HealthAnalyticsServiceClient {
	conn, err := grpc.NewClient(cfg.HEALTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}
	return pb.NewHealthAnalyticsServiceClient(conn)
}
