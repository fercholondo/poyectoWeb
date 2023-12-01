package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"poyectoWeb/models"
	repositorio "poyectoWeb/repository"
)

var (
	updateQuery = "UPDATE empleados SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM empleados WHERE id=$1;"
	selectQuery = "SELECT id, identificacion, nombres, apellidos, fecha_ingreso, cargo, es_vinculado, salario FROM empleados WHERE id=$1;"
	listQuery   = "SELECT id, identificacion, nombres, apellidos, fecha_ingreso, cargo, es_vinculado, salario FROM empleados LIMIT $1 OFFSET $2;"
	createQuery = "INSERT INTO empleados (identificacion, nombres, apellidos, fecha_ingreso, cargo, es_vinculado, salario) VALUES (:identificacion, :nombres, :apellidos, :fecha_ingreso, :cargo, :es_vinculado, :salario) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Empleado]
}

func NewController(repo repositorio.Repository[models.Empleado]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("se necesita un repositorio no nulo para instanciar un controlador")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnEmpleado(reqBody []byte, id string) error {
	nuevosValoresEmpleado := make(map[string]interface{}) //revisar esta linea
	err := json.Unmarshal(reqBody, &nuevosValoresEmpleado)
	if err != nil {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un empleado, con error: %s", err.Error())
	}

	if len(nuevosValoresEmpleado) == 0 {
		log.Printf("no se proporcionaron valores para actualizar")
		return fmt.Errorf("no se proporcionaron valores para actualizar")
	}

	query := construirUpdateQuery(nuevosValoresEmpleado)
	nuevosValoresEmpleado["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresEmpleado)
	if err != nil {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un empleado, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]interface{}) string {
	columns := []string{}
	for key := range nuevosValores {
		if key != "id" {
			columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
		}
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) EliminarUnEmpleado(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un empleado, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnEmpleado(id string) ([]byte, error) {
	empleado, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un empleado, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un empleado, con error: %s", err.Error())
	}

	// Conversión de FechaIngreso a un formato de fecha apropiado
	/*
		if emp, ok := empleado.(*models.Empleado); ok {
			if t, err := time.Parse("2006-01-02", emp.FechaIngreso); err == nil {
				emp.FechaIngreso = t.Format("2006-01-02") // El formato deseado para la fecha
			}
		}
	*/
	// Conversión de Salario a un tipo numérico si se almacena como cadena de números
	if empleado.Salario != "" {
		salario, err := strconv.ParseFloat(empleado.Salario, 64)
		if err == nil {
			empleado.Salario = fmt.Sprintf("%.2f", salario) // Ajusta el salario a dos decimales si es necesario
		}
	}

	empleadoJson, err := json.Marshal(empleado)
	if err != nil {
		log.Printf("fallo al leer un empleado, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un empleado, con error: %s", err.Error())
	}
	return empleadoJson, nil
}

func (c *Controller) LeerEmpleados(limit, offset int) ([]byte, error) {
	empleados, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer empleados, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer empleados, con error: %s", err.Error())
	}

	for _, emp := range empleados {
		if emp.Salario != "" {
			salario, err := strconv.ParseFloat(emp.Salario, 64)
			if err == nil {
				emp.Salario = fmt.Sprintf("%.2f", salario) // Ajusta el salario a dos decimales si es necesario
			}
		}
	}

	jsonEmpleados, err := json.Marshal(empleados)
	if err != nil {
		log.Printf("fallo al leer empleados, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer empleados, con error: %s", err.Error())
	}
	return jsonEmpleados, nil
}

func (c *Controller) CrearEmpleado(reqBody []byte) (int64, error) {
	nuevoEmpleado := &models.Empleado{}
	err := json.Unmarshal(reqBody, nuevoEmpleado)
	if err != nil {
		log.Printf("fallo al crear un nuevo empleado, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo empleado, con error: %s", err.Error())
	}

	valoresColumnasNuevoEmpleado := map[string]interface{}{
		"identificacion": nuevoEmpleado.Identificacion,
		"nombres":        nuevoEmpleado.Nombres,
		"apellidos":      nuevoEmpleado.Apellidos,
		"fecha_ingreso":  nuevoEmpleado.FechaIngreso,
		"cargo":          nuevoEmpleado.Cargo,
		"es_vinculado":   nuevoEmpleado.EsVinculado,
		"salario":        nuevoEmpleado.Salario,
	}

	nuevoID, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoEmpleado)
	if err != nil {
		log.Printf("fallo al crear un nuevo empleado, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo empleado, con error: %s", err.Error())
	}
	return nuevoID, nil
}
