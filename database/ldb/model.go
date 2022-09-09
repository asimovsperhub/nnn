package ldb

type S3 struct {
	Upload map[string]string `json:"upload"`
}

type Ipfs struct {
	Upload map[string]string `json:"upload"`
}

type Card struct {
	Name map[string]string `json:"name"`
}
type NotWhite struct {
	Name map[string]string `json:"name"`
}

type TokenIdName struct {
	TokenName map[string]string `json:"token_name"`
}
