Import-Module -Name powershell-yaml
New-Item -ItemType Directory apps
$apps = $(kubectl get -o jsonpath='{range .items[*]}{.spec.project}|{.metadata.name}|{.metadata.namespace}|{.spec.destination}{"\n"}{end}' -A applications.argoproj.io)

$apps | ForEach-Object {
    $project = $_.Split('|')[0]
    $appName = $_.Split('|')[1]
    $namespace = $_.Split('|')[2]
    $destination = $_.Split('|')[3] | ConvertFrom-Json
    New-Item -ItemType File -Path "apps" -Name "$namespace-$appName.yaml" -Force
    $app = [PSCustomObject]@{
      apiVersion = "applications.argocd.crossplane.io/v1alpha1"
      kind = "Application"
      metadata = @{
        name = "$namespace-$appName"
        annotations = @{
          "crossplane.io/external-name" = $appName
        }
      }
      spec = @{
        forProvider = @{
          project = $project
          destination = $destination
          appNamespace = $namespace
        }
        providerConfigRef = @{
          name = "eus2-preprod"
        }
        managementPolicies = @("Observe")
      }
    }
    Set-Content -Path "apps/$namespace-$appName.yaml" -Value $($app | ConvertTo-Yaml)
}
