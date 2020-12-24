// Code generated by "stringer -type=Type"; DO NOT EDIT.

package layer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AccA-0]
	_ = x[AccB-1]
	_ = x[Beard-2]
	_ = x[Beard1-3]
	_ = x[Beard2-4]
	_ = x[BeastEars-5]
	_ = x[Bitten-6]
	_ = x[Body-7]
	_ = x[Cloak-8]
	_ = x[Cloak1-9]
	_ = x[Cloak2-10]
	_ = x[Clothing-11]
	_ = x[Clothing1-12]
	_ = x[Clothing2-13]
	_ = x[Ears-14]
	_ = x[Emotion-15]
	_ = x[Eyebrows-16]
	_ = x[Eyes-17]
	_ = x[Face-18]
	_ = x[FacialMark-19]
	_ = x[FrontHair-20]
	_ = x[FrontHair1-21]
	_ = x[FrontHair2-22]
	_ = x[Glasses-23]
	_ = x[Head-24]
	_ = x[Mouth-25]
	_ = x[Nose-26]
	_ = x[RearHair-27]
	_ = x[RearHair1-28]
	_ = x[RearHair2-29]
	_ = x[Scar-30]
	_ = x[Soil-31]
	_ = x[Tail-32]
	_ = x[Tail1-33]
	_ = x[Tail2-34]
	_ = x[Wing-35]
	_ = x[Wing1-36]
	_ = x[Wing2-37]
}

const _Type_name = "AccAAccBBeardBeard1Beard2BeastEarsBittenBodyCloakCloak1Cloak2ClothingClothing1Clothing2EarsEmotionEyebrowsEyesFaceFacialMarkFrontHairFrontHair1FrontHair2GlassesHeadMouthNoseRearHairRearHair1RearHair2ScarSoilTailTail1Tail2WingWing1Wing2"

var _Type_index = [...]uint8{0, 4, 8, 13, 19, 25, 34, 40, 44, 49, 55, 61, 69, 78, 87, 91, 98, 106, 110, 114, 124, 133, 143, 153, 160, 164, 169, 173, 181, 190, 199, 203, 207, 211, 216, 221, 225, 230, 235}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
