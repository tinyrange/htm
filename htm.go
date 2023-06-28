package htm

import (
	"context"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Node interface {
	AddChild(node Node) error
	AddAttribute(key string, value string) error
}

type htmlNode struct {
	node *html.Node
}

// AddAttribute implements Node.
func (n *htmlNode) AddAttribute(key string, value string) error {
	if n.node.Type == html.TextNode {
		return fmt.Errorf("cannot add attributes to text nodes")
	}

	n.node.Attr = append(n.node.Attr, html.Attribute{Key: key, Val: value})

	return nil
}

// AddChild implements Node.
func (n *htmlNode) AddChild(node Node) error {
	if n.node.Type == html.TextNode {
		return fmt.Errorf("cannot add attributes to text nodes")
	}

	newChild := node.(*htmlNode).node
	n.node.AppendChild(newChild)

	return nil
}

var (
	_ Node = &htmlNode{}
)

func newHtmlNode(tag string) Node {
	return &htmlNode{node: &html.Node{Type: html.ElementNode, Data: tag}}
}

type Fragment interface {
	Children(ctx context.Context) ([]Fragment, error)
	Render(ctx context.Context, parent Node) error
}

type htmlFragment struct {
	tag      string
	children []Fragment
}

// Children implements Fragment.
func (h *htmlFragment) Children(ctx context.Context) ([]Fragment, error) {
	return h.children, nil
}

// Render implements Fragment.
func (h *htmlFragment) Render(ctx context.Context, parent Node) error {
	newNode := newHtmlNode(h.tag)

	for _, child := range h.children {
		err := child.Render(ctx, newNode)
		if err != nil {
			return err
		}
	}

	return parent.AddChild(newNode)
}

var (
	_ Fragment = &htmlFragment{}
)

func NewHtmlFragment(tag string, children ...Fragment) Fragment {
	return &htmlFragment{tag: tag, children: children}
}

type Text string

func (Text) Children(ctx context.Context) ([]Fragment, error) {
	return []Fragment{}, nil
}

// Render implements Fragment.
func (t Text) Render(ctx context.Context, parent Node) error {
	return parent.AddChild(&htmlNode{node: &html.Node{Type: html.TextNode, Data: string(t)}})
}

var (
	_ Fragment = Text("")
)

type attr struct {
	key   string
	value string
}

// Children implements Fragment.
func (*attr) Children(ctx context.Context) ([]Fragment, error) {
	return []Fragment{}, nil
}

// Render implements Fragment.
func (a *attr) Render(ctx context.Context, parent Node) error {
	return parent.AddAttribute(a.key, a.value)
}

var (
	_ Fragment = &attr{}
)

func Attr(key string, value string) Fragment {
	return &attr{key: key, value: value}
}

type topLevel struct {
	top Node
}

// AddAttribute implements Node.
func (*topLevel) AddAttribute(key string, value string) error {
	return fmt.Errorf("cannot add attributes to a top level node")
}

// AddChild implements Node.
func (t *topLevel) AddChild(node Node) error {
	if t.top == nil {
		t.top = node
		return nil
	} else {
		return fmt.Errorf("cannot have multiple top level nodes")
	}
}

var (
	_ Node = &topLevel{}
)

func Render(ctx context.Context, w io.Writer, frag Fragment) error {
	top := &topLevel{top: nil}

	err := frag.Render(ctx, top)
	if err != nil {
		return err
	}

	node := top.top.(*htmlNode).node

	err = html.Render(w, node)
	if err != nil {
		return err
	}

	return nil
}
