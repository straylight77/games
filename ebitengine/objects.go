package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//func init() {
//	s1 := rand.NewSource(time.Now().UnixNano())
//	r1 := rand.New(s1)
//}

/**************************************************************************
 *                              GameObject                                *
 **************************************************************************/
type GameObject struct {
	X, Y          float64
	Angle         float64
	VelX, VelY    float64
	Img           *ebiten.Image
	Width, Height float64
}

// ------------------------------------------------------------------------
func (obj *GameObject) Update(maxX, maxY float64) {
	obj.X += obj.VelX
	obj.Y += obj.VelY
	obj.checkBoundary(maxX, maxY)
}

// ------------------------------------------------------------------------
func (obj *GameObject) Speed() float64 {
	return math.Sqrt(obj.VelX*obj.VelX + obj.VelY*obj.VelY)
}

// ------------------------------------------------------------------------
func (obj *GameObject) Direction() float64 {
	return math.Atan(obj.VelY / obj.VelX)
}

// ------------------------------------------------------------------------
func (obj *GameObject) checkBoundary(maxX, maxY float64) {
	hpad := obj.Width / 2
	vpad := obj.Height / 2

	if obj.X+hpad < 0 { // left side
		obj.X = maxX + hpad
	} else if obj.X-hpad > maxX { // right side
		obj.X = -hpad
	}

	if obj.Y+vpad < 0 { // top
		obj.Y = maxY + vpad
	} else if obj.Y-vpad > maxY { // bottom
		obj.Y = -vpad
	}
}

// ------------------------------------------------------------------------
func (obj *GameObject) LoadSprite(fname string) {
	var err error
	obj.Img, _, err = ebitenutil.NewImageFromFile(fname)
	if err != nil {
		panic(err)
	}
	size := obj.Img.Bounds().Size()
	obj.Width = float64(size.X)
	obj.Height = float64(size.Y)
}

// ------------------------------------------------------------------------
func (obj *GameObject) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-obj.Width/2, -obj.Height/2)
	op.GeoM.Rotate(obj.Angle * 2 * math.Pi / 360)
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(obj.Img, op)
}

/**************************************************************************
 *                               Asteroid                                 *
 **************************************************************************/
type Asteroid struct {
	GameObject
	stage int
}

// ------------------------------------------------------------------------
func getRandDirection() float64 {
	return rand.Float64() * 360
}

// ------------------------------------------------------------------------
func getRandSpeed() float64 {
	return rand.Float64()*(5-2) + 2
}

// ------------------------------------------------------------------------
func MakeRandomAsteroid() Asteroid {
	return MakeAsteroid(250, 250, 45, 3)
}

// ------------------------------------------------------------------------
func MakeAsteroid(x, y, direction, speed float64) Asteroid {
	vx := math.Cos(direction*math.Pi/180) * speed
	vy := math.Sin(direction*math.Pi/180) * speed
	a := Asteroid{
		GameObject{X: x, Y: y, VelX: vx, VelY: vy},
		3,
	}
	a.LoadSprite("spin-45.png")
	return a
}

/**************************************************************************
 *                                 Player                                 *
 **************************************************************************/
type Player struct {
	GameObject
	Angle float64
}