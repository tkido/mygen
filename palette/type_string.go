// Code generated by "stringer -type=Type"; DO NOT EDIT.

package palette

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Skin-1]
	_ = x[Eyes-2]
	_ = x[Hair-3]
	_ = x[HairSub-4]
	_ = x[FacialMark-5]
	_ = x[BeastEars-6]
	_ = x[Clothing-7]
	_ = x[ClothingSub1-8]
	_ = x[ClothingSub2-9]
	_ = x[ClothingSub3-10]
	_ = x[Cloak-11]
	_ = x[CloakSub-12]
	_ = x[AccA-13]
	_ = x[AccASub1-14]
	_ = x[AccASub2-15]
	_ = x[AccB-16]
	_ = x[AccBSub1-17]
	_ = x[AccBSub2-18]
	_ = x[AccBSub3-19]
	_ = x[Glasses-20]
	_ = x[GlassesSub1-21]
	_ = x[GlassesSub2-22]
	_ = x[Tail-23]
	_ = x[Wing-24]
}

const _Type_name = "SkinEyesHairHairSubFacialMarkBeastEarsClothingClothingSub1ClothingSub2ClothingSub3CloakCloakSubAccAAccASub1AccASub2AccBAccBSub1AccBSub2AccBSub3GlassesGlassesSub1GlassesSub2TailWing"

var _Type_index = [...]uint8{0, 4, 8, 12, 19, 29, 38, 46, 58, 70, 82, 87, 95, 99, 107, 115, 119, 127, 135, 143, 150, 161, 172, 176, 180}

func (i Type) String() string {
	i -= 1
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
