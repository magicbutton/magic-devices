<#---
title: Import info SQL
tag: import-sql
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

$workdir = (Resolve-Path $workdir).Path

$sqldir = join-path $workdir "sqlimport"

Push-Location
Set-Location $psscriptroot

Get-ChildItem -Path  $sqldir  -Filter "*.createtablesql.sql" | ForEach-Object {
    $sql = get-content $_.FullName  -Raw
    
    # Write-Host   magic-devices sql exec $sql
    $rowsAffected = magic-devices sql exec $sql
    write-host $rowsAffected "rows affected executing" $_.Name

  
}
Get-ChildItem -Path  $sqldir  -Filter "*.inserttablesql*.sql" | ForEach-Object {
    $sql = get-content $_.FullName  -Raw
    
    # Write-Host   magic-devices sql exec $sql
    $rowsAffected = magic-devices sql exec $sql
    write-host $rowsAffected "rows affected executing" $_.Name

  
}
# Pop-Location