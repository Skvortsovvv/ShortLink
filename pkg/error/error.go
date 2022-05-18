package error

import "encoding/json"

type Error struct {
	ErrMsg string `json:"error"`
}

func (er *Error) Error() []byte {
	jsonResponse, err := json.Marshal(er)
	if err != nil {
		return []byte{}
	}
	return jsonResponse
}
