package dto_article

type CreateArticleRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"desc" form:"desc" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
	PostedAt    string `json:"posted" form:"posted" validate:"required"`
	Creator     string `json:"creator" form:"creator" validate:"required"`
}

type UpdateArticleRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"desc" form:"desc"`
	Image       string `json:"image" form:"image"`
	PostedAt    string `json:"posted" form:"posted"`
	Creator     string `json:"creator" form:"creator"`
}
