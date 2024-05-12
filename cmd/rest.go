package restserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rarrazaan/be-player-performance-app/internal/dependency"
)

func Serve(dep *dependency.Dependency) {
	ddep := dependency.NewDirectDependency(dep)
	router := dependency.NewRouter(*dep, *ddep)
	router.ContextWithFallback = true

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")),
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
