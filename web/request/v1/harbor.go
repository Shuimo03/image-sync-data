package v1

type Registry struct {
	Page int `json:"page" binding:"required"`
	Size int `json:"size" binding:"required"`
}
