package generatestring

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	res := GenerateWord("siapa")
	fmt.Println(len(res))
	fmt.Println(res)
}
