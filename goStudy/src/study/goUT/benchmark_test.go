package testing
import (
    "testing"
    "bytes"
)
// window测试命令为 go test -v -bench="."
// lunix下使用 go test 命令行时，-bench=.应写为-bench="."。
// -benchmen 16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配
//go test -v polymorphsim_test.go >> 1.txt 将基准测试结果打印到文件中，后续使用
func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)

		}
	}
	b.StopTimer()

}