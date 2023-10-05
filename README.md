

# go-result

Result and Option for golang

# Install

```
go get github.com/ayasechan/go-result
```

# Example

```golang
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
```