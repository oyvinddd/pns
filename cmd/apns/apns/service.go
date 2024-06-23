package apns

import "github.com/jackc/pgx"

type Service struct {
	dbConn *pgx.Conn
}

func NewService(dbConn *pgx.Conn) *Service {
	return &Service{dbConn: dbConn}
}
