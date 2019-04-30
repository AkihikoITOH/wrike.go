package wrike

import (
	"fmt"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// DeleteFolder deletes an existing folder by given ID.
// For details refer to https://developers.wrike.com/documentation/api/methods/delete-folder
func (api *API) DeleteFolder(id types.FolderID) (*types.Folders, error) {
	path := fmt.Sprintf("/folders/%s", id)

	data, err := api.delete(path)
	if err != nil {
		return nil, err
	}

	return types.NewFoldersFromJSON(data)
}
