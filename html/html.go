package html

import (
	"github.com/tinyrange/htm/v2"
)

func Html(children ...htm.Fragment) htm.Fragment { return htm.NewHtmlFragment("html", children...) }
func Head(children ...htm.Fragment) htm.Fragment { return htm.NewHtmlFragment("head", children...) }
func Body(children ...htm.Fragment) htm.Fragment { return htm.NewHtmlFragment("body", children...) }
func Div(children ...htm.Fragment) htm.Fragment  { return htm.NewHtmlFragment("div", children...) }

func Title(title string) htm.Fragment { return htm.NewHtmlFragment("title", htm.Text(title)) }

func LinkCSS(url string, children ...htm.Fragment) htm.Fragment {
	childList := []htm.Fragment{
		htm.Attr("rel", "stylesheet"),
		htm.Attr("href", url),
	}
	childList = append(childList, children...)
	return htm.NewHtmlFragment("link", childList...)
}

func JavaScriptSrc(url string, children ...htm.Fragment) htm.Fragment {
	childList := []htm.Fragment{htm.Attr("src", url)}
	childList = append(childList, children...)
	return htm.NewHtmlFragment("script", childList...)
}
