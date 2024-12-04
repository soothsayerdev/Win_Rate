package winrate

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"winrate/middlewares"
	"winrate/routes"
	"winrate/utils"
)

func main() {
	// Conection to database
	utils.ConnectDatabase()

	// Initialize router
	router := mux.NewRouter()

    // Apply middlewares
    router.Use(middlewares.CORSMiddleware)

    // Register routes
    routes.RegisterUserRoutes(router)
	routes.RegisterDeckRoutes(router)
	routes.RegisterMatchRoutes(router)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server started on port %s", port)
    err := http.ListenAndServe(":"+port, router)
    if err!= nil {
        log.Fatal(err)
    }
}