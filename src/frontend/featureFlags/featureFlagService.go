package featureFlags

import (
	"context"
	//"fmt"
	//"os"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/sirupsen/logrus"
)

type featureFlagService struct {
	client *openfeature.Client
	log    *logrus.Logger
}

func NewFeatureFlagService(ctx context.Context, log *logrus.Logger) (*featureFlagService, error) {
	// Register your feature flag provider
	err := openfeature.SetProvider(NewEnvVarProvider(""))
	if err != nil {
		return nil, err
	}

	// Create a named client for the frontend service
	client := openfeature.NewClient("frontend")

	log.Info("Feature flag service initialized successfully")

	return &featureFlagService{
		client: client,
		log:    log,
	}, nil
}

// Method to check boolean flags with logging
func (ffs *featureFlagService) GetBooleanFlag(ctx context.Context, flagKey string, defaultValue bool) bool {
	value, err := ffs.client.BooleanValue(ctx, flagKey, defaultValue, openfeature.EvaluationContext{})
	if err != nil {
		ffs.log.WithFields(logrus.Fields{
			"flagKey": flagKey,
			"default": defaultValue,
			"error":   err.Error(),
		}).Warn("Failed to evaluate boolean flag, using default")
		return defaultValue
	}

	ffs.log.WithFields(logrus.Fields{
		"flagKey": flagKey,
		"value":   value,
	}).Info("Feature flag evaluated")

	return value
}

// Method to check string flags with logging
func (ffs *featureFlagService) getStringFlag(ctx context.Context, flagKey string, defaultValue string) string {
	value, err := ffs.client.StringValue(ctx, flagKey, defaultValue, openfeature.EvaluationContext{})
	if err != nil {
		ffs.log.WithFields(logrus.Fields{
			"flagKey": flagKey,
			"default": defaultValue,
			"error":   err.Error(),
		}).Warn("Failed to evaluate string flag, using default")
		return defaultValue
	}

	ffs.log.WithFields(logrus.Fields{
		"flagKey": flagKey,
		"value":   value,
	}).Info("Feature flag evaluated")

	return value
}
