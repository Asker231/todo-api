package todo

import (
	"fmt"

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

func(repo *RepositoryTodo)DeleteTodo(id int)(error){
	var todo Todo
	result := repo.Database.Delete(&todo,"id = ?",id)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
		return nil
	}
	return nil
}

func(repo *RepositoryTodo)UpdateTodo(id int)(*Todo,error){
	var todo Todo
	result := repo.Database.First(&todo,"id = ?",id)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
		return nil,result.Error
	}	
	result.Model(&todo).Update("complited",true)
	return &todo,nil
}

func(repo *RepositoryTodo)GetAll()(*[]Todo,error){
	var todos []Todo
	result := repo.Database.Find(&todos)
	if result.Error != nil{
		return nil,result.Error
	}
	return &todos,nil
}

func(repo *RepositoryTodo)GetTodoById(id int)(*Todo,error){
	var todo Todo
	result := repo.Database.First(&todo,"id = ?",id)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
		return nil,result.Error
	}
	return &todo,nil	
}