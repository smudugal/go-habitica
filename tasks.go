package habitica

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TaskResponse struct {
	Success bool `json:"success"`
	Data    Task `json:"data"`
}

type TasksResponse struct {
	Success bool   `json:"success"`
	Data    []Task `json:"data"`
}

type TaskService struct {
	client *HabiticaClient
}

func newTaskService(h *HabiticaClient) *TaskService {
	return &TaskService{
		client: h,
	}
}

func (t *TaskService) Get(ctx context.Context, id string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(
		http.MethodGet,
		fmt.Sprintf("tasks/%s", id),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	var taskResp TaskResponse
	err = json.NewDecoder(resp.Body).Decode(&taskResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return &taskResp, err
}

func (t *TaskService) List(ctx context.Context) (*TasksResponse, error) {
	req, err := t.client.NewRequest(http.MethodGet, "tasks/user")
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	var tasksResp TasksResponse
	err = json.NewDecoder(resp.Body).Decode(&tasksResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return &tasksResp, err
}

func (t *TaskService) Update(ctx context.Context, id string, task *Task) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPut, fmt.Sprintf("tasks/%s", id))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()
	return nil, nil
}
