package entities

// Book
type Book struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Authors       []string       `json:"authors"`
	ImageLinks    BookImageLinks `json:"image_links"`
	PrintType     string         `json:"print_type"`
	Language      string         `json:"language"`
	PublishedDate *int32         `json:"published_date"`
	PageCount     *int32         `json:"page_count"`
	Description   *string        `json:"description"`
}

// BookImageLinks
type BookImageLinks struct {
	SmallThumbnail string `json:"small_thumbnail"`
	Thumbnail      string `json:"thumbnail"`
}
