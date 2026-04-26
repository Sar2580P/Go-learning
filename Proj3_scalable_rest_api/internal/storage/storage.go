package storage

import "Proj3_scalable_rest_api/internal/types"

type Storage interface{
	CreateStudent(name, email string, age int) (int64, error)

	GetStudentById(id int64) (types.Student, error)

}