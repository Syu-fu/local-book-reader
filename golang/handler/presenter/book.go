package presenter

type Book struct {
	BookId       string `json:"bookId"`
	Volume       string `json:"volume"`
	DisplayOrder int    `json:"displayOrder"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Publisher    string `json:"publisher"`
	Direction    string `json:"direction"`
}
