package utils

import (
	"hash/fnv"
	"strings"

	"github.com/sony/sonyflake"

	"github.com/google/uuid"
)

func GenerateUUIDBytes() []byte {
	u := uuid.New()
	return u[:]
}

var MachineID uint16

func GetMachineID() {
	MachineID = ReadMachineIDFromFile()

}
func ReadMachineIDFromFile() uint16 {
	data := GenerateUUIDBytes()
	h := fnv.New32()
	h.Write(data)
	return uint16(h.Sum32() & 0xffff)

}

var Sf *sonyflake.Sonyflake

func SF() {
	Sf, _ = sonyflake.New(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return MachineID, nil
		},
	})

}

// 生成UUID，去除其中的连字符
func GenerateUUIDWithoutHyphen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
