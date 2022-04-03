package ungar

import "testing"

func TestConvert(t *testing.T) {
	testCases := []struct {
		desription string
		input      string
		expected   [7]uint16
	}{
		{
			"Sevens full", 
			"As 7s 7c 7h Ad 8d 2h", 
			[7]uint16{335, 111, 109, 106, 331, 139, 18},
		},
		{
			"Four of a kind: sevens", 
			"7s 7c 7d 7h 3c 3d As", 
			[7]uint16{111, 109, 107, 106, 29, 27, 335},
		},
		{
			"Straight Ace to Five", 
			"Ah 2s 3s 4h 5d Jd Qs", 
			[7]uint16{330, 23, 31, 42, 59, 235, 255},
		},
		{
			"Wrong combination of seven Aces", 
			"Ac Ac Ac Ac Ac Ac Ac", 
			[7]uint16{333, 333, 333, 333, 333, 333, 333},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desription, func(t *testing.T) {
			if result := Convert(testCase.input); result != testCase.expected {
				t.Errorf("Wrong conversion! Got: %v, expected: %v", result, testCase.expected)
			}
		})
	}
}
