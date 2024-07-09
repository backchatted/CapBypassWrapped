package capbypass

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	payload := &CapBypassPayload{
		ClientKey: c.apiKey,
	}

	data, _ := json.Marshal(payload)

	resp, err := http.Post(BASEURL+"/getBalance", "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	response := &CapBypassResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *CapBypass) Solve(task map[string]any) (*CapBypassResponse, error) {
	_ = checkTask(task)
	payload := &CapBypassPayload{
		ClientKey: c.apiKey,
		Task:      &task,
	}

	data, _ := json.Marshal(payload)

	createTaskResp, err := http.Post(BASEURL+"/createTask", "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer createTaskResp.Body.Close()

	createTaskBody, _ := io.ReadAll(createTaskResp.Body)
	createTaskResponse := &CapBypassResponse{}
	err = json.Unmarshal(createTaskBody, createTaskResponse)
	if err != nil {
		return nil, err
	}

	if createTaskResp.StatusCode != 200 {
		return nil, fmt.Errorf(string(createTaskBody))
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

		statusBody, _ := io.ReadAll(statusResp.Body)
		statusResponse := &CapBypassResponse{}
		err = json.Unmarshal(statusBody, statusResponse)
		if err != nil {
			return nil, err
		}

		if statusResponse.Status == "DONE" {
			return statusResponse, nil
		}

		if statusResponse.ErrorId == 1 {
			return nil, errors.New(string(statusBody))
		}
	}
	return nil, errors.New("could not solve")
}
