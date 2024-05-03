<#---
title: Import Excel

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


$excelfilename = join-path $workdir "Estrazione Catalogo NEAR_20240404.xlsx"
$sheetname = "Applicazioni v2"
$namespace = "excel"
$tablename = "applications"

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

