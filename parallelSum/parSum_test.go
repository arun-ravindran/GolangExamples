package parallelSum

import (
	"fmt"
	"testing"
)

func test(t *testing.T, fileName string, num int, expected int) {
	result := sum(num, fileName)
	if result != expected {
		t.Fatal(fmt.Sprintf(
			"Sum of %s failed: got %d, expected %d\n", fileName, result, expected))
	}
}

func Test0(t *testing.T) {
	test(t, "parSum_test0.txt", 2, 10)
}

func Test1(t *testing.T) {
	test(t, "parSum_test1.txt", 1, 499500)
}

func Test2(t *testing.T) {
	test(t, "parSum_test1.txt", 10, 499500)
}

func Test3(t *testing.T) {
	test(t, "parSum_test2.txt", 1, 117652)
}

func Test4(t *testing.T) {
	test(t, "parSum_test2.txt", 10, 117652)
}
