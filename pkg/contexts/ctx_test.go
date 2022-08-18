/**
@author ZhengHao Lou
Date    2022/08/18
*/
package contexts

import (
	"context"
	"fmt"
)

func ExampleCtxWithValues() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "a")
	ctx = context.WithValue(ctx, "key2", "b")
	var doSomething = func(ctx context.Context) {
		fmt.Printf("doSomething, key1: %v\n", ctx.Value("key1"))
		fmt.Printf("doSomething, key2: %v\n", ctx.Value("key2"))
	}
	
	doSomething(ctx)
	// Output:
	// doSomething, key1: a
	// doSomething, key2: b
}
