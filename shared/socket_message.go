package shared

type SocketMessage struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
