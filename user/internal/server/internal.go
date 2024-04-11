package server

import (
	"errors"
	"github.com/MaksKazantsev/grpc_service/user/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

const (
	internal string = "internal error"
)

func (s *server) handleError(err error) error {
	var e *models.Error

	ok := errors.As(err, &e)

	if ok {
		switch e.Status {
		case http.StatusInternalServerError:
			s.log.Error(e.Message)
			return status.Error(codes.Internal, internal)
		case http.StatusBadRequest:
			s.log.Error(e.Message)
			return status.Error(codes.InvalidArgument, e.Message)
		case http.StatusNotFound:
			return status.Error(codes.NotFound, e.Message)
		}
	}

	s.log.Error("unexpected error", slog.String("error:", err.Error()))
	return e

}
