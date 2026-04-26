package sqlite

import (
	"Proj3_scalable_rest_api/internal/config"
	"Proj3_scalable_rest_api/internal/storage"
	"Proj3_scalable_rest_api/internal/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // '_' --> this is not being used directly in the code, but behind the scene used.
)

var _ storage.Storage = (*Sqlite)(nil)  // compile-time assertion to ensure that Sqlite struct implements the Storage interface

type Sqlite struct{
	Db *sql.DB
}

// constructor
func New(cfg *config.Config) (*Sqlite, error){
	db, err:= sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// creating the students table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT ,
	email TEXT,
	age INTEGER
	)`)

	if err!=nil{
		return nil , err 
	}
	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name, email string, age int) (int64, error){
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)  // execute the db insertion
	if err != nil {return 0, err}
	
	last_id, err:= result.LastInsertId()  // get the id of the newly inserted student record
	if err != nil {return 0, err}

	return last_id, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error){
	// prepare the query to get the student by id
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	// execute the query and scan the result into the student struct
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {

		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}

		return types.Student{}, fmt.Errorf("query error: %w", err)
	}
	return student, nil
}


func (s *Sqlite) GetStudents() ([]types.Student, error){
	stmt, err:= s.Db.Prepare("SELECT id, name, email, age FROM students")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	rows, err:= stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {   // iterator to scan each row
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}