package resp

type Address struct {
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Total int           `json:"total"`
	List  []interface{} `json:"list"`
	Final int64         `json:"final"`
}
