//go:build !windows && !plan9
// +build !windows,!plan9

package sftp

import (
	"path"
	"strings"
)

func (s *Server) toLocalPath(p string) string {
	if s.workDir != "" {
		if !path.IsAbs(p) {
			p = path.Join(s.workDir, p)
		} else {
			if !strings.HasPrefix(p, s.workDir) {
				p = ""
			}
		}
	}

	return p
}
