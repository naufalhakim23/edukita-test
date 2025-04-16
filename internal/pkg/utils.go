package pkg

import (
	"edukita-teaching-grading/configs"
	"edukita-teaching-grading/docs"
)

func SwaggerInfo(config *configs.Config) {
	docs.SwaggerInfo.Version = config.Application.Name
	docs.SwaggerInfo.Host = config.Application.SwaggerPath
	docs.SwaggerInfo.BasePath = config.Application.SwaggerPath
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Title = config.Application.Name
	docs.SwaggerInfo.Description = `
		This is a sample server for teaching grading system
		- Author: hakim
		- Version: 1.0.0
	`
}
