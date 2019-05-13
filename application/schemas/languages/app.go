package languages

type App struct {
	Title          string `json:"title"`
	TitleSeparator string `json:"title_separator"`
	Writer         Writer `json:"writer"`
	Next           string `json:"next"`
	Error          string `json:"error"`
}
