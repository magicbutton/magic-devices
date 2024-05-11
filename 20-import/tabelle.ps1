<#---
title: Import Sheet Tabelle
tag: tabelle
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


$excelfilename = join-path $workdir "devices.xlsx"
$sheetname = "tabelle"
$namespace = "exceldevice"
$tablename = "tabelle"

write-host "Workdir: $workdir"

$exportdir = join-path $workdir "sqlimport"
if (-not (Test-Path $exportdir)) {
    New-Item -Path $exportdir -ItemType Directory | Out-Null
}

Push-Location
Set-Location $exportdir

try {
    magic-mix from excel to sql $excelfilename $sheetname $namespace $tablename 
}
catch {
    Write-Host "Error: $_" -ForegroundColor Red
}
finally {
    Pop-Location
}

