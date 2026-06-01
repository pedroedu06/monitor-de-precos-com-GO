package request

type ProductRequest struct {
	URL string `json:"url" binding:"required"`
	PrecoAlvo float64 `json:"preco_alvo" binding:"required"`
}

