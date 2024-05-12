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

type AuthPass struct {
	Pass string `json:"password"`
}

type AuthPassError struct {
	MyTocken string `json:"token,omitempty"`
	Err      string `json:"error,omitempty"`
}
