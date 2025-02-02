package infraestructure

import (
	"api-hexagonal/career/domain/entities"
	"api-hexagonal/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateCareer(career *entities.Career) (id int64, err error){
	query := "INSERT INTO careers (name, duration, type) VALUES (?,?,?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, career.Name, career.Duration, career.Type)

	if err != nil {
		return 0, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("Error al obtener el id generado al insertar una nueva carrera: %v", err)
	}

	return id, nil
}