package params

// type RespTaskError struct {
// 	Id  string `json:"id,omitempty"`
// 	Err string `json:"error,omitempty"`
// }

type AuthPass struct {
	Pass string `json:"password"`
}

type AuthPassError struct {
	MyTocken string `json:"token,omitempty"`
	Err      string `json:"error,omitempty"`
}
