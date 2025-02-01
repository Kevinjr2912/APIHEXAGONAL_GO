package infraestructure

import (
	"api-hexagonal/core"
	"api-hexagonal/students/domain/entities"
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

func (mysql *MySQL) CreateStudent(student *entities.Student) (id int64, err error) {
	query := "INSERT INTO students (name, age, phone_number) VALUES (?,?,?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, student.Name, student.Age, student.PhoneNumber)

	if err != nil {
		return 0, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("Error al obtener el id generado al insertar un nuevo estudiante %v", err)
	}

	return id, nil
}