package main

import (
	"errors"
	"fmt"
	"forum/3rd_party_service/cloudinary_service"
	"forum/3rd_party_service/mail_sender"
	"forum/3rd_party_service/redis_service"
	"forum/database"
	"forum/logger"
	"forum/repository"
	"forum/server/http_server"
	"forum/server/http_server/handler"
	"forum/server/socket_server"
	"forum/server/socket_server/event"
	"forum/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Initialize()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	logger.GetLogInstance().Info("=====================Application Starting=======================")

	// Start HTTP server
	go func() {
		httpServer := http_server.GetHTTPServer()
		logger.GetLogInstance().Info(fmt.Sprintf("HTTP Server running on port: %d", httpServer.Config.Port))
		err := httpServer.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetLogInstance().Error(fmt.Sprintf("Error starting HTTP server: %s", err))
		}
	}()

	// Start Socket server
	go func() {
		socketServer := socket_server.GetSocketServer()
		logger.GetLogInstance().Info(fmt.Sprintf("Socket server running on port: %d", socketServer.Config.Port))
		err := socketServer.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetLogInstance().Error(fmt.Sprintf("Error starting Socket server: %s", err))
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
	cloudinary_service.InitializeFileUploader()
	mail_sender.InitializeMailSender()
	redis_service.InitializeRedis()
	database.InitializeDatabaseConnection()
	repository.InitializeRepository(database.GetDatabaseConnection())
	http_server.InitializeHTTPServer()
	handler.InitializeHandler(http_server.GetHTTPServer().RouterGroup)
	socket_server.InitializeSocketServer()
	event.RegisterEvent(socket_server.GetSocketServer().Router)
}

func CleanupUnfinishedTasks() {
	logger.GetLogInstance().Info("==================== Application Stopping ======================")

	if httpServer := http_server.GetHTTPServer(); httpServer != nil {
		if err := httpServer.Close(); err != nil {
			logger.GetLogInstance().Error(fmt.Sprintf("Error closing HTTP server: %s", err))
		}
	}
	if socketServer := socket_server.GetSocketServer(); socketServer != nil {
		if err := socketServer.Close(); err != nil {
			logger.GetLogInstance().Error(fmt.Sprintf("Error closing Socket server: %s", err))
		}
	}

	utils.ShutdownPool()
	logger.GetLogInstance().Info("===================== Application Stopped ======================")

	logger.CleanupQueuedLogs()
}
