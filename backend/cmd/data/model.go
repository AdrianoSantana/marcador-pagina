package data

type Book struct {
	ID             uint64 `json:"id"`
	Title          string `json:"titulo"`
	Last_date_read string `json:"ultima_leitura"`
	Chapter        uint64 `json:"capitulo"`
	Page           uint64 `json:"pagina"`
}
