package orm

import (
	"fmt"
	"runtime/debug"

	imodels "github.com/nbcx/go-orm/orm/internal/models"
)

var defaultModelCache = imodels.NewModelCacheHandler()

// RegisterModel Register models
func RegisterModel(models ...interface{}) {
	RegisterModelWithPrefix("", models...)
}

// RegisterModelWithPrefix Register models with a prefix
func RegisterModelWithPrefix(prefix string, models ...interface{}) {
	if err := defaultModelCache.Register(prefix, true, models...); err != nil {
		panic(err)
	}
}

// RegisterModelWithSuffix Register models with a suffix
func RegisterModelWithSuffix(suffix string, models ...interface{}) {
	if err := defaultModelCache.Register(suffix, false, models...); err != nil {
		panic(err)
	}
}

// BootStrap Bootstrap models.
// make All model parsed and can not add more models
func BootStrap() {
	if dataBaseCache.getDefault() == nil {
		fmt.Println("must have one Register DataBase alias named `default`")
		debug.PrintStack()
		return
	}
	defaultModelCache.Bootstrap()
}

// ResetModelCache Clean model cache. Then you can re-RegisterModel.
// Common use this api for test case.
func ResetModelCache() {
	defaultModelCache.Clean()
}
