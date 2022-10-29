package presenter

type Book struct {
	BookId       string `json:"bookId"`
	Volume       string `json:"volume"`
	DisplayOrder int    `json:"displayOrder"`
	Thumbnail    string `json:"thumbnail"`
	Title        string `json:"title"`
	Filepath     string `json:"filepath"`
	Author       string `json:"author"`
	Publisher    string `json:"publisher"`
	Direction    string `json:"direction"`
}
