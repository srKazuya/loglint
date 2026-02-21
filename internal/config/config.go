package config

type RuleConfig struct {
    Enabled  *bool    `mapstructure:"enabled"`
    Severity string   `mapstructure:"severity"`
    Patterns []string `mapstructure:"patterns"`
}

type Config struct {
    Rules map[string]RuleConfig `mapstructure:"rules"`
}

var Global Config

func IsEnabled(name string, defaultValue bool) bool {
    rule, ok := Global.Rules[name]
    if !ok || rule.Enabled == nil {
        return defaultValue
    }
    return *rule.Enabled
}

func GetPatterns(name string, defaultPatterns []string) []string {
    rule, ok := Global.Rules[name]
    if !ok || len(rule.Patterns) == 0 {
        return defaultPatterns
    }
    return rule.Patterns
}

func GetSeverity(name string, defaultValue string) string {
	rule, ok := Global.Rules[name]
	if !ok || rule.Severity == "" {
		return defaultValue
	}
	return rule.Severity
}