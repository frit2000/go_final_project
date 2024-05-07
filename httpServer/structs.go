package httpServer

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat,omitempty"`
}

type RespTaskAdd struct {
	Id  int64 `json:"id,omitempty"`
	Err error `json:"error,omitempty"`
}
