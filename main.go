package main

import (
	"fmt"
	"forum/3rd_party_service/cloudinary"
	"forum/3rd_party_service/mail_sender"
	"forum/3rd_party_service/redis"
	"forum/database"
	"forum/handler"
	"forum/logger"
	"forum/repository"
	"forum/server/http"
	"forum/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Initialize()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	logger.GetLogInstance().Info("=====================Application Starting=======================")

	go func() {
		httpServer := http.GetHTTPServer()
		logger.GetLogInstance().Info(fmt.Sprintf("HTTP Server running on port: %d", httpServer.Config.Port))
		if err := httpServer.Run(); err != nil {
			logger.GetLogInstance().Error(fmt.Sprintf("Error starting HTTP server: %s", err))
		}
	}()

	logger.GetLogInstance().Info("=====================Application Started========================")

	sig := <-sigChan
	logger.GetLogInstance().Info(fmt.Sprintf("Received signal: %v. Initiating shutdown...", sig))

	CleanupUnfinishedTasks()

	os.Exit(0)
}

func Initialize() {
	logger.InitializeLogger()
	cloudinary.InitializeFileUploader()
	mail_sender.InitializeMailSender()
	redis.InitializeRedis()
	database.InitializeDatabaseConnection()
	repository.InitializeRepository(database.GetDatabaseConnection())
	http.InitializeHTTPServer()
	handler.InitializeHandler(http.GetHTTPServer().RouterGroup)
}

func CleanupUnfinishedTasks() {
	logger.GetLogInstance().Info("==================== Application Stopping ======================")

	utils.ShutdownPool()
	logger.CleanupQueuedLogs()

	logger.GetLogInstance().Info("===================== Application Stopped ======================")
}
