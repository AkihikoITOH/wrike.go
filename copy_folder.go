package wrike

import (
	"fmt"

	"github.com/AkihikoITOH/wrike.go/parameters"
	"github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// CopyFolder copies a folder.
// For details refer to https://developers.wrike.com/documentation/api/methods/copy-folder
func (api *API) CopyFolder(id types.FolderID, params parameters.CopyFolder) (*types.Folders, error) {
	path := fmt.Sprintf("/copy_folder/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.post(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewFoldersFromJSON(data)
}
