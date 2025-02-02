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

func (mysql *MySQL) GetAllCareers() (careersArray *[]entities.Career, err error) {
	var careers []entities.Career
	var career entities.Career 

	query := "SELECT * FROM careers"

	rows := mysql.conn.FetchRows(query)

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&career.Id, &career.Name, &career.Duration, &career.Type); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %w", err)
		}

		careers = append(careers, career)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error después de iterar sobre las filas: %w", err)
	}

	return &careers, nil
	
}

func (mysql *MySQL) UpdateCareer(id int64, career *entities.Career) (err error) {
	query := "UPDATE careers SET name = ?, duration = ?, type = ? WHERE id_career = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, career.Name, career.Duration, career.Type, id)

	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("Error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
        return fmt.Errorf("No se encontró una carrea con el Id proporcionado para actualizar su información")
    }

	return nil	
}