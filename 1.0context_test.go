package learn_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContexts(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
