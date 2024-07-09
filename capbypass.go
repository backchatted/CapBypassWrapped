package capbypass

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BASEURL      = "https://capbypass.com/api"
	TASK_TIMEOUT = 250
)

func New(apiKey string) *CapBypass {
	return &CapBypass{apiKey: apiKey}
}

func (c *CapBypass) Balance() (*CapBypassResponse, error) {
	jsonPayload := &CapBypassPayload{
		ClientKey: c.apiKey,
	}

	b, _ := json.Marshal(jsonPayload)

	resp, err := http.Post(BASEURL+"/getBalance", "application/json", bytes.NewReader(b))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response CapBypassResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *CapBypass) Solve(task CapBypassPayload) (*CapBypassResponse, error) {
	_ = checkTask(task)
	payload := &CapBypassPayload{
		ClientKey: c.apiKey,
		Task:      task.Task,
	}

	data, _ := json.Marshal(payload)

	createTaskResp, err := http.Post(BASEURL+"/createTask", "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer createTaskResp.Body.Close()

	var createTaskResponse CapBypassResponse
	if err := json.NewDecoder(createTaskResp.Body).Decode(&createTaskResponse); err != nil {
		return nil, err
	}

	if createTaskResp.StatusCode != 200 {
		return nil, fmt.Errorf(createTaskResponse.Status)
	}

	for i := 0; i < TASK_TIMEOUT; i++ {
		statusPayload := &CapBypassPayload{
			ClientKey: c.apiKey,
			TaskId:    createTaskResponse.TaskId,
		}
		statusData, _ := json.Marshal(statusPayload)

		statusResp, err := http.Post(BASEURL+"/getTaskResult", "application/json", bytes.NewReader(statusData))
		time.Sleep(time.Second * 1)

		if err != nil {
			return nil, err
		}
		defer statusResp.Body.Close()

		var statusResponse CapBypassResponse
		if err := json.NewDecoder(statusResp.Body).Decode(&statusResponse); err != nil {
			return nil, err
		}

		if statusResponse.Status == "DONE" {
			return &statusResponse, nil
		}

		if statusResponse.ErrorId == 1 {
			return nil, errors.New(statusResponse.Status)
		}
	}
	return nil, errors.New("could not solve")
}
