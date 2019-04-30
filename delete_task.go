package wrike

import (
	"fmt"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// DeleteTask deletes an existing task by given ID.
// For details refer to https://developers.wrike.com/documentation/api/methods/delete-tasks
func (api *API) DeleteTask(id types.TaskID) (*types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", id)

	data, err := api.delete(path)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}
