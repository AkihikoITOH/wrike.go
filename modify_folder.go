package wrike

import (
	"fmt"
	"strings"

	"github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// ModifyFolder modifiies a folder with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-folder
func (api *API) ModifyFolder(id types.FolderID, params parameters.ModifyFolder) (*types.Folders, error) {
	path := fmt.Sprintf("/folders/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewFoldersFromJSON(data)
}

// ModifyFolders modifiies folders with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-folder
func (api *API) ModifyFolders(ids []types.FolderID, params parameters.ModifyFolders) (*types.Folders, error) {
	s := make([]string, 0, len(ids))
	for _, id := range ids {
		s = append(s, string(id))
	}
	path := fmt.Sprintf("/folders/%s", strings.Join(s, ","))

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewFoldersFromJSON(data)
}
