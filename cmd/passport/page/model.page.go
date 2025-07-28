package page

type Model struct {
	Title string
}

type HtmlError struct {
	Title      string
	Path       string
	Header     string
	Info       string
	ShowButton bool
	Button     string
}
