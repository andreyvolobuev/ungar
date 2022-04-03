package ungar

import "testing"

var testCases = []struct{
        description string
        input       [7]uint16
        expected    uint16
    }{
        {
            "Qs Js 9c 6d 4c 3s 2h",
            [7]uint16{255, 239, 157, 91, 45, 31, 18},
            uint16(7047),
        },
        {
            "Ks Js 9c 6d 4c 3s 2h",
            [7]uint16{303, 239, 157, 91, 45, 31, 18},
            uint16(6838),
        },
        {
            "Ks Js 9c Td 8c 3s 2h",
            [7]uint16{303, 239, 157, 187, 141, 31, 18},
            uint16(6798),
        },
        {
            "Ks Js 9c Td 8c 6s 2h",
            [7]uint16{303, 239, 157, 187, 141, 95, 18},
            uint16(6798),
        },
        {
            "Ks Js 9c Td Tc 6s 2h",
            [7]uint16{303, 239, 157, 187, 189, 95, 18},
            uint16(4270),
        },
        {
            "Ks Js 9c Jd Tc 8s 2h",
            [7]uint16{303, 239, 157, 235, 189, 143, 18},
            uint16(4050),
        },
        {
            "Ks Js Kc Jd Tc 8s 2h",
            [7]uint16{303, 239, 301, 235, 189, 143, 18},
            uint16(2613),
        },
        {
            "Ks Js Kc Jd Qc Qs 2h",
            [7]uint16{303, 239, 301, 235, 253, 255, 18},
            uint16(2601),
        },
        {
            "Ks Js Kc Jd Qc Qs Ah",
            [7]uint16{303, 239, 301, 235, 253, 255, 330},
            uint16(2600),
        },
        {
            "Ks Js Kc Jd Qc Qs Th",
            [7]uint16{303, 239, 301, 235, 253, 255, 186},
            uint16(2601),
        },
        {
            "7c 7d 7s 9c Js Qh 3d",
            [7]uint16{109, 107, 111, 157, 239, 250, 27},
            uint16(2168),
        },
        {
            "7c 7d 7s Tc Js Qh 8d",
            [7]uint16{109, 107, 111, 189, 239, 250, 139},
            uint16(2168),
        },
        {
            "7c 7d 7s Tc Ks Qh 8d",
            [7]uint16{109, 107, 111, 189, 303, 250, 139},
            uint16(2159),
        },
        {
            "7c 7d 7s Tc Ks Qh Jd",
            [7]uint16{109, 107, 111, 189, 303, 250, 139},
            uint16(2159),
        },
        {
            "7c 7d 6s 5c 4s 3h Kd",
            [7]uint16{109, 107, 95, 61, 47, 26, 299},
            uint16(1607),
        },
        {
            "7c 7d 6s 5c 4s 3h Ad",
            [7]uint16{109, 107, 95, 61, 47, 26, 331},
            uint16(1607),
        },
        {
            "7c Ad 6s 5c 4s 3h 2d",
            [7]uint16{109, 331, 95, 61, 47, 26, 19},
            uint16(1607),
        },
        {
            "7c Ad As 5c 4s 3h 2d",
            [7]uint16{109, 331, 335, 61, 47, 26, 19},
            uint16(1609),
        },
        {
            "7c Ac As 5c Jc Qc 2d",
            [7]uint16{109, 333, 335, 61, 237, 253, 19},
            uint16(509),
        },
        {
            "7c Ac As 5c Jc Qc Ah",
            [7]uint16{109, 333, 335, 61, 237, 253, 330},
            uint16(509),
        },
        {
            "7c Ac As 5c Qd Qc Qh",
            [7]uint16{109, 333, 335, 61, 251, 253, 250},
            uint16(191),
        },
        {
            "7c Kc Ks Kh 2d 2c 2h",
            [7]uint16{109, 301, 303, 298, 19, 21, 18},
            uint16(190),
        },
        {
            "7c Ac As Ah 2d 2c 2h",
            [7]uint16{109, 333, 335, 330, 19, 21, 18},
            uint16(178),
        },
        {
            "Kc Ac As Ah 2d 2c 2h",
            [7]uint16{301, 333, 335, 330, 19, 21, 18},
            uint16(178),
        },
        {
            "Ad Ac As Ah 2d 2c 2h",
            [7]uint16{331, 333, 335, 330, 19, 21, 18},
            uint16(22),
        },
        {
            "Ad Ac As Ah Kd Kc Kh",
            [7]uint16{331, 333, 335, 330, 299, 301, 298},
            uint16(11),
        },
        {
            "Ad Ac As Ah Kd Qc Qh",
            [7]uint16{331, 333, 335, 330, 299, 253, 250},
            uint16(11),
        },
        {
            "Ad 2d 3d 4d 5d Qc Qh",
            [7]uint16{331, 19, 27, 43, 59, 253, 250},
            uint16(10),
        },
        {
            "Ad 2d 3d 4d 5d 6d Qh",
            [7]uint16{331, 19, 27, 43, 59, 91, 250},
            uint16(9),
        },
        {
            "Kd 2d 3d 4d 5d 6d Qh",
            [7]uint16{299, 19, 27, 43, 59, 91, 250},
            uint16(9),
        },
        {
            "Kd Ad Jd Td 5c 6s Qd",
            [7]uint16{299, 331, 235, 187, 61, 95, 251},
            uint16(1),
        },
    }

func TestEval (t *testing.T) {
    for _, testCase := range testCases {
        t.Run(testCase.description, func(t *testing.T) {
            if result := Evaluate(testCase.input); result != testCase.expected {
                t.Errorf("Wrong evaluation! Got: %v, expected: %v", result, testCase.expected)
            }
        })
    }
}
