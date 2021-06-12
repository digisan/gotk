package filedir

import (
	"strings"
	"sync"
)

var (
	sJoin       = strings.Join
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sReplaceAll = strings.ReplaceAll
	mtx4crtdir  = &sync.Mutex{}
)
