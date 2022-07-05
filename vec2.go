package gobes

import (
	"math"
)

// Vector2 Holds a vector in 2 dimensions.
type Vector2 struct {
	// Holds the value along the X axis
	X float64
	// Holds the value along the Y axis
	Y float64
}

// Calculates and returns the sum of two given vectors
func (v Vector2) addVector(w Vector2) Vector2 {
	return Vector2{
		v.X + w.X,
		v.Y + w.Y,
	}
}

// Calculates and returns the sum of a given vector and a scalar
func (v Vector2) addScalar(s float64) Vector2 {
	return Vector2{
		v.X + s,
		v.Y + s,
	}
}

// Calculates and returns the sum of two vectors with the second one scaled by the given amount
func (v Vector2) addScaledVector(w Vector2, s float64) Vector2 {
	return v.addVector(w.multiplyScalar(s))
}

// Calculates and returns the difference of two given vectors
func (v Vector2) subVector(w Vector2) Vector2 {
	return Vector2{
		v.X - w.X,
		v.Y - w.Y,
	}
}

// Calculates and returns the difference of a given vector and a scalar
func (v Vector2) subScalar(s float64) Vector2 {
	return Vector2{
		v.X - s,
		v.Y - s,
	}
}

// Calculates and returns the product of a given vector and a scalar
func (v Vector2) multiplyScalar(s float64) Vector2 {
	return Vector2{
		v.X * s,
		v.Y * s,
	}
}

// Calculates and returns the scalar product of two given vectors
func (v Vector2) scalarProduct(w Vector2) float64 {
	return v.X*w.X + v.Y*w.Y
}

// Flips all the components of the vector
func (v Vector2) invert() Vector2 {
	return Vector2{
		-v.X,
		-v.Y,
	}
}

// Returns the magnitude/length of the vector
func (v Vector2) magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Turns a non-zero vector into a vector of unit length
func (v Vector2) normalize() Vector2 {
	if length := v.magnitude(); length > 0 {
		return v.multiplyScalar(1 / length) // v = v/length
	}
	return v
}
