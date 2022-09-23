package rand

import (
	"fmt"
	"math/rand"
	"time"
)

type randClass struct {
	CharType []string
}

/**
 * rand
 * @Description: 随机生成Strlen长度的,包含include种字符的随机字符串
 * @receiver r
 * @param len
 * @param include
 */
func (rc randClass) Rand(Strlen, include int) (string, error) {
	var index int
	if Strlen < include {
		return "", fmt.Errorf("密码的总长度小于随机字符种类")
	}
	if include > len(rc.CharType) {
		return "", fmt.Errorf("随机字符种类大于字符种类")
	}
	//将所以字符种类填充进集合
	maps := make(map[int]interface{})
	for i := 0; i < len(rc.CharType); i++ {
		maps[i] = nil
	}
	//删除多余的set
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	for i := 0; i < len(rc.CharType)-include; i++ {
		delete(maps, r.Intn(len(maps)))
	}
	bytesMap := make(map[int]byte)

	//按include每种字符先填一次
	for k, _ := range maps {
		index = r.Intn(len(rc.CharType[k]))
		bytesMap[len(bytesMap)+1] = rc.CharType[k][index]
	}

	//将随机产生的字符map转化为slice,方便生成随机值
	charTypes := make([]int, len(maps))
	i := 0
	for k, _ := range maps {
		charTypes[i] = k
	}

	//随机生成Strlen-include的值填充进入map
	for i := 0; i < Strlen-include; i++ {
		index = r.Intn(len(charTypes))
		bytesMap[len(bytesMap)+1] = rc.CharType[charTypes[index]][r.Intn(len(rc.CharType))]
	}

	//map转换slice
	bytes := make([]byte, Strlen)
	i = 0
	for _, v := range bytesMap {
		bytes[i] = v
		i++
	}

	return string(bytes), nil
}

func New(strs []string) randClass {
	return randClass{CharType: strs}
}
