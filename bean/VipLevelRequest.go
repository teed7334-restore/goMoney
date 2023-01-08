package bean

type VipLevelRequest struct {
	Params params `url:"params" json:"params"`
	Path   string `url:"path" json:"path"`
}

type params struct {
	Nonce int64 `url:"nonce" json:"nonce"`
}
