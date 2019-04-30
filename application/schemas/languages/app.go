package languages

type App struct {
	Title          string `json:"title"`
	TitleSeparator string `json:"title_separator"`
	Writer         Writer `json:"writer"`
}
