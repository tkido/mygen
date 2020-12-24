// Code generated by "stringer -type=Type"; DO NOT EDIT.

package part

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AccA-0]
	_ = x[AccB-1]
	_ = x[BeastEars-2]
	_ = x[Beard-3]
	_ = x[Body-4]
	_ = x[Clothing-5]
	_ = x[Cloak-6]
	_ = x[Ears-7]
	_ = x[Eyebrows-8]
	_ = x[Eyes-9]
	_ = x[Face-10]
	_ = x[FacialMark-11]
	_ = x[FrontHair-12]
	_ = x[Glasses-13]
	_ = x[Head-14]
	_ = x[Mouth-15]
	_ = x[Nose-16]
	_ = x[RearHair-17]
	_ = x[Soil-18]
	_ = x[Tail-19]
	_ = x[Wing-20]
}

const _Type_name = "AccAAccBBeastEarsBeardBodyClothingCloakEarsEyebrowsEyesFaceFacialMarkFrontHairGlassesHeadMouthNoseRearHairSoilTailWing"

var _Type_index = [...]uint8{0, 4, 8, 17, 22, 26, 34, 39, 43, 51, 55, 59, 69, 78, 85, 89, 94, 98, 106, 110, 114, 118}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
