package request

type CategoryRequest struct {
	Name string `json:"name" validate:"required, lowercase"`
}
