package result

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

var errAlwaysFailed = errors.New("always failed")

func ExampleResult_downloadFile() {
	download := func(url string, dst io.Writer) Result[struct{}] {
		return WithRecover(func() Result[struct{}] {
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

func TestAsResult(t *testing.T) {
	type args struct {
		value struct{}
		err   error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "should ok",
			args: args{value: struct{}{}, err: nil},
			want: nil,
		},
		{
			name: "should error",
			args: args{value: struct{}{}, err: errAlwaysFailed},
			want: errAlwaysFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsResult(tt.args.value, tt.args.err).UnwrapErrUnchecked(); !errors.Is(got, tt.want) {
				t.Errorf("AsResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
