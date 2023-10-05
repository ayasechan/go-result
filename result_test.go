package result

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ExampleResult_downloadFile() {
	download := func(url string, dst io.Writer) (ret *Result[struct{}]) {
		return withRecover(func() *Result[struct{}] {
			req := AsResult(http.NewRequest(http.MethodGet, url, nil)).Unwrap()
			resp := AsResult(http.DefaultClient.Do(req)).Unwrap()
			defer resp.Body.Close()

			AsResult(io.Copy(dst, resp.Body)).Unwrap()
			return Ok(struct{}{})
		})
	}

	dstFd := AsResult(os.CreateTemp("", ".temp")).Unwrap()
	defer dstFd.Close()
	defer os.Remove(dstFd.Name())

	download("https://httpbin.org/", dstFd).InspectErr(func(err error) {
		fmt.Printf("%+v", err)
		os.Exit(1)
	})
}

func withRecover[T any](fn func() *Result[T]) (ret *Result[T]) {
	defer func() {
		err := recover()
		if err != nil {
			ret = Err[T](err.(error))
		}
	}()
	ret = fn()
	return
}
