package models

type Empleado struct {
	ID             uint64 `db:"id" json:"id"`
	Identificacion int    `db:"identificacion" json:"identificacion"`
	Nombres        string `db:"nombres" json:"nombres"`
	Apellidos      string `db:"apellidos" json:"apellidos"`
	FechaIngreso   string `db:"fecha_ingreso" json:"fecha_ingreso"`
	Cargo          string `db:"cargo" json:"cargo"`
	EsVinculado    bool   `db:"es_vinculado" json:"es_vinculado"`
	Salario        string `db:"salario" json:"salario"`
}
