package handler

import (
	"Api_Gateway/config"
	pb "Api_Gateway/genproto/users"
	"Api_Gateway/pkg"
	"Api_Gateway/pkg/logger"
	"log/slog"
	"time"
)

type str string

type Handler struct {
	User           pb.AuthServiceClient
	Admin          pb.AdminClient
	Log            *slog.Logger
	ContextTimeout time.Duration
	UserIDKey      str
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		User:           pkg.NewUserClient(cfg),
		Admin:          pkg.NewAdminClient(cfg),
		Log:            logger.NewLogger(),
		ContextTimeout: 5 * time.Second,
		UserIDKey:      str("user_id"),
	}
}
