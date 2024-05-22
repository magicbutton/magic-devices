<#---
title: Job deploy to production
tag: jobdeployproduction
api: post
---
#>

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions") {
  $path = join-path $PSScriptRoot ".." ".."
}
else {
  $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path

$inputFile = join-path  $koksmatDir "koksmat.json"

if (!(Test-Path -Path $inputFile) ) {
  Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"
$appname = $json.appname
$imagename = $json.imagename

<#
Envs
#>

$envs = @()
function env($name, $value ) {
  if ($null -eq $value) {
    throw "Environment value for $name is not set"
  }
  return @{name = $name; value = $value }
}



$envs += env "PNPAPPID" $env:PNPAPPID
$envs += env "PNPTENANTID" $env:PNPTENANTID
$envs += env "PNPCERTIFICATE" $env:PNPCERTIFICATE
$envs += env "PNPSITE" $env:PNPSITE
$envs += env "SITEURL" $env:SITEURL
$envs += env "NATS" "nats://nats:4222"
$envs += env "POSTGRES_DB" $env:POSTGRES_DB
$envs += env "DEVICEKPIBLOB" $env:DEVICEKPIBLOB
$configEnv = ""
foreach ($item in $envs) {

  $configEnv += @"
              - name: $($item.name)
                value: $($item.value)

"@
}

<#
Then we build the deployment file
#>

$image = "$($imagename)-app:$($version)"

$config = @"
apiVersion: batch/v1
kind: CronJob
metadata:
  name: $appname-job
spec:
  schedule: '* 3 * * *'
  concurrencyPolicy: Allow
  suspend: false
  jobTemplate:
    metadata:
      creationTimestamp: null
    spec:
      parallelism: 1
      completions: 1
      backoffLimit: 3
      template:
        metadata:
          creationTimestamp: null
          labels:
            app: $appname-job
        spec:

          containers:
            - name: $appname-app
              image: $image
              command: [$appname]
              args: ["run","devicekpi"]               
              env:
$configEnv                           
              
              
          restartPolicy: Never
          terminationGracePeriodSeconds: 30
          dnsPolicy: ClusterFirst
          securityContext: {}
          schedulerName: default-scheduler
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1
                  

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -