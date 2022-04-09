package copier

import (
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

type Student struct {
	ID    string
	Name  string
	Owner *Student
}

type Teacher struct {
	ID    string
	Name  string
	Owner *Teacher
}

func TestCopier(t *testing.T) {
	s := &Student{"1", "Wangtao", &Student{"2", "wt", nil}}
	th := &Teacher{}
	copier.Copy(th, s)
	fmt.Printf("t: %v\n", th)
	fmt.Printf("th: %v\n", th.Owner)
}
