package io

import (
	"strings"
	"sync"
)

var (
	sJoin       = strings.Join
	sReplaceAll = strings.ReplaceAll
	mtx4crtdir  = &sync.Mutex{}
)
