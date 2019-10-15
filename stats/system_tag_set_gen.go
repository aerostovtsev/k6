// Code generated by "enumer -type=SystemTagSet -transform=snake -trimprefix=Tag -output system_tag_set_gen.go"; DO NOT EDIT.

//
package stats

import (
	"fmt"
)

const _SystemTagSetName = "protosubprotostatusmethodurlnamegroupcheckerrorerror_codetls_versionitervuocsp_statusip"

var _SystemTagSetMap = map[SystemTagSet]string{
	1:     _SystemTagSetName[0:5],
	2:     _SystemTagSetName[5:13],
	4:     _SystemTagSetName[13:19],
	8:     _SystemTagSetName[19:25],
	16:    _SystemTagSetName[25:28],
	32:    _SystemTagSetName[28:32],
	64:    _SystemTagSetName[32:37],
	128:   _SystemTagSetName[37:42],
	256:   _SystemTagSetName[42:47],
	512:   _SystemTagSetName[47:57],
	1024:  _SystemTagSetName[57:68],
	2048:  _SystemTagSetName[68:72],
	4096:  _SystemTagSetName[72:74],
	8192:  _SystemTagSetName[74:85],
	16384: _SystemTagSetName[85:87],
}

func (i SystemTagSet) String() string {
	if str, ok := _SystemTagSetMap[i]; ok {
		return str
	}
	return fmt.Sprintf("SystemTagSet(%d)", i)
}

var _SystemTagSetValues = []SystemTagSet{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384}

var _SystemTagSetNameToValueMap = map[string]SystemTagSet{
	_SystemTagSetName[0:5]:   1,
	_SystemTagSetName[5:13]:  2,
	_SystemTagSetName[13:19]: 4,
	_SystemTagSetName[19:25]: 8,
	_SystemTagSetName[25:28]: 16,
	_SystemTagSetName[28:32]: 32,
	_SystemTagSetName[32:37]: 64,
	_SystemTagSetName[37:42]: 128,
	_SystemTagSetName[42:47]: 256,
	_SystemTagSetName[47:57]: 512,
	_SystemTagSetName[57:68]: 1024,
	_SystemTagSetName[68:72]: 2048,
	_SystemTagSetName[72:74]: 4096,
	_SystemTagSetName[74:85]: 8192,
	_SystemTagSetName[85:87]: 16384,
}

// SystemTagSetString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func SystemTagSetString(s string) (SystemTagSet, error) {
	if val, ok := _SystemTagSetNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to SystemTagSet values", s)
}

// SystemTagSetValues returns all values of the enum
func SystemTagSetValues() []SystemTagSet {
	return _SystemTagSetValues
}

// IsASystemTagSet returns "true" if the value is listed in the enum definition. "false" otherwise
func (i SystemTagSet) IsASystemTagSet() bool {
	_, ok := _SystemTagSetMap[i]
	return ok
}
