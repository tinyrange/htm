package html

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tinyrange/htm/v2"
)

func Text(s string) htm.Fragment                   { return htm.Text(s) }
func Textf(format string, a ...any) htm.Fragment   { return Text(fmt.Sprintf(format, a...)) }
func Html(children ...htm.Fragment) htm.Fragment   { return htm.NewHtmlFragment("html", children...) }
func Head(children ...htm.Fragment) htm.Fragment   { return htm.NewHtmlFragment("head", children...) }
func Body(children ...htm.Fragment) htm.Fragment   { return htm.NewHtmlFragment("body", children...) }
func Div(children ...htm.Fragment) htm.Fragment    { return htm.NewHtmlFragment("div", children...) }
func Span(children ...htm.Fragment) htm.Fragment   { return htm.NewHtmlFragment("span", children...) }
func Button(children ...htm.Fragment) htm.Fragment { return htm.NewHtmlFragment("button", children...) }
func A(children ...htm.Fragment) htm.Fragment      { return htm.NewHtmlFragment("a", children...) }

func Link(target string, children ...htm.Fragment) htm.Fragment {
	childList := []htm.Fragment{
		htm.Attr("href", target),
	}
	childList = append(childList, children...)
	return A(childList...)
}

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

func Title(title string) htm.Fragment { return htm.NewHtmlFragment("title", htm.Text(title)) }

type Id string

func (i Id) Children(ctx context.Context) ([]htm.Fragment, error) {
	return htm.Attr("id", string(i)).Children(ctx)
}
func (i Id) Render(ctx context.Context, parent htm.Node) error {
	return htm.Attr("id", string(i)).Render(ctx, parent)
}

var (
	_ htm.Fragment = Id("")
)

func NewId() Id {
	return Id("i" + strconv.FormatUint(rand.Uint64(), 36))
}

type urlFragment struct{}

// Children implements htm.Fragment.
func (*urlFragment) Children(ctx context.Context) ([]htm.Fragment, error) {
	return []htm.Fragment{}, nil
}

// Render implements htm.Fragment.
func (*urlFragment) Render(ctx context.Context, parent htm.Node) error {
	req, ok := htm.RequestFromContext(ctx)
	if !ok {
		return fmt.Errorf("failed to get request")
	}

	return htm.Text(req.URL.String()).Render(ctx, parent)
}

func Url() htm.Fragment {
	return &urlFragment{}
}
