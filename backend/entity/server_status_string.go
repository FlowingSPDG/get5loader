// Code generated by "stringer -type=SERVER_STATUS"; DO NOT EDIT.

package entity

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SERVER_STATUS_UNKNOWN-0]
	_ = x[SERVER_STATUS_STANDBY-1]
	_ = x[SERVER_STATUS_INUSE-2]
}

const _SERVER_STATUS_name = "SERVER_STATUS_UNKNOWNSERVER_STATUS_STANDBYSERVER_STATUS_INUSE"

var _SERVER_STATUS_index = [...]uint8{0, 21, 42, 61}

func (i SERVER_STATUS) String() string {
	if i < 0 || i >= SERVER_STATUS(len(_SERVER_STATUS_index)-1) {
		return "SERVER_STATUS(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SERVER_STATUS_name[_SERVER_STATUS_index[i]:_SERVER_STATUS_index[i+1]]
}
