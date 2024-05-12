package config

import (
	"encoding/json"
	"fmt"

	"github.com/scch94/Gconfiguration"
)

type MicropagosDatabaseConfiguration struct {
	nombre string `json:"nombre"`
}

func (m MicropagosDatabaseConfiguration) ConfigurationString() string {
	configJSON, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("Error al convertir la configuración a JSON: %v", err)
	}
	return string(configJSON)
}

var Config MicropagosDatabaseConfiguration

func Upconfig() error {
	fmt.Println("starting to get the config struct ")
	err := Gconfiguration.GetConfig(&Config)
	if err != nil {
		fmt.Println("error in Gconfiguration.GetConfig() ", err)
		return err
	}
	return nil
}
