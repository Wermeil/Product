package http

import (
	"Ctrl/internal/models"
	userService2 "Ctrl/internal/services"
	"context"
	"fmt"
)

type CombinedHandler struct {
	*UserHandlerService
	*TaskHandlerService
}

type TaskHandlerService struct {
	taskService userService2.TasksService
}
type UserHandlerService struct {
	userService userService2.UserService
}

func NewUserHandler(u userService2.UserService) *UserHandlerService {
	return &UserHandlerService{userService: u}
}

func NewTaskHandler(t userService2.TasksService) *TaskHandlerService {
	return &TaskHandlerService{taskService: t}
}

func (t *TaskHandlerService) GetTask(ctx context.Context, request GetTaskRequestObject) (GetTaskResponseObject, error) {
	tasks, err := t.taskService.GetAllTask()
	if err != nil {
		return nil, err
	}
	response := GetTask200JSONResponse{}
	for _, task := range tasks {
		obj := Tasks{
			Id:       &task.ID,
			IsDone:   &task.IsDone,
			TaskName: &task.TaskName,
			UserId:   &task.UserId,
		}
		response = append(response, obj)
	}
	return response, nil
}

func (t *TaskHandlerService) PostTask(ctx context.Context, request PostTaskRequestObject) (PostTaskResponseObject, error) {
	body := models.Tasks{
		IsDone:   *request.Body.IsDone,
		TaskName: *request.Body.TaskName,
		UserId:   uint(*request.Body.UserId),
	}
	obj, err := t.taskService.CreateTask(body)
	if err != nil {
		return nil, err
	}
	var response = PostTask201JSONResponse{
		Id:       &obj.ID,
		IsDone:   &obj.IsDone,
		TaskName: &obj.TaskName,
		UserId:   &obj.UserId,
	}
	return response, nil

}

func (t *TaskHandlerService) DeleteTasksId(ctx context.Context, request DeleteTasksIdRequestObject) (DeleteTasksIdResponseObject, error) {
	id := fmt.Sprintf("%v", request.Id)
	if err := t.taskService.DeleteTask(id); err != nil {
		return nil, err
	}
	return DeleteTasksId204Response{}, nil
}

func (t *TaskHandlerService) PatchTasksId(ctx context.Context, request PatchTasksIdRequestObject) (PatchTasksIdResponseObject, error) {
	idString := fmt.Sprintf("%v", request.Id)
	id := uint(request.Id)
	userId := uint(*request.Body.UserId)
	task := models.Tasks{
		ID:       id,
		TaskName: *request.Body.TaskName,
		IsDone:   *request.Body.IsDone,
		UserId:   userId,
	}
	obj, err := t.taskService.ChangeTask(idString, task)
	if err != nil {
		return nil, err
	}
	response := PatchTasksId200JSONResponse{
		Id:       &id,
		IsDone:   &obj.IsDone,
		TaskName: &obj.TaskName,
		UserId:   &obj.UserId,
	}
	return response, nil
}

func (u *UserHandlerService) GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error) {
	arr, err := u.userService.GetUser()
	if err != nil {
		return nil, err
	}
	response := GetUser200JSONResponse{}

	for _, ass := range arr {
		TasksInDb, _ := u.userService.GetTasksForUser(ass.ID)
		var tasks []Tasks
		for _, task := range TasksInDb {
			obj := Tasks{
				Id:       &task.ID,
				IsDone:   &task.IsDone,
				TaskName: &task.TaskName,
				UserId:   &task.UserId,
			}
			tasks = append(tasks, obj)
		}
		task := Users{
			Email:    &ass.Email,
			Id:       &ass.ID,
			Password: &ass.Password,
			Tasks:    &tasks,
		}
		response = append(response, task)
	}
	return response, nil
}

func (u *UserHandlerService) PostUser(ctx context.Context, request PostUserRequestObject) (PostUserResponseObject, error) {
	body := models.Users{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}
	obj, err := u.userService.CreateUser(body)
	if err != nil {
		return nil, err
	}
	response := PostUser201JSONResponse{
		Email:    &obj.Email,
		Id:       &obj.ID,
		Password: &obj.Password,
	}
	return response, nil

}

func (u *UserHandlerService) DeleteUserId(ctx context.Context, request DeleteUserIdRequestObject) (DeleteUserIdResponseObject, error) {
	id := fmt.Sprintf("%v", request.Id)
	err := u.userService.DeleteUserById(id)
	if err != nil {
		return nil, err
	}
	return DeleteUserId204Response{}, nil
}

func (u *UserHandlerService) PatchUserId(ctx context.Context, request PatchUserIdRequestObject) (PatchUserIdResponseObject, error) {
	id := fmt.Sprintf("%v", request.Id)
	userID := uint(request.Id)
	user := models.Users{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}
	if err := u.userService.ChangeUserById(id, user); err != nil {
		return nil, err
	}
	response := PatchUserId200JSONResponse{
		Email:    request.Body.Email,
		Id:       &userID,
		Password: request.Body.Password,
	}
	return response, nil
}
