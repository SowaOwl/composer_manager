package compose

type UpdateContainerRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

type CreateContainerRequest struct {
	Name        string             `json:"name"`
	Body        string             `json:"body"`
	PublicTypes []createPublicType `json:"public_types"`
}

type createPublicType struct {
	TypeID uint   `json:"type_id"`
	Data   string `json:"data"`
}

type CreateDataTypeRequest struct {
	Name       string `json:"name"`
	ManyName   string `json:"many_name"`
	IsTabulate bool   `json:"is_tabulate"`
}
