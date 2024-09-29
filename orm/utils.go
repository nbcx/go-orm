package orm

import (
	"github.com/nbcx/go-orm/orm/internal/models"
	"github.com/nbcx/go-orm/orm/internal/utils"
)

type StrTo = utils.StrTo

func SetNameStrategy(s string) {
	if models.SnakeAcronymNameStrategy != s {
		models.NameStrategy = models.DefaultNameStrategy
	}
	models.NameStrategy = s
}
