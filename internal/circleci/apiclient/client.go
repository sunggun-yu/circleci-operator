package apiclient

import (
	"errors"
	"net/http"

	cciapi "github.com/CircleCI-Public/circleci-cli/api"
	ccigraphql "github.com/CircleCI-Public/circleci-cli/api/graphql"
	cciapiproject "github.com/CircleCI-Public/circleci-cli/api/project"
	ccisettings "github.com/CircleCI-Public/circleci-cli/settings"
	"github.com/go-logr/logr"
	v1 "github.com/sunggun-yu/circleci-operator/api/circleci/v1"
	"github.com/sunggun-yu/circleci-operator/internal/utils"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type ApiClient struct {
	config *ccisettings.Config
	log    logr.Logger
}

// validateConfig validates the configuration by calling the whoami endpoint and returns the owner of the token.
func (c *ApiClient) validateConfig() error {
	gclient := ccigraphql.NewClient(c.config.HTTPClient, c.config.Host, c.config.Endpoint, c.config.Token, c.config.Debug)
	resp, err := cciapi.WhoamiQuery(gclient)
	if err != nil || resp == nil || resp.Me.Name == "" {
		c.log.Error(err, "API token is not valid")
		return errors.New("API token is not valid")
	}
	c.log.V(10).Info("CircleCI Token is valid.", "whoaami", resp)
	return nil
}

func NewCircleciAPIClient(log logr.Logger, k8sclient k8sclient.Client, spec *v1.ProviderSpec) (*ApiClient, error) {

	token, err := utils.GetSecretValue(k8sclient, spec.Auth.Token.SecretKeyRef)
	if err != nil {
		return nil, err
	}

	config := &ccisettings.Config{
		Debug:        false,
		Token:        token,
		Host:         spec.Host,
		RestEndpoint: spec.RestEndpoint,
		Endpoint:     spec.Endpoint,
		HTTPClient:   &http.Client{},
	}

	c := &ApiClient{
		config: config,
		log:    log.WithName("apiclient"),
	}

	err = c.validateConfig()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// NewContextClient creates a new context REST API client
func (c *ApiClient) NewContextClient() (*cciapi.ContextRestClient, error) {
	return cciapi.NewContextRestClient(*c.config)
}

// NewProjectClient creates a new project REST API client
func (c *ApiClient) NewProjectClient() (cciapiproject.ProjectClient, error) {
	return cciapiproject.NewProjectRestClient(*c.config)
}
