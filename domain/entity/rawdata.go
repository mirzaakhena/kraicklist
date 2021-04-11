package entity

type RawData struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	ThumbURL  string   `json:"thumb_url"`
	Tags      []string `json:"tags"`
	UpdatedAt int64    `json:"updated_at"`
	ImageURLs []string `json:"image_urls"`
}

type DataRequest struct {

}

func NewData(req DataRequest) (*RawData, error) {

	var obj RawData

	return &obj, nil
}

