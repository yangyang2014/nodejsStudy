package testing
import (
	// "fmt"
	// "testing"
	// "bytes"
	// go get -u github.com/stretchr/testify/assert
// 	"github.com/stretchr/testify/assert"
)
//文件名称形如 xxx_test.go; 函数方法形如 Testxxx(t *testing.TS){T.Errorf()}
//运行时使用 go test（打印日志和覆盖率 go test -v -cover）
// func TestSquare(t *testing.T) {
// 	inputs := [...]int{1,2,3}
// 	expected := [...]int{1,4,9}
// 	for i := 0 ; i<len(inputs); i++ {
// 		ret:=Square(inputs[i]);
// 		if ret != expected[i] {
// 			t.Errorf("input is %d,the expected is %d, the actual %d",inputs[i],expected[i],ret)
// 		}
// 	}
// }


// func TestFatalInCode(t *testing.T){
// 	fmt.Println("Start")
// 	t.Fatal("Fatal")
// 	fmt.Println("Error")
// }

// func TestErrorInCode(t *testing.T){
// 	fmt.Println("Start2")
// 	t.Error("error")
// 	fmt.Println("End2")
// }
// //BenchMark 性能测试
// /**
// func BenchmarkConcatStringByAdd(b *testing.B) {
// 	//与性能测试无关的代码
// 	b.ResetTime()
// 	for i:=0; i<b.N; i++ {
// 		//测试代码
// 	}
// 	b.StopTimer();

// }

// go test -bench=.
// **/
// func TestConcatStringByAdd(t *testing.T) {
// 	elems := []string{"1","2","3","4","5"}
// 	ret := ""
// 	for _, elem := range elems {
// 		ret += elem
// 	}
// 	if(ret != "12345") {
// 		t.Errorf("input is %d,the expected is %d, the actual %d",elems,"12345",ret)
// 	}
// }

// func TestConcatStringByBytesBuffer(t *testing.T) {
// 	elems := []string{"1","2","3","4","5"}
// 	var buf bytes.Buffer
	
// 	for _, elem := range elems {
// 		buf.WriteString(elem)
// 	}
// 	if(buf.String() != "12345") {
// 		t.Errorf("input is %d,the expected is %d, the actual %d",elems,"12345",buf.String())
// 	}
// }

// func BenchmarkConcatStringByAdd(b *testing.B) {
// 	elems := []string{"1","2","3","4","5"}
// 	b.ResetTimer()
// 	for i:=0 ; i<b.N; i++{
		
// 		ret := ""
// 		for _, elem := range elems {
// 			ret += elem
// 		}
// 	}
// 	b.StopTimer()
// }

// func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
// 	elems := []string{"1","2","3","4","5"}
// 	b.ResetTimer()
// 	for i:=0 ; i<b.N; i++{
// 		var buf bytes.Buffer
// 		for _, elem := range elems {
// 			buf.WriteString(elem)
// 		}
// 	}
// 	b.StopTimer()
// }