package wrike

import (
	"fmt"
	"strings"

	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// QueryTasks fetches a list of tasks.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-tasks
func (api *API) QueryTasks(params parameters.QueryTasks) (*types.Tasks, error) {
	path := "/tasks"

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}

// QueryTasksInFolder fetches a list of tasks.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-tasks
func (api *API) QueryTasksInFolder(folderID types.FolderID, params parameters.QueryTasks) (*types.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", folderID)

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}

// QueryTasksByIDs fetches a list of tasks.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-tasks
func (api *API) QueryTasksByIDs(ids []types.TaskID, params parameters.QueryTasksByIDs) (*types.Tasks, error) {
	s := make([]string, 0, len(ids))
	for _, id := range ids {
		s = append(s, string(id))
	}
	path := fmt.Sprintf("/tasks/%s", strings.Join(s, ","))

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}
