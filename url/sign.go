package url

import (
	"fmt"
	"net/url"

	"github.com/deoops-net/easy/hash"
)

// SignUrl use md5 to sign sorted kv form patched with key
// e.g.  md5(sort("foo=bar&bar=baz") + "&key=value")
func SignUrl(v url.Values, key, value string) (r string, err error) {
	s, err := url.QueryUnescape(v.Encode())
	if err != nil {
		return
	}

	f := fmt.Sprintf("%s&%s=%s", s, key, value)
	r = hash.MD5(f)

	return
}
