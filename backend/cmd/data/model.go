package data

type Book struct {
	id             uint64 `json:"id"`
	title          string `json:"titulo"`
	last_date_read string `json:"ultima_leitura"`
	chapter        uint64 `json:"capitulo"`
	page           uint64 `json:"pagina"`
}
