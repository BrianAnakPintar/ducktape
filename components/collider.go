package components

import "github.com/yohamta/donburi"

type ColliderData struct {
	HalfWidth, HalfHeight float64 // for center-based positioning
	Static bool // to determine one-time insertion vs updates
}

var Collider = donburi.NewComponentType[ColliderData]()
