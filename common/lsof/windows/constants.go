// +build windows

package windows

const (
	AF_INET  = 2
	AF_INET6 = 23
)

const (
	MAX_MODULE_NAME32 = 255
	MAX_PATH          = 260
)

// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
const (
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010
	TH32CS_INHERIT      = 0x80000000
	TH32CS_SNAPALL      = TH32CS_SNAPHEAPLIST | TH32CS_SNAPMODULE | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD
)

const (
	MAX_ADAPTER_NAME       = 128
	MAX_INTERFACE_NAME_LEN = 256
	MAXLEN_PHYSADDR        = 8
	MAXLEN_IFDESCR         = 256
)
