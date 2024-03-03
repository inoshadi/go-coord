/**
 * Author: Sutrisno Hadi
 * File: go-coord.go
 */

package coord

import "testing"

func TestCountDistance(t *testing.T) {
	a := Coord{
		Lat: -6.157690584152987,
		Lng: 106.59008232643828,
	}
	b := Coord{
		Lat: -6.209604,
		Lng: 106.632629,
	}

	expected := 7446.398397265567
	if got := CountDistance(a, b); got != expected {
		t.Errorf("CountDistance(%v, %v) = %v, didn't return %v", a, b, got, expected)
	}
}
