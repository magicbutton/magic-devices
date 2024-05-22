<#---
title: Load Transport Statistics into SQL
input: InboundTransportRuleHits.json
tag: load-exchangestats
api: post
---

#>


if ($null -eq $env:WORKDIR ) {
   $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
   New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

$inputfile = join-path $workdir "InboundTransportRuleHits.json"

$env:POSTGRES_DB = "postgres://ps_user:SecurePassword@localhost:5432/devices?sslmode=disable"
magic-files import $inputfile