package main

import ("fmt" 
	"net/http"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"github.com/yhat/scrape"
)

func main() {
	// Загружаем и парсим страницу
	resp, err := http.Get("http://gorod48.ru/")
	if err != nil {	panic(err) }
	root, err := html.Parse(resp.Body)
	if err != nil {	panic(err) }

	// Определяем функцию-сответствия для поиска нужного элемента сайта
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil {
			return scrape.Attr(n.Parent, "class") == "item link"
		}
		return false
	}

	// Ищем все вхождения данного элемента на сайте
	articles := scrape.FindAll(root, matcher)
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}