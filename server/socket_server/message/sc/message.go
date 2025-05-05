package sc

type SCLogin struct {
	Status int `json:"status"`
}

const LoginSuccess = 0
const LoginFailure = 1
