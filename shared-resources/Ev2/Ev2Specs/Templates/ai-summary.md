# AI-Summary
## Directory Summary
This directory contains templates for setting up Azure resources and role assignments in the Ev2 pipeline. It includes a Bicep template for configuring Azure Container Registry push operations and an ARM template for assigning roles at the subscription level.

**Tags:** Azure, templates, role assignment, Bicep, ARM

## File Details
    
### /shared-resources/Ev2/Ev2Specs/Templates/AcrPush.SharedResources.Template.bicep
This Bicep template is used in the Ev2 pipeline to configure resources for pushing images to an Azure Container Registry (ACR). It defines parameters for the resource name and location, and creates a managed identity and role assignment for image push operations in the ACR. The template is specific to Ev2 services and uses a user-assigned identity to manage access.

### /shared-resources/Ev2/Ev2Specs/Templates/RoleAssignment.Subscription.Template.json
This JSON file is an Azure Resource Manager (ARM) template for assigning roles at the subscription level. It defines parameters for a principal ID and a built-in role type, which can be one of several predefined Azure roles. The template uses these parameters to create a role assignment resource, specifying the role definition ID and the principal ID.
