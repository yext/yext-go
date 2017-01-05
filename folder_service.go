package yext

const folderPath = "folders"

var (
	FolderListMaxLimit = 1000
)

// Folder is a representation of a Folder in Yext Location Manager.
// For details see http://developer.yext.com/docs/api-reference/#operation/getLocationFolders
type Folder struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}

type FolderService struct {
	client *Client
}

type FolderListResponse struct {
	Count   int       `json:"count"`
	Folders []*Folder `json:"folders"`
}

func (s *FolderService) ListAll() ([]*Folder, error) {
	var folders []*Folder
	var lg listRetriever = func(opts *ListOptions) (int, int, error) {
		flr, _, err := s.List(opts)
		if err != nil {
			return 0, 0, err
		}
		folders = append(folders, flr.Folders...)
		return len(flr.Folders), flr.Count, err
	}

	if err := listHelper(lg, FolderListMaxLimit); err != nil {
		return nil, err
	} else {
		return folders, nil
	}
}

func (s *FolderService) List(opts *ListOptions) (*FolderListResponse, *Response, error) {
	requrl, err := addListOptions(folderPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &FolderListResponse{}
	r, err := s.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	return v, r, err
}
