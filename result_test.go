package result

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ExampleResult_downloadFile() {
	download := func(url string, dst io.Writer) (ret *Result[struct{}]) {
		// You might want to wrap this part as a function
		defer func() {
			err := recover()
			if err != nil {
				ret = Err[struct{}](err.(error))
			}
		}()

		req := AsResult(http.NewRequest(http.MethodGet, url, nil)).Unwrap()
		resp := AsResult(http.DefaultClient.Do(req)).Unwrap()
		defer resp.Body.Close()

		AsResult(io.Copy(dst, resp.Body)).Unwrap()
		return Ok(struct{}{})
	}

	dstFd := AsResult(os.CreateTemp("", ".temp")).Unwrap()
	defer dstFd.Close()
	defer os.Remove(dstFd.Name())

	download("https://httpbin.org/", dstFd).InspectErr(func(err error) {
		fmt.Printf("%+v", err)
		os.Exit(1)
	})
}
