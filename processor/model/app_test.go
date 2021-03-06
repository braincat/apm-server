package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
)

func TestAppTransformDefinition(t *testing.T) {
	myfn := func(fn TransformApp) string { return "ok" }
	res := myfn((*App).Transform)
	assert.Equal(t, "ok", res)
}

func TestAppTransform(t *testing.T) {

	version := "5.1.3"
	pid := 1234
	processTitle := "node"
	langName := "ecmascript"
	langVersion := "8"
	rtName := "node"
	rtVersion := "8.0.0"
	fwName := "Express"
	fwVersion := "1.2.3"
	agentName := "elastic-node"
	agentVersion := "1.0.0"
	tests := []struct {
		App    App
		Output common.MapStr
	}{
		{
			App: App{},
			Output: common.MapStr{
				"agent": common.MapStr{
					"name":    "",
					"version": "",
				},
				"name": "",
			},
		},
		{
			App: App{
				Name:         "myapp",
				Version:      &version,
				Pid:          &pid,
				ProcessTitle: &processTitle,
				Argv: []string{
					"node",
					"server.js",
				},
				Language: Language{
					Name:    &langName,
					Version: &langVersion,
				},
				Runtime: Runtime{
					Name:    &rtName,
					Version: &rtVersion,
				},
				Framework: Framework{
					Name:    &fwName,
					Version: &fwVersion,
				},
				Agent: Agent{
					Name:    agentName,
					Version: agentVersion,
				},
			},
			Output: common.MapStr{
				"name":          "myapp",
				"version":       "5.1.3",
				"pid":           1234,
				"process_title": "node",
				"argv": []string{
					"node",
					"server.js",
				},
				"language": common.MapStr{
					"name":    "ecmascript",
					"version": "8",
				},
				"runtime": common.MapStr{
					"name":    "node",
					"version": "8.0.0",
				},
				"framework": common.MapStr{
					"name":    "Express",
					"version": "1.2.3",
				},
				"agent": common.MapStr{
					"name":    "elastic-node",
					"version": "1.0.0",
				},
			},
		},
	}

	for _, test := range tests {
		output := test.App.Transform()
		assert.Equal(t, test.Output, output)
	}

}
