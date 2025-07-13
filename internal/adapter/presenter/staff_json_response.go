package presenter

type StaffJsonResponse struct {
	ID        uint64 `json:"id" example:"1"`
	Name      string `json:"name" example:"John Doe"`
	Role      string `json:"role" example:"COOK, ATTENDANT or MANAGER"`
	CreatedAt string `json:"created_at" example:"2024-02-09T10:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2024-02-09T10:00:00Z"`
}

type StaffJsonPaginatedResponse struct {
	JsonPagination
	Staffs []StaffJsonResponse `json:"staffs"`
}
