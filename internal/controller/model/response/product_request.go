package response

type ProductResponse struct {
	URL string `json:"url" binding:"required"`
	PrecoAlvo float64 `json:"preco_alvo" binding:"required"`
}

type ProductsResponse struct {
    ID         string  `json:"id"`
    URL        string  `json:"url"`
    PrecoAlvo  float64 `json:"preco_alvo"`
    Ativo      bool    `json:"ativo"`
}