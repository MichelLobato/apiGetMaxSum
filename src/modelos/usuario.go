package modelos

// Struct que recebe o Json presente no Body
type Resp struct {
	Lista []int `json:"list"`
}

// Struct para Resposta do JSON
type RespJSON struct {
	Sum       int `json:"\n"`
	Positions []int
}
