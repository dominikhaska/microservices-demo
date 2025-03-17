package featureFlags

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/open-feature/go-sdk/openfeature"
)

// EnvVarProvider implements the FeatureProvider interface and reads flags from environment variables
type EnvVarProvider struct {
	// Optional prefix for all environment variable names
	Prefix string
}

// NewEnvVarProvider creates a new EnvVarProvider with an optional prefix
func NewEnvVarProvider(prefix string) *EnvVarProvider {
	return &EnvVarProvider{
		Prefix: prefix,
	}
}

// Metadata returns the metadata of the provider
func (e EnvVarProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{Name: "EnvVarProvider"}
}

// getEnvVarName formats the flag key as an environment variable name
func (e EnvVarProvider) getEnvVarName(flag string) string {
	// Convert to uppercase and replace non-alphanumeric with underscore
	envName := strings.ToUpper(flag)
	envName = strings.ReplaceAll(envName, "-", "_")
	envName = strings.ReplaceAll(envName, ".", "_")

	if e.Prefix != "" {
		envName = e.Prefix + "_" + envName
	}

	return envName
}

// BooleanEvaluation returns a boolean flag from environment or default
func (e EnvVarProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	envName := e.getEnvVarName(flag)
	envValue := os.Getenv(envName)

	if envValue != "" {
		if value, err := strconv.ParseBool(envValue); err == nil {
			return openfeature.BoolResolutionDetail{
				Value: value,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Variant: "env-var",
					Reason:  openfeature.TargetingMatchReason,
				},
			}
		}
	}

	return openfeature.BoolResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Variant: "default-variant",
			Reason:  openfeature.DefaultReason,
		},
	}
}

// StringEvaluation returns a string flag from environment or default
func (e EnvVarProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	envName := e.getEnvVarName(flag)
	envValue := os.Getenv(envName)

	if envValue != "" {
		return openfeature.StringResolutionDetail{
			Value: envValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				Variant: "env-var",
				Reason:  openfeature.TargetingMatchReason,
			},
		}
	}

	return openfeature.StringResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Variant: "default-variant",
			Reason:  openfeature.DefaultReason,
		},
	}
}

// FloatEvaluation returns a float flag from environment or default
func (e EnvVarProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	envName := e.getEnvVarName(flag)
	envValue := os.Getenv(envName)

	if envValue != "" {
		if value, err := strconv.ParseFloat(envValue, 64); err == nil {
			return openfeature.FloatResolutionDetail{
				Value: value,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Variant: "env-var",
					Reason:  openfeature.TargetingMatchReason,
				},
			}
		}
	}

	return openfeature.FloatResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Variant: "default-variant",
			Reason:  openfeature.DefaultReason,
		},
	}
}

// IntEvaluation returns an int flag from environment or default
func (e EnvVarProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	envName := e.getEnvVarName(flag)
	envValue := os.Getenv(envName)

	if envValue != "" {
		if value, err := strconv.ParseInt(envValue, 10, 64); err == nil {
			return openfeature.IntResolutionDetail{
				Value: value,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Variant: "env-var",
					Reason:  openfeature.TargetingMatchReason,
				},
			}
		}
	}

	return openfeature.IntResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Variant: "default-variant",
			Reason:  openfeature.DefaultReason,
		},
	}
}

// ObjectEvaluation returns an object flag (always default as env vars can't store objects)
func (e EnvVarProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Variant: "default-variant",
			Reason:  openfeature.DefaultReason,
		},
	}
}

// Hooks returns hooks
func (e EnvVarProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}
