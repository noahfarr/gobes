package gobes

import (
	"math"
)

// Vector3 Holds a vector in 3 dimensions.
type Vector3 struct {
	// Holds the value along the X axis
	X float64
	// Holds the value along the Y axis
	Y float64
	// Holds the value along the Z axis
	Z float64
}

// Calculates and returns the sum of two given vectors
func (v Vector3) addVector(w Vector3) Vector3 {
	return Vector3{
		v.X + w.X,
		v.Y + w.Y,
		v.Z + w.Z,
	}
}

// Calculates and returns the sum of a given vector and a scalar
func (v Vector3) addScalar(s float64) Vector3 {
	return Vector3{
		v.X + s,
		v.Y + s,
		v.Z + s,
	}
}

// Calculates and returns the sum of two vectors with the second one scaled by the given amount
func (v Vector3) addScaledVector(w Vector3, s float64) Vector3 {
	return v.addVector(w.multiplyScalar(s))
}

// Calculates and returns the difference of two given vectors
func (v Vector3) subVector(w Vector3) Vector3 {
	return Vector3{
		v.X - w.X,
		v.Y - w.Y,
		v.Z - w.Z,
	}
}

// Calculates and returns the difference of a given vector and a scalar
func (v Vector3) subScalar(s float64) Vector3 {
	return Vector3{
		v.X - s,
		v.Y - s,
		v.Z - s,
	}
}

// Calculates and returns the product of a given vector and a scalar
func (v Vector3) multiplyScalar(s float64) Vector3 {
	return Vector3{
		v.X * s,
		v.Y * s,
		v.Z * s,
	}
}

// Calculates and returns the scalar product of two given vectors
func (v Vector3) scalarProduct(w Vector3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Calculates and returns the cross product of two given vectors
func (v Vector3) crossProduct(w Vector3) Vector3 {
	return Vector3{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// Flips all the components of the vector
func (v Vector3) invert() Vector3 {
	return Vector3{
		-v.X,
		-v.Y,
		-v.Z,
	}
}

// Returns the magnitude/length of the vector
func (v Vector3) magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Turns a non-zero vector into a vector of unit length
func (v Vector3) normalize() Vector3 {
	if length := v.magnitude(); length > 0 {
		return v.multiplyScalar(1 / length) // v = v/length
	}
	return v
}
