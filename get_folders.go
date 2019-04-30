package wrike

import (
	"fmt"
	"strings"

	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// GetFolders fetches complete information about specified folders.
// For details refer to https://developers.wrike.com/documentation/api/methods/get-folder
func (api *API) GetFolders(ids []types.FolderID, params *parameters.GetFolders) (*types.Folders, error) {
	s := make([]string, 0, len(ids))
	for _, id := range ids {
		s = append(s, string(id))
	}
	path := fmt.Sprintf("/folders/%s", strings.Join(s, ","))

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewFoldersFromJSON(data)
}
