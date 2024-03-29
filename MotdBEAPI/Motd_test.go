/*
 * @Author: NyanCatda
 * @Date: 2024-03-30 02:33:22
 * @LastEditTime: 2024-03-30 02:34:31
 * @LastEditors: NyanCatda
 * @Description: 测试用例
 * @FilePath: \MCBE-Server-Motd\MotdBEAPI\Motd_test.go
 */
package MotdBEAPI

import "testing"

func TestBE(t *testing.T) {
	Host := ""

	Data, err := MotdBE(Host)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(Data)
}

func TestJava(t *testing.T) {
	Host := ""

	Data, err := MotdJava(Host)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(Data)
}
