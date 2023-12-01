package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"poyectoWeb/controllers"

	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("se requiere un controlador no nulo para instanciar un manejador")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarUnEmpleado(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un empleado, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.ActualizarUnEmpleado(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un empleado, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarUnEmpleado(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := h.controller.EliminarUnEmpleado(id)
	if err != nil {
		log.Printf("fallo al eliminar un empleado, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al eliminar un empleado, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) LeerUnEmpleado(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	empleado, err := h.controller.LeerUnEmpleado(id)
	if err != nil {
		log.Printf("fallo al leer un empleado, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al leer un empleado, con error: %s", err.Error()), http.StatusNotFound)
		return
	}
	/*
		jsonEmpleado, err := json.Marshal(empleado)
		if err != nil {
			log.Printf("fallo al convertir a JSON, con error: %s", err.Error())
			http.Error(writer, fmt.Sprintf("fallo al convertir a JSON, con error: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	*/
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(empleado)
}

func (h *Handler) LeerEmpleados(writer http.ResponseWriter, req *http.Request) {
	empleados, err := h.controller.LeerEmpleados(100, 0)
	if err != nil {
		log.Printf("fallo al leer empleados, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al leer empleados, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	/*
		jsonEmpleados, err := json.Marshal(empleados)
		if err != nil {
			log.Printf("fallo al convertir a JSON, con error: %s", err.Error())
			http.Error(writer, fmt.Sprintf("fallo al convertir a JSON, con error: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	*/

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(empleados)
}

func (h *Handler) CrearEmpleado(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al leer el cuerpo de la solicitud, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al leer el cuerpo de la solicitud, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	nuevoID, err := h.controller.CrearEmpleado(body)
	if err != nil {
		log.Printf("fallo al crear un empleado, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al crear un empleado, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("ID del nuevo empleado: %d", nuevoID)))
}
