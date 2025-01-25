package todo

import (
	"gorm.io/gorm"
)
type RepositoryTodo struct{
	Database  gorm.DB

}

func NewTodoRepository(db gorm.DB)*RepositoryTodo{
	return &RepositoryTodo{
		Database: db,
	}
}

func(repo *RepositoryTodo)Create(todo *Todo)(error){
	result := repo.Database.Create(todo)
	if result.Error != nil{
		return nil
	}
	return nil
}