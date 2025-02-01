package db

import (
	"cmp"
	"encoding/json"
	"errors"
	"maps"
	"os"
	"slices"

	"github.com/SalmandaAK/task-cli/task/domain"
)

var (
	errTaskNotFound = errors.New("task not found")
)

type TaskJSONRepository struct {
	filePath string
	tasks    map[int]*domain.Task
}

func NewTaskJSONRepository(filePath string) *TaskJSONRepository {
	return &TaskJSONRepository{
		filePath: filePath,
		tasks:    make(map[int]*domain.Task),
	}
}

func (r *TaskJSONRepository) FindAllTasks() ([]*domain.Task, error) {
	err := r.loadData()
	if err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	if len(r.tasks) > 0 {
		tasks = slices.SortedFunc(maps.Values(r.tasks), func(t1, t2 *domain.Task) int {
			return cmp.Compare(t1.Id, t2.Id)
		})
	}
	return tasks, nil
}

func (r *TaskJSONRepository) CreateTask(t *domain.Task) error {
	if err := r.loadData(); err != nil {
		return err
	}
	t.Id = r.generateId()
	r.tasks[t.Id] = t
	return r.saveData()
}

func (r *TaskJSONRepository) FindTaskById(id int) (*domain.Task, error) {
	if err := r.loadData(); err != nil {
		return nil, err
	}
	t, found := r.tasks[id]
	if !found {
		return nil, errTaskNotFound
	}
	return t, nil
}

func (r *TaskJSONRepository) UpdateTask(updatedTask *domain.Task) error {
	r.tasks[updatedTask.Id] = updatedTask
	return r.saveData()
}

func (r *TaskJSONRepository) DeleteTask(deletedTask *domain.Task) error {
	delete(r.tasks, deletedTask.Id)
	return r.saveData()
}

func (r *TaskJSONRepository) loadData() error {
	jsonBlob, err := os.ReadFile(r.filePath)
	if err != nil {
		if err == err.(*os.PathError) {
			return nil
		} else {
			return err
		}
	}
	if len(jsonBlob) == 0 {
		return nil
	}
	return json.Unmarshal(jsonBlob, &r.tasks)
}

func (r *TaskJSONRepository) saveData() error {
	jsonBlob, err := json.MarshalIndent(r.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, jsonBlob, 0666)
}

// Generate Id generates id by adding 1 into the first unused Id in a list of ids in ascending sort. This function use binary search algorithm to find the first unused Id.
// Generate Id will return 1 if id list is empty or the first Id is not 1.
func (r *TaskJSONRepository) generateId() int {
	idList := slices.Sorted(maps.Keys(r.tasks))
	if len(idList) == 0 || idList[0] != 1 {
		return 1
	}

	// lo and hi are the indices of lowest used Id and highest used Id respectively.
	lo, hi := 0, len(idList)-1

	// If the len(idList) == idList[hi] (highest Id), then there's no unused Id up to the highest used Id). So the next available Id is the highest used Id + 1.
	if len(idList) == idList[hi] {
		return idList[hi] + 1
	}

	// If len(idList) < idList[hi], there's any unused Id between lowest used Id, idList[0] and highest used Id (idList[len(idList) - 1])
	// We are going to find the first unused empty Id by using binary search algorithm until the range between hi and lo is only 1 (hi will always be greater than lo).
	for hi-lo != 1 {
		i := int(uint(lo+hi) >> 1)
		if len(idList[:i+1]) == idList[i] {
			lo = i
		} else {
			hi = i
		}
	}

	// The first unused Id will be positioned after lo, so the next available id will be idList[lo] + 1.
	return idList[lo] + 1
}
