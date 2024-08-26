package internal

import (
	"fmt"

	"github.com/mach-composer/mach-composer-plugin-helpers/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

type BluestonePimPlugin struct {
	provider     string
	environment  string
	globalConfig *BluestonePimGlobalConfig
	siteConfigs  map[string]*BluestonePimSiteConfig
	enabled      bool
}

func NewBluestonePimPlugin() schema.MachComposerPlugin {
	state := &BluestonePimPlugin{
		provider:    "0.1.0",
		siteConfigs: map[string]*BluestonePimSiteConfig{},
	}
	return plugin.NewPlugin(&schema.PluginSchema{
		Identifier:          "bluestonepim",
		Configure:           state.Configure,
		IsEnabled:           func() bool { return state.enabled },
		GetValidationSchema: state.GetValidationSchema,

		SetGlobalConfig: state.SetGlobalConfig,
		SetSiteConfig:   state.SetSiteConfig,

		// Renders
		RenderTerraformProviders: state.RenderTerraformProviders,
		RenderTerraformResources: state.RenderTerraformResources,
		RenderTerraformComponent: state.RenderTerraformComponent,
	})
}

func (p *BluestonePimPlugin) Configure(environment string, provider string) error {
	p.environment = environment
	if provider != "" {
		p.provider = provider
	}
	return nil
}

func (p *BluestonePimPlugin) GetValidationSchema() (*schema.ValidationSchema, error) {
	result := getSchema()
	return result, nil
}

func (p *BluestonePimPlugin) SetGlobalConfig(data map[string]any) error {
	cfg := BluestonePimGlobalConfig{}

	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.globalConfig = &cfg
	p.enabled = true

	return nil
}

func (p *BluestonePimPlugin) SetSiteConfig(site string, data map[string]any) error {
	cfg := BluestonePimSiteConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.siteConfigs[site] = &cfg
	p.enabled = true
	return nil
}

func (p *BluestonePimPlugin) getSiteConfig(site string) *BluestonePimSiteConfig {
	result := &BluestonePimSiteConfig{}
	if p.globalConfig != nil {
		result.ClientID = p.globalConfig.ClientID
		result.ClientSecret = p.globalConfig.ClientSecret
		result.AuthUrl = p.globalConfig.AuthUrl
		result.ApiUrl = p.globalConfig.ApiUrl
	}

	cfg, ok := p.siteConfigs[site]
	if ok {
		if cfg.ClientID != "" {
			result.ClientID = cfg.ClientID
		}
		if cfg.ClientSecret != "" {
			result.ClientSecret = cfg.ClientSecret
		}
		if cfg.AuthUrl != nil {
			result.AuthUrl = cfg.AuthUrl
		}
		if cfg.ApiUrl != nil {
			result.ApiUrl = cfg.ApiUrl
		}
	}

	return result
}

func (p *BluestonePimPlugin) RenderTerraformStateBackend(_ string) (string, error) {
	return "", nil
}

func (p *BluestonePimPlugin) RenderTerraformProviders(site string) (string, error) {
	cfg := p.getSiteConfig(site)

	if cfg == nil {
		return "", nil
	}

	result := fmt.Sprintf(`
		bluestonepim = {
			source = "labd/bluestonepim"
			version = "%s"
		}
	`, helpers.VersionConstraint(p.provider))

	return result, nil
}

func (p *BluestonePimPlugin) RenderTerraformResources(site string) (string, error) {
	cfg := p.getSiteConfig(site)

	if cfg == nil {
		return "", nil
	}

	template := `
		provider "bluestonepim" {
			{{ renderProperty "client_id" .ClientID }}
			{{ renderProperty "client_secret" .ClientSecret }}
			{{if .AuthUrl}}
			{{ renderProperty "auth_url" .AuthUrl }}
			{{end}}
			{{if .ApiUrl}}
			{{ renderProperty "api_url" .ApiUrl }}
			{{end}}
		}
	`
	return helpers.RenderGoTemplate(template, cfg)
}

func (p *BluestonePimPlugin) RenderTerraformComponent(_ string, _ string) (*schema.ComponentSchema, error) {
	result := &schema.ComponentSchema{
		Providers: []string{"bluestonepim = bluestonepim"},
	}

	return result, nil
}
