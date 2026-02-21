package plugin

import (
	"golang.org/x/tools/go/analysis"
	"github.com/golangci/plugin-module-register/register"
	

	logcheck "github.com/srKazuya/loglint/internal"
)

func init() {
	register.Plugin("loglint", New)
}

type plugin struct{}
var _ register.LinterPlugin = (*plugin)(nil)


func New(settings any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

func (*plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return logcheck.AllAnalyzers(), nil
}


func (*plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo 
}