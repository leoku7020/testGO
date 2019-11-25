package customeError

import (
    "fmt"
)

type MyError struct {
    Title string
    Message string
}

func (e MyError) Error() string {
    return fmt.Sprintf("%v: %v", e.Title, e.Message)
}

func main() {
    err := MyError{"Error Title 1", "Error Message 1"}
    fmt.Println(err)

    err = MyError{
        Title: "Error Title 2",
        Message: "Error Message 2",
    }
    fmt.Println(err)
}
