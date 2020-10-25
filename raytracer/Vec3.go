package raytracer

import "math"

// Vec3 (x, y, z)
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

//NewVec3 return a new Vec3
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// Add v2 to v
func (v *Vec3) Add(v2 Vec3) *Vec3 {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
	return v
}

// Add return v + v2
func Add(v Vec3, v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

// Neg return -v
func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Neg return v - v2
func Neg(v Vec3, v2 Vec3) Vec3 {
	return Vec3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

// Mul v * float
func (v *Vec3) Mul(f float64) *Vec3 {
	v.X *= f
	v.Y *= f
	v.Z *= f
	return v
}

// Mul return v * v2
func Mul(v Vec3, v2 Vec3) Vec3 {
	return Vec3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

// MulF return v * float
func MulF(v Vec3, f float64) Vec3 {
	return Vec3{v.X * f, v.Y * f, v.Z * f}
}

// Div v / float
func (v *Vec3) Div(f float64) *Vec3 {
	return v.Mul(1 / f)
}

// Div return v / float
func Div(v Vec3, f float64) Vec3 {
	return MulF(v, 1/f)
}

// Length return sqrt(LengthSquared)
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared return lengthSquared
func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Dot return Dot Product
func Dot(v Vec3, v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Cross return Vector product
func Cross(v Vec3, v2 Vec3) Vec3 {
	return Vec3{v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X}
}

// UnitVector return the unitVector
func (v Vec3) UnitVector() Vec3 {
	return Div(v, v.Length())
}
