package tasksService

type TasksService interface {
	GetAllTask() ([]Tasks, error)
	CreateTask(tasks Tasks) (Tasks, error)
	GetTaskById(id string) (Tasks, error)
	DeleteTask(id string) error
	ChangeTask(id string, tasks Tasks) (Tasks, error)
	GetTaskByUserId(userId uint) ([]Tasks, error)
}

type TaskRepo struct {
	repo TaskRepository
}

func (s *TaskRepo) GetTaskByUserId(userId uint) ([]Tasks, error) {
	return s.repo.GetTaskByUserId(userId)
}

func (s *TaskRepo) GetAllTask() ([]Tasks, error) {
	return s.repo.GetAllTask()
}

func (s *TaskRepo) CreateTask(tasks Tasks) (Tasks, error) {
	return s.repo.CreateTask(tasks)
}

func (s *TaskRepo) GetTaskById(id string) (Tasks, error) {
	return s.repo.GetTaskById(id)
}

func (s *TaskRepo) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskRepo) ChangeTask(id string, tasks Tasks) (Tasks, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Tasks{}, err
	}
	response := Tasks{
		ID:       task.ID,
		TaskName: tasks.TaskName,
		IsDone:   tasks.IsDone,
		UserId:   task.UserId,
	}
	if err := s.repo.SaveTask(response); err != nil {
		return Tasks{}, err
	}
	return response, nil
}

func NewTaskService(r TaskRepository) TasksService {
	return &TaskRepo{repo: r}
}
