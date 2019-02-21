package oops

type `Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (rect *Rect) CalArea() float64 {
	return rect.Width * rect.Height
}
