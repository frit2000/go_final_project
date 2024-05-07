package httpServer

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat,omitempty"`
}

type RespTaskError struct {
	Id  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}
