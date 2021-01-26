package dtos

type Dummy struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	CompanyID int64  `json:"companyId"`
}

type DummyCreateRequest struct {
	Title     string       `json:"title" binding:"required"`
	CompanyID int64        `json:"companyId" binding:"required"`
	Metadata  *interface{} `json:"metadata,omitempty"`
}

type DummyPatchRequest struct {
	Title string `form:"title," json:"title"`
}
