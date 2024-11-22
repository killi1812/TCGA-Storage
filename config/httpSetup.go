package config

import (
	"fmt"
)

type Controller interface {
	RegisterEndpoints() error
}

func RegiserControllers(controllers []Controller) error {
	for _, controller := range controllers {
		err := controller.RegisterEndpoints()
		if err != nil {
			return fmt.Errorf("Error registering endpoints from controller \n%s", err.Error())
		}
	}
	return nil
}
