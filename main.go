package main

import (
	"fmt"
	"forum/3rd_party_service/cloudinary"
	"forum/3rd_party_service/mail_sender"
	"forum/3rd_party_service/redis"
	"forum/database"
	"forum/logger"
	"forum/repository"
	"forum/server_http"
)

func main() {
	Initialize()
	defer CleanupUnfinishedTasks()

	logger.GetLogInstance().Info("================================================================")
	logger.GetLogInstance().Info("=====================Application Starting=======================")
	logger.GetLogInstance().Info("================================================================")

	go func() {
		httpServer := server_http.GetHTTPServer()
		logger.GetLogInstance().Info(fmt.Sprintf("HTTP Server running on port: %d", httpServer.Config.Port))
		if err := httpServer.Run(); err != nil {
			logger.GetLogInstance().Error(fmt.Sprintf("Error starting HTTP server: %s", err))
		}
	}()

	logger.GetLogInstance().Info("================================================================")
	logger.GetLogInstance().Info("=====================Application Started========================")
	logger.GetLogInstance().Info("================================================================")

	select {}
}

func Initialize() {
	logger.InitializeLogger()
	cloudinary.InitializeFileUploader()
	mail_sender.InitializeMailSender()
	redis.InitializeRedis()
	database.InitializeDatabaseConnection()
	repository.InitializeRepository(database.GetDatabaseConnection())
	server_http.InitializeHTTPServer()
}

func CleanupUnfinishedTasks() {
	logger.GetLogInstance().Info("================================================================")
	logger.GetLogInstance().Info("=======================Application Stop=========================")
	logger.GetLogInstance().Info("================================================================")
	logger.CleanupQueuedLogs()
}
