package main 
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/abh1SHAKE/go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	err := http.ListenAndServe("localhost:1734", r)
	if err != nil {
		log.Error(err)
	}
}