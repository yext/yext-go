package yext

const folderPath = "folders"

type FolderService struct {
	client *Client
}

type folderListResponse struct {
	Folders []*Folder `json:"folders"`
}

func (s *FolderService) List() ([]*Folder, error) {
	v := &folderListResponse{}
	err := s.client.DoRequest("GET", folderPath, v)
	return v.Folders, err
}
