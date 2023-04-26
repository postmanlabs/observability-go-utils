package math

import (
	"math"
	"testing"
)

func TestRoundTo3SigFigs(t *testing.T) {
	testCases := map[float64]float64{
		0:      0,
		7:      7,
		42:     42,
		999:    999,
		1337:   1340,
		2344:   2340,
		2345:   2350,
		9994:   9990,
		9995:   10000,
		9999:   10000,
		10000:  10000,
		10049:  10000,
		10050:  10100,
		23449:  23400,
		23450:  23500,
		99949:  99900,
		99950:  100000,
		99999:  100000,
		100000: 100000,
		100499: 100000,
		100500: 101000,
	}

	epsilon := 1e-7
	for input, expected := range testCases {
		for _, sign := range []float64{1, -1} {
			for _, exp := range []int{0, 1, 2, 3, 4, -1, -2, -3, -4} {
				scale := math.Pow10(exp)
				actual := RoundToSigFigs(sign*input*scale, 3)
				if math.Abs(sign*expected*scale-actual) >= epsilon {
					t.Errorf("Expected RoundTo3SigFigs(%f) = %f but got %f", sign*input*scale, sign*expected*scale, actual)
				}
			}
		}
	}
}

func TestCeilTo3SigFigs(t *testing.T) {
	testCases := map[float64]float64{
		0:       0,
		7:       7,
		-7:      -7,
		42:      42,
		-42:     -42,
		999:     999,
		-999:    -999,
		1337:    1340,
		-1337:   -1330,
		2344:    2350,
		-2344:   -2340,
		2345:    2350,
		-2345:   -2340,
		9994:    10000,
		-9994:   -9990,
		9995:    10000,
		-9995:   -9990,
		9999:    10000,
		-9999:   -9990,
		10000:   10000,
		-10000:  -10000,
		10049:   10100,
		-10049:  -10000,
		10050:   10100,
		-10050:  -10000,
		23449:   23500,
		-23449:  -23400,
		23450:   23500,
		-23450:  -23400,
		99949:   100000,
		-99949:  -99900,
		99950:   100000,
		-99950:  -99900,
		99999:   100000,
		-99999:  -99900,
		100000:  100000,
		-100000: -100000,
		100499:  101000,
		-100499: -100000,
		100500:  101000,
		-100500: -100000,
	}

	epsilon := 1e-7
	for input, expected := range testCases {
		// XXX Don't scale the inputs. Unlike in the RoundToSigFigs tests, this test
		// is more sensitive to rounding error during scaling.
		actual := CeilToSigFigs(input, 3)
		if math.Abs(expected-actual) >= epsilon {
			t.Errorf("Expected CeilTo3SigFigs(%f) = %f but got %f", input, expected, actual)
		}
	}
}

func TestFloorTo3SigFigs(t *testing.T) {
	testCases := map[float64]float64{
		0:       0,
		7:       7,
		-7:      -7,
		42:      42,
		-42:     -42,
		999:     999,
		-999:    -999,
		1337:    1330,
		-1337:   -1340,
		2344:    2340,
		-2344:   -2350,
		2345:    2340,
		-2345:   -2350,
		9994:    9990,
		-9994:   -10000,
		9995:    9990,
		-9995:   -10000,
		9999:    9990,
		-9999:   -10000,
		10000:   10000,
		-10000:  -10000,
		10049:   10000,
		-10049:  -10100,
		10050:   10000,
		-10050:  -10100,
		23449:   23400,
		-23449:  -23500,
		23450:   23400,
		-23450:  -23500,
		99949:   99900,
		-99949:  -100000,
		99950:   99900,
		-99950:  -100000,
		99999:   99900,
		-99999:  -100000,
		100000:  100000,
		-100000: -100000,
		100499:  100000,
		-100499: -101000,
		100500:  100000,
		-100500: -101000,
	}

	epsilon := 1e-7
	for input, expected := range testCases {
		// XXX Don't scale the inputs. Unlike in the RoundToSigFigs tests, this test
		// is more sensitive to rounding error during scaling.
		actual := FloorToSigFigs(input, 3)
		if math.Abs(expected-actual) >= epsilon {
			t.Errorf("Expected floorTo3SigFigs(%f) = %f but got %f", input, expected, actual)
		}
	}
}
