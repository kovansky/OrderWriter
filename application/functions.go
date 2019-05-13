package application

func GetWinTitle(title string) string {
	return title + " " + Language.App.TitleSeparator + " " + Language.App.Title
}
