# AI-Summary
## Directory Summary
This directory contains scripts and templates for managing Azure resources within the mygreeterv3 project. It includes a Bash script for converting JSON to YAML configurations and a Bicep template for deploying Azure resources at the subscription level.

**Tags:** Azure, deployment, Bash script, resource management, template

## File Details
    
### /mygreeterv3/server/resources/azuresdk/saveOutputs.sh
This is a Bash script named 'saveOutputs.sh' that reads a client ID from an input JSON file and generates a YAML configuration file with this client ID. The script requires two arguments: the input JSON file and the output YAML file. It checks if the correct number of arguments is provided and if the input file exists before proceeding.

### /mygreeterv3/server/resources/azuresdk/Azuresdk.ServiceResources.Template.bicep
This Bicep template file defines the deployment configuration for Azure resources at the subscription level. It includes parameters for resource names, subscription ID, location, and resource group name. The template references existing resources such as resource groups and service bus namespaces, and defines modules for AKS clusters, managed identities, and role assignments. The template outputs the client ID of the managed identity.
