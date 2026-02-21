package plugin

import (

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	logcheck "github.com/srKazuya/loglint/internal"
	"github.com/srKazuya/loglint/internal/config"
)

func init() {
	register.Plugin("loglint", New)
}

type LoglintPlugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	if err := config.LoadFromAny(settings); err != nil {
		return nil, err
	}

	return &LoglintPlugin{}, nil
}

func (p *LoglintPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return logcheck.AllAnalyzers(), nil
}

func (p *LoglintPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}