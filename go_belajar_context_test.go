package belajargolanjutan

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	// ada dua cara membuat context dengan background dan todo
	// keduanya sama" membuat context kosong
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
