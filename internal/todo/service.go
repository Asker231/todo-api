package todo

type TodoService struct{
	todorepo *RepositoryTodo
}

func NewServiceTodo(repo *RepositoryTodo)*TodoService{
	return &TodoService{
		todorepo: repo,
	}
}

//Create method
func(service *TodoService)ServiceCreateTodo(todo *Todo)(string,error){
	err := service.todorepo.Create(todo)
	if err != nil{
		return err.Error(),nil
	}
	return "Create todo",nil
}

//Get All method
func(service *TodoService)ServiceGetAllTodos()(*[]Todo,error){
	todos,err := service.todorepo.GetAll()
	if err != nil{
		return nil,err
	}
	return todos,nil
}

//Get By Id method
func(service *TodoService)ServiceGetByIdTodo(id int)(*Todo,error){
	todo,err := service.todorepo.GetTodoById(id)
	if err != nil{
		return nil,err
	}
	return todo,nil
}

//Update Todo method
func(service *TodoService)ServiceUpdateTodo(todo Todo)(error){
	_,err := service.todorepo.UpdateTodo(&todo)
	if err != nil{
		return err
	}
	return err
}

//Delete Todo method
func(service *TodoService)ServiceDeleteTodo(id int)(error){
	 err :=  service.todorepo.DeleteTodo(id)
	 if  err != nil{
		return err
	}
	return err

}