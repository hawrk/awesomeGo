// Package bench
/*
 Author: hawrkchen
 Date: 2022/3/8 10:54
 Desc:
*/
package bench

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// 用 jsoniter 替换原生的json 会有很大的性能提升
/*
import "encoding/json"
json.Marshal(&data)
替换成
import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Marshal(&data)

import "encoding/json"
json.Unmarshal(input, &data)
替换成
import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Unmarshal(input, &data)

*/

type ColorGroup struct {
	ID      int `json:"id"`
	Name    string `json:"name"`
	Colors  []string `json:"colors"`
}

func JsoniterEncode() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	_, err := jsoniter.Marshal(group)
	if err != nil {
		fmt.Println("jsoniter marshal err :", err)
	}
	//fmt.Println("jsoniter marshal out:", string(b))   // 输出字符串，否则输出二进制
}

func JsonEncode() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	_, err := json.Marshal(group)
	if err != nil {
		fmt.Println(" json marshal err:", err)
	}
	//fmt.Println("json marshal out:", string(b))
}

