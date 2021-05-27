package kubebench

import (
	"embed"
)

//go:embed jobs/*.yaml
var yamlDir embed.FS

func embedYAMLs(kubebenchYAML string) ([]byte, error) {

	var (
		data []byte
		err  error
	)
	data, err = yamlDir.ReadFile("jobs/" + kubebenchYAML)
	if err != nil {
		return nil, err
	}
	return data, nil
}
