package middleware

import (
	"net/http"

	"github.com/dhuki/rest-template/common"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/juju/ratelimit"
)

func TokenBucketLimiter(bucket *ratelimit.Bucket, logger log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if bucket.TakeAvailable(1) == 0 {
				ErrorEncoder(ctx, common.ErrLimitExceed, w)
				ErrorHandlerFunc(logger)(ctx, common.ErrLimitExceed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
