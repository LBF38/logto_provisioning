package internal

import (
	"context"
	"fmt"
	"log"

	openapi "github.com/lbf38/logto_provisioning/logtoClient/client"
)

// Create OpenAPI instance
var (
	OpenapiCfg    = openapi.NewConfiguration()
	client = openapi.NewAPIClient(OpenapiCfg)
)

func CreateResource(ctx context.Context, name, indicator string) (*openapi.ListResources200ResponseInner, error) {
	req := *openapi.NewCreateResourceRequest(name, indicator)
	resp, r, err := client.ResourcesAPI.CreateResource(ctx).CreateResourceRequest(req).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `ResourcesAPI.CreateResource``: %v\nFull HTTP response: %v", err, r)
	}
	// response from `CreateResource`: ListResources200ResponseInner
	return resp, nil
}

func CreateResourceScope(ctx context.Context, scope, description string, createdResource *openapi.ListResources200ResponseInner) (*openapi.ListResources200ResponseInnerScopesInner, error) {
	req := openapi.NewCreateResourceScopeRequest(scope)
	req.SetDescription(description)
	resp, r, err := client.ResourcesAPI.CreateResourceScope(ctx, createdResource.Id).CreateResourceScopeRequest(*req).Execute()
	if err != nil {
		return nil, fmt.Errorf("error with Resources scope provisioning: %v\nFull HTTP response: %v", err, r)
	}
	return resp, nil
}

func CreateRole(ctx context.Context, role_name string, scopes []string) (*openapi.ListApplicationRoles200ResponseInner, error) {
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
func CleanLogtoInstance(ctx context.Context) error {

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
