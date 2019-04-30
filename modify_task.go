package wrike

import (
	"fmt"
	"strings"

	"github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// ModifyTask modifies a task with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-tasks
func (api *API) ModifyTask(id types.TaskID, params parameters.ModifyTask) (*types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}

// ModifyTasks modifies multiple tasks with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-tasks
func (api *API) ModifyTasks(ids []types.TaskID, params parameters.ModifyTasks) (*types.Tasks, error) {
	s := make([]string, 0, len(ids))
	for _, id := range ids {
		s = append(s, string(id))
	}
	path := fmt.Sprintf("/tasks/%s", strings.Join(s, ","))

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}
