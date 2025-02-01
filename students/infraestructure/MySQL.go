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

func (mysql *MySQL) GetAllStudents() (studentsArray *[]entities.Student, err error){
	var students []entities.Student
	var student entities.Student

	query := "SELECT * FROM students"

	rows := mysql.conn.FetchRows(query)

	defer rows.Close()

	for rows.Next() {
		
		if err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.PhoneNumber); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %w", err)
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error después de iterar sobre las filas: %w", err)
	}

	return &students, nil

}

func (mysql *MySQL) UpdateStudent(id int64, student *entities.Student) (err error) {
	query := "UPDATE students SET name = ?,  age = ?, phone_number = ? WHERE id_student = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, student.Name, student.Age, student.PhoneNumber, id)

	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("Error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
        return fmt.Errorf("No se encontró un estudiante con el ID proporcionado para actualizar su información")
    }

	return nil

}


func (mysql *MySQL) DeleteStudent(id int64) (err error) {
	query := "DELETE FROM students WHERE id_student = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)

    if err != nil {
        return fmt.Errorf("Error al ejecutar la consulta: %v", err)
    }

    rowsAffected, err := result.RowsAffected()

    if err != nil {
		return fmt.Errorf("Error al obtner las filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
        return fmt.Errorf("No se encontró un estudiante con el ID proporcionado para eliminar")
    }

    return nil
}