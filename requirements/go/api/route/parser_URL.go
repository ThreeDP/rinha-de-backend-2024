package route

import (
	"strings"
)

func (s *BankServer) splitFields(str string, step string) []string {
	parts := strings.Split(str, step)
	var res []string
	for _, part := range parts {
		if len(part) > 0 {
			res = append(res, part)
		}
	}
	return res
}

func (s *BankServer) parseURLPath(uri, path string) (map[string]string, bool) {
	uriKeys := map[string]string{}
	sUri, sPath := s.splitFields(uri, "/"), s.splitFields(path, "/")
	if len(sUri) == len(sPath) {
		for i, u := range sUri {
			key, ok := strings.CutPrefix(u, ":")
			field := u
			if ok {
				_, ok := uriKeys[key]
				if !ok {
					uriKeys[key] = sPath[i]
				}
				field = uriKeys[key]
			}
			if field != sPath[i] {
				return nil, false
			}
		}
		return uriKeys, true
	}
	return nil, false
}