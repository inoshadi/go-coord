/**
 * Author: Sutrisno Hadi
 * File: go-coord.go
 */

package coord

import "math"

type (
	// Coord struct consist of Latitude and Longitude coordinates
	Coord struct {
		Lat float64
		Lng float64
	}
)

const (
	// Corr is the correction value to get the nearest result
	Corr = 0.1059
)

// CountDistance between two coordinates
func CountDistance(coord1 Coord, coord2 Coord) float64 {
	plt := math.Pow(coord1.Lat-coord2.Lat, 2)
	plg := math.Pow(coord1.Lng-coord2.Lng, 2)
	sqrt := math.Sqrt(plt + plg)
	mtr := sqrt * 100000
	return mtr + (mtr * Corr)
}
