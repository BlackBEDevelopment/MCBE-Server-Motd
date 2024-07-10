/*
 * @Author: NyanCatda
 * @Date: 2024-03-30 02:33:22
 * @LastEditTime: 2024-07-10 17:32:54
 * @LastEditors: NyanCatda
 * @Description: 测试用例
 * @FilePath: \MCBE-Server-Motd\MotdBEAPI\Motd_test.go
 */
package MotdBEAPI

import (
	"encoding/json"
	"testing"
)

func TestBE(t *testing.T) {
	Host := ""

	Data, err := MotdBE(Host)
	if err != nil {
		t.Error(err)
		return
	}

	DataJson, err := json.Marshal(Data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(DataJson))
}

func TestJava(t *testing.T) {
	Host := ""

	Data, err := MotdJava(Host)
	if err != nil {
		t.Error(err)
		return
	}

	DataJson, err := json.Marshal(Data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(DataJson))
}
