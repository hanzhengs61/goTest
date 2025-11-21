package main

import "testing"

func TestHello(t *testing.T) {
	// 简单调用方法与比对,失败打印失败原因.
	/*got := Hello("Golang")
	want := "Hello,Golang"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}*/

	// 开启多个测试用例,可以测试不同场景的同一个方法.
	/*t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Golang")
		want := "Hello,Golang"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
	t.Run("say 'Hello, Golang' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello,Golang"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})*/

	// 封装测试方法,简化代码量.
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		// 标记当前函数为测试辅助函数,当测试失败时,错误信息会出现在调用方的地方,而不是在这里.
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	/*t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Golang")
		want := "Hello,Golang"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, Golang' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello,Golang"
		assertCorrectMessage(t, got, want)
	})*/

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola,Elodie"
		assertCorrectMessage(t, got, want)
	})

}
