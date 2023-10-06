package result

import (
	"errors"
	"testing"
)

func TestWithRecover(t *testing.T) {
	type args struct {
		fn func() Result[struct{}]
	}
	tests := []struct {
		name    string
		args    args
		wantRet error
	}{
		{
			name:    "should ok",
			args:    args{func() Result[struct{}] { return Ok(struct{}{}) }},
			wantRet: nil,
		},
		{
			name:    "should error",
			args:    args{func() Result[struct{}] { panic(errAlwaysFailed) }},
			wantRet: errAlwaysFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := WithRecover(tt.args.fn).UnwrapErrUnchecked(); !errors.Is(gotRet, tt.wantRet) {
				t.Errorf("WithRecover() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
