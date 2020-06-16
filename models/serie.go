package models

//Serie Struct to sent info
type Serie struct {
	ID       int    `json:"id_serie"`
	Name     string `json:"name_serie"`
	Episodes int    `json:"episodes_serie"`
	Status   string `json:"status_serie"`
	Type     string `json:"type_serie"`
}
