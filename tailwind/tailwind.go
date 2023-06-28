package tailwind

import "github.com/tinyrange/htm/v2"

type fontWeight struct {
	Thin       htm.Fragment // font-thin	font-weight: 100;
	ExtraLight htm.Fragment // font-extralight	font-weight: 200;
	Light      htm.Fragment // font-light	font-weight: 300;
	Normal     htm.Fragment // font-normal	font-weight: 400;
	Medium     htm.Fragment // font-medium	font-weight: 500;
	SemiBold   htm.Fragment // font-semibold	font-weight: 600;
	Bold       htm.Fragment // font-bold	font-weight: 700;
	ExtraBold  htm.Fragment // font-extrabold	font-weight: 800;
	Black      htm.Fragment // font-black	font-weight: 900;
}

type font struct {
	Weight fontWeight
}

const (
	Container = htm.Class("container")
)

var (
	Font = font{Weight: fontWeight{
		Thin:       htm.Class("font-thin"),
		ExtraLight: htm.Class("font-extralight"),
		Light:      htm.Class("font-light"),
		Normal:     htm.Class("font-normal"),
		Medium:     htm.Class("font-medium"),
		SemiBold:   htm.Class("font-semibold"),
		Bold:       htm.Class("font-bold"),
		ExtraBold:  htm.Class("font-extrabold"),
		Black:      htm.Class("font-black"),
	}}
)
