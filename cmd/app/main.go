package main

import (
	"fmt"
	"net/http"

	"myapp/server/router"
	"myapp/config"
	lr "myapp/util/logger"

)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	appRouter := router.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}

// Greet : A function to say hi to people dropping by our server.
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
