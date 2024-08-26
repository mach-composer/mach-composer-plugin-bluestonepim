# Bluestone Pim Plugin for MACH composer

This repository contains the Bluestone PIM plugin for Mach Composer. It requires MACH
composer >= 2.5.x

This plugin uses the [Bluestone PIM Terraform Provider](https://github.com/labd/terraform-provider-bluestonepim)

## Usage

```yaml
mach_composer:
  version: 1
  plugins:
    bluestonepim:
      source: mach-composer/bluestonepim
      version: 0.1.0

global:
  # ...k
  bluestonepim:
    client_id: "my-client-id"
    client_secret: "my-client-secret"

sites:
  - identifier: my-site
    # ...
    bluestonepim:
      client_id: "my-test-client-id"
      client_secret: "my-test-client-secret"
      auth_url: "https://idp.test.bluestonepim.com/op/token"
      api_url: "https://api.test.bluestonepim.com"
    components:
      - name: my-component
        # ...
```
