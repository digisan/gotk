package queue

import (
	"fmt"
	"strings"
)

type (
	sBuilder = strings.Builder
)

var (
	fSf        = fmt.Sprintf
	fPln       = fmt.Println
	sTrimRight = strings.TrimRight
)
