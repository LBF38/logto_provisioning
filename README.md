# Logto Provisioning

This is a POC for enabling bulk management and provisioning to [Logto](https://logto.io).

This simple project is written using Go and produces a simple binary for provisioning config to your Logto instance.

## Objectives

As proposed in [this GitHub issue](https://github.com/logto-io/logto/issues/5966), it is difficult to make bulk changes in Logto, on RBAC, organizations or applications.

Therefore, this tool aims to provide a simple way to fix this:

- [ ] Provision RBAC roles
- [ ] Provision applications
- [ ] Provision organizations

The tool can provision to any Logto instance (Cloud or self-hosted) via [the Management API](https://logto.io/products/management-api).

## Ideas

The first and simplest idea is to define a JSON schema for the YAML configuration file to provision.
Then, users can define their resources, RBAC roles and organizations based on this schema and run the tool against the file to provision data to Logto instance.

Here is a first draft on the configuration file schema:

```yaml
logto:
    url: https://[tenant-id].logto.app/api # your logto instance. Management API endpoint
    appID: APP_ID # app ID configured for the tool to access M2M API => Better defined in env vars.
    appSecret: APP_SECRET # same as APP_ID. => Better defined in env vars.
resources:
  - baseUrl: https://api.store.io
    /orders:
      - read:order
      - write:order
      - delete:order
    /products:
      - read:product
      - write:product
      - delete:product
roles:
  order_admin:
    # the resources/permissions attributions are autocompleted from above definition in resources.
    https://api.store.io/orders: 
      - read:order 
      - write:order
      - delete:order
    https://api.store.io/products:
      - read:product
  product_admin:
    https://api.store.io/products:
      - read:product
      - write:product
      - delete:product
users: # WIP
    - username: john_doe
      password: pwd
organizations: # WIP
    acme_inc:
        - john_doe
```
