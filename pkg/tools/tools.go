//go:build tools

package tools

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/swaggo/swag/cmd/swag"
)
