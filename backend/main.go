package main

import (
	"backend/src/api/handler"
	dbRepo "backend/src/repository"
	"backend/src/usecase/phoneNumber"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var r *gin.Engine

func init() {

	logger := getLogger()

	db, err := sqlx.Connect("sqlite3", "./sample.db")

	if err != nil {
		logger.Error().Msgf("Error in connection to sqlite: %s", err)
	}
	db.MustExec(dbRepo.Schema)

	repo := dbRepo.NewSqlLiteDB(db)
	srv := phoneNumber.LoadService(repo, &logger)
	r = handler.NewGinHandler(srv)

}

func main() {
	r.Run(":8080")
}

func getLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	return log.With().Logger()
}
