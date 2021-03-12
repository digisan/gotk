package slice

import (
	"fmt"
	"strings"
)

var (
	fPln        = fmt.Println
	sJoin       = strings.Join
	sSplit      = strings.Split
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sTrimPrefix = strings.TrimPrefix
	sTrimSuffix = strings.TrimSuffix
	sReplaceAll = strings.ReplaceAll
	sContains   = strings.Contains
)
