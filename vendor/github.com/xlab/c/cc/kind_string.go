// generated by stringer -type Kind; DO NOT EDIT

package cc

import "fmt"

const _Kind_name = "UndefinedVoidPtrCharSCharUCharShortUShortIntUIntLongULongLongLongULongLongFloatDoubleLongDoubleBoolFloatComplexDoubleComplexLongDoubleComplexStructUnionEnumTypedefNameFunctionArray"

var _Kind_index = [...]uint8{0, 9, 13, 16, 20, 25, 30, 35, 41, 44, 48, 52, 57, 65, 74, 79, 85, 95, 99, 111, 124, 141, 147, 152, 156, 167, 175, 180}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return fmt.Sprintf("Kind(%d)", i)
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
