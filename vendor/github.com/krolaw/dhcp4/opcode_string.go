// Code generated by "stringer -type=OpCode"; DO NOT EDIT.

package dhcp4

import "strconv"

const _OpCode_name = "BootRequestBootReply"

var _OpCode_index = [...]uint8{0, 11, 20}

func (i OpCode) String() string {
	i -= 1
	if i >= OpCode(len(_OpCode_index)-1) {
		return "OpCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _OpCode_name[_OpCode_index[i]:_OpCode_index[i+1]]
}
