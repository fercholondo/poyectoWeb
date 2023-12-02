package main

import (
	"log"
	"net/http"
	"poyectoWeb/controllers"
	myhandlers "poyectoWeb/handlers"
	"poyectoWeb/models"
	repositorio "poyectoWeb/repository"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq" //
)

func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl)
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	db, err := ConectarDB("postgres://oxvihvzr:x1z5Razkbc9SI7zD5g83BPKg0LNGwVSR@berry.db.elephantsql.com/oxvihvzr", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	repo, err := repositorio.NewRepository[models.Empleado](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := myhandlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}
	router := mux.NewRouter()

	router.Handle("/empleados", http.HandlerFunc(handler.LeerEmpleados)).Methods(http.MethodGet)
	router.Handle("/empleados", http.HandlerFunc(handler.CrearEmpleado)).Methods(http.MethodPost)
	router.Handle("/empleados/{id}", http.HandlerFunc(handler.LeerUnEmpleado)).Methods(http.MethodGet)
	router.Handle("/empleados/{id}", http.HandlerFunc(handler.ActualizarUnEmpleado)).Methods(http.MethodPatch)
	router.Handle("/empleados/{id}", http.HandlerFunc(handler.EliminarUnEmpleado)).Methods(http.MethodDelete)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))

}
