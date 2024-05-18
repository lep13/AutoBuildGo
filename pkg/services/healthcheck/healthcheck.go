package healthcheck

import "github.com/lep13/AutoBuildGo/pkg/model"

func GetHealthStatus() *model.Health {
	health := new(model.Health)
	health.Status = "UP"
	return health
}
