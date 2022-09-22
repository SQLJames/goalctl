// cgo -godefs types_illumos.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build amd64 && illumos
// +build amd64,illumos

package unix

const (
	TUNNEWPPA = 0x540001
	TUNSETPPA = 0x540002

	I_STR     = 0x5308
	I_POP     = 0x5303
	I_PUSH    = 0x5302
	I_LINK    = 0x530c
	I_UNLINK  = 0x530d
	I_PLINK   = 0x5316
	I_PUNLINK = 0x5317

	IF_UNITSEL = -0x7ffb8cca
)

type strbuf struct {
	Maxlen int32
	Len    int32
	Buf    *int8
}

type Strioctl struct {
	Cmd    int32
	Timout int32
	Len    int32
	Dp     *int8
}

type Lifreq struct {
	Name   [32]int8
	Lifru1 [4]byte
	Type   uint32
	Lifru  [336]byte
}
