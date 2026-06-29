# Solutions Archive

Verified solutions in order of completion.

---

## F01 — Launch a goroutine, print "hello" from inside it

> Solution not captured in git history.

---

## F02 — Launch a goroutine and wait for it to finish using WaitGroup

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go train(&wg)

	wg.Wait()
}

func train(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from Goroutine")
}
```

---

## F03 — Send a value into an unbuffered channel from a goroutine, receive in main

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	v := <-ch
	fmt.Println(v)
}
```

---

## F04 — Create a buffered channel of size 3, send 3 values without blocking

> Solution not cleanly captured in git history.

---

## F05 — Close a channel after sending; iterate its values with range

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	var once sync.Once

	closeChannel := func() {
		once.Do(func() {
			close(ch)
		})
	}

	ch <- 1
	ch <- 2
	ch <- 3
	go closeChannel()
	go closeChannel()

	for val := range ch {
		fmt.Println(val)
	}
}
```

---

## F06 — Use select to receive from whichever of two channels is ready first

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go demo(ch, 1)
	go demo(ch2, 2)

	select {
	case v := <-ch:
		fmt.Println(v)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}

func demo(ch chan int, value int) {
	ch <- value
}
```

---

## F07 — Use select with a default case to do a non-blocking receive

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)

	go demo(ch, 1)

	select {
	case v := <-ch:
		fmt.Println(v)
	case ch2 <- 2:
		fmt.Println("send success")
	default:
		fmt.Println("default run")
	}
}

func demo(ch chan int, value int) {
	ch <- value
}
```

---

## F08 — Protect a shared integer counter with sync.Mutex across 5 goroutines

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go inc(&counter, &mu, &wg)
	}

	wg.Wait()
	fmt.Println(counter)
}

func inc(n *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer wg.Done()
	defer mu.Unlock()
	*n++
}
```

---

## F09 — Use sync.Once to initialize a config struct exactly once

```go
package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Port string
	Msg  string
}

var once sync.Once
var config Config

func Init(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(func() {
		config = Config{Port: "8080", Msg: "Initialized"}
		fmt.Println(config)
	})
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go Init(&wg)
	go Init(&wg)
	go Init(&wg)

	wg.Wait()
}
```

---

## F10 — Use sync.WaitGroup to launch 5 goroutines and wait for all to finish

```go
package main

import (
	"fmt"
	"sync"
)

func Init(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("run ", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Init(&wg, i)
	}

	wg.Wait()
}
```

---

## F17 — Define a sentinel error with errors.New; check it with errors.Is

```go
package main

import (
	"errors"
	"fmt"
)

var UserError = errors.New("user not found")

func main() {
	err := Work()
	if errors.Is(err, UserError) {
		fmt.Println("error matched")
	} else {
		fmt.Println("error not matched")
	}
}

func Work() error {
	return UserError
}
```

---

## F16 — Use defer mu.Unlock() immediately after mu.Lock() in a function

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	Work(&mu)
}

func Work(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	return
	fmt.Println("work done")
}
```

---

## F15 — Use defer to ensure a file (or fake closer) is always closed

```go
package main

import (
	"fmt"
)

type User struct {
	Name string
}

func (u User) Close() {
	fmt.Println("closed")
}

func main() {
	var user User
	defer user.Close()

	user.Name = "shivam"
	return

	fmt.Println("User Name: ", user.Name)
}
```

---

## F14 — Use context.WithValue to attach a request-ID string, read it inside a func

```go
package main

import (
	"context"
	"fmt"
)

type key string

const userId key = "userId"

func main() {
	ctx := context.WithValue(context.Background(), userId, "shivam")
	read(ctx)
}

func read(ctx context.Context) {
	value, ok := ctx.Value(userId).(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no ok")
	}
}
```

---

## F13 — Create a context with a 1-second timeout; show it expires

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg.Add(1)
	go inc(ctx, &wg)

	wg.Wait()
	fmt.Println(ctx.Err())
}

func inc(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("context done")
	}
}
```

---

## F12 — Create a context with cancel, pass it to a goroutine, call cancel()

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go inc(ctx, &wg)

	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
}

func inc(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("context done")
	}
}
```

---

## F11 — Increment a counter from 100 goroutines using atomic.AddInt64

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go inc(&counter, 1, &wg)
	}

	wg.Wait()
	fmt.Println(counter)
}

func inc(counter *int64, new int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(counter, new)
}
```

## F19 — Create a custom error type with an extra Code field; implement error interface

```go
package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Code string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field: %s", e.Code)
}

func validate(code string) error {
	original := &ValidationError{Code: code}
	return fmt.Errorf("request invalid: %w", original)
}

func main() {
	err := validate("email")

	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("Caught ValidationError, field:", ve.Code)
	}
}
```

## F20 — Use recover() inside a deferred function to catch a panic

```go
package main

import (
	"fmt"
)

func main() {
	err := work()
	if err != nil {
		fmt.Println(err)
	}
}

func work() (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	panic("panic occurred")
	fmt.Println("Task Scuccess")
	return nil
}
```
