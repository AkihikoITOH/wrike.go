package wrike

import (
	"fmt"

	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// GetFolderTree fetches a list of entries required to build a folder tree for the current account.
// For details refer to https://developers.wrike.com/documentation/api/methods/get-folder-tree
func (api *API) GetFolderTree(params *parameters.GetFolderTree) (*types.FolderTree, error) {
	path := "/folders"

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewFolderTreeFromJSON(data)
}

// GetFolderSubtree fetches a list of tree entries for subtree of the folder.
// For details refer to https://developers.wrike.com/documentation/api/methods/get-folder-tree
func (api *API) GetFolderSubtree(id types.FolderID, params *parameters.GetFolderSubtree) (*types.FolderTree, error) {
	path := fmt.Sprintf("/folders/%s/folders", id)

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewFolderTreeFromJSON(data)
}
