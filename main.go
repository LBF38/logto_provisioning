package main

import (
	"context"
	"fmt"
	"log"

	openapi "github.com/lbf38/logto_provisioning/logtoClient/client"
	"golang.org/x/oauth2"
)

// Create OpenAPI instance
var (
	cfg    = openapi.NewConfiguration()
	client = openapi.NewAPIClient(cfg)
)

func main() {
	fmt.Println("Logto Provisioning")

	// Load provisioning config from file
	config, err := NewConfig("provisioning/logto_provisioning_dev.yaml")

	if err != nil {
		log.Fatalf("Failed loading config file: %v\n", err)
	}

	// DONE: handle the logto URL defined in the provisioning file.
	// => it should be replaced in the OpenAPI servers config.
	if config.Logto.Url != "" {
		cfg.Servers = openapi.ServerConfigurations{
			{
				URL:         config.Logto.Url,
				Description: "Logto instance URL",
			},
		}
	}

	// Fetch access token for interacting w/ Management API
	accessTokenResponse, err := config.FetchAccessTokenResponse()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}
	log.Printf("Access Token response: %v\n", accessTokenResponse)

	fmt.Printf("OpenAPI configuration: %+v\n", *cfg)
	fmt.Printf("OpenAPI client: %+v\n", *client)

	// OpenAPI OAuth2
	token := oauth2.Token{
		AccessToken: accessTokenResponse.AccessToken,
	}
	tokenSource := oauth2.StaticTokenSource(&token)
	auth := context.WithValue(context.Background(), openapi.ContextOAuth2, tokenSource)

	//! Provision Logto instance
	// Resources provisioning

	var resourceScopesMap = make(map[string]map[string]string)
	for _, resource := range config.Resources {
		for endpoint, scopes := range resource.Endpoints {
			indicator := resource.BaseUrl + endpoint
			createdResource, err := createResource(auth, endpoint, indicator)
			if err != nil {
				log.Fatalf("Resources provisioning: %v", err)
			}
			resourceScopesMap[indicator] = make(map[string]string)
			for _, scope := range scopes {
				createdScope, err := createResourceScope(auth, scope, "description", createdResource)
				if err != nil {
					log.Fatalf("Resources provisioning: %v", err)
				}
				fmt.Println(createdScope)
				resourceScopesMap[indicator][scope] = createdScope.Id
			}
		}
	}

	// Roles provisioning
	for _, role := range config.Roles {
		for role_name, resources := range role {
			fmt.Println(resources)
			// DONE: link each role w/ resources scopes
			scopeIds := []string{}
			for indicator, scopes := range resources {
				for _, scope := range scopes {
					scopeIds = append(scopeIds, resourceScopesMap[indicator][scope])
				}
			}
			fmt.Println("scopeIds: ", scopeIds)

			createdRole, err := createRole(auth, role_name, scopeIds)
			if err != nil {
				log.Fatalf("Roles provisioning: %v", err)
			}
			fmt.Println(createdRole)
		}
	}

	// Users provisioning

	// Organizations provisioning
}

func createResource(ctx context.Context, name, indicator string) (*openapi.ListResources200ResponseInner, error) {
	req := *openapi.NewCreateResourceRequest(name, indicator)
	resp, r, err := client.ResourcesAPI.CreateResource(ctx).CreateResourceRequest(req).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `ResourcesAPI.CreateResource``: %v\nFull HTTP response: %v", err, r)
	}
	// response from `CreateResource`: ListResources200ResponseInner
	return resp, nil
}

func createResourceScope(ctx context.Context, scope, description string, createdResource *openapi.ListResources200ResponseInner) (*openapi.ListResources200ResponseInnerScopesInner, error) {
	req := openapi.NewCreateResourceScopeRequest(scope)
	req.SetDescription(description)
	resp, r, err := client.ResourcesAPI.CreateResourceScope(ctx, createdResource.Id).CreateResourceScopeRequest(*req).Execute()
	if err != nil {
		return nil, fmt.Errorf("error with Resources scope provisioning: %v\nFull HTTP response: %v", err, r)
	}
	return resp, nil
}

func createRole(ctx context.Context, role_name string, scopes []string) (*openapi.ListApplicationRoles200ResponseInner, error) {
	req := openapi.NewCreateRoleRequest(role_name, "some description")
	if scopes != nil {
		req.SetScopeIds(scopes)
	}
	req.SetIsDefault(false)
	log.Printf("createRole request: %+v\n", req)

	resp, r, err := client.RolesAPI.CreateRole(ctx).CreateRoleRequest(*req).Execute()
	if r.StatusCode < 300 {
		return resp, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error with Roles provisioning: %v\nFull HTTP response: %v", err, r)
	}

	return resp, nil
}

// Remove all created objects in Logto instance
// such as API resources, Roles, Users and Organizations.
func cleanLogtoInstance(ctx context.Context) error {

	resourcesResp, _, err := client.ResourcesAPI.ListResources(ctx).Execute()
	if err != nil {
		return fmt.Errorf("Resources cleaning error: %v", err)
	}

	for _, resource := range resourcesResp {
		if resource.Id == "management-api" {
			// This is the default Logto Management API
			// We can't delete it as it's locked.
			continue
		}
		_, err = client.ResourcesAPI.DeleteResource(ctx, resource.Id).Execute()
		if err != nil {
			return fmt.Errorf("Resources cleaning error: %v", err)
		}
	}

	rolesResp, r, err := client.RolesAPI.ListRoles(ctx).Execute()
	if err != nil && r.StatusCode >= 300 {
		return fmt.Errorf("Roles cleaning error: %v", err)
	}

	for _, role := range rolesResp {
		_, err = client.RolesAPI.DeleteRole(ctx, role.Id).Execute()
	}
	if err != nil {
		return fmt.Errorf("Roles cleaning error: %v", err)
	}

	usersResp, _, err := client.UsersAPI.ListUsers(ctx).Execute()
	if err != nil {
		return fmt.Errorf("Users cleaning error: %v", err)
	}

	for _, user := range usersResp {
		_, err = client.UsersAPI.DeleteUser(ctx, user.Id).Execute()
		if err != nil {
			return fmt.Errorf("Users cleaning error: %v", err)
		}
	}

	orgsResp, _, err := client.OrganizationsAPI.ListOrganizations(ctx).Execute()
	if err != nil {
		return fmt.Errorf("Users cleaning error: %v", err)
	}

	for _, org := range orgsResp {
		_, err = client.OrganizationsAPI.DeleteOrganization(ctx, org.Id).Execute()
		if err != nil {
			return fmt.Errorf("Users cleaning error: %v", err)
		}
	}

	return nil

}
