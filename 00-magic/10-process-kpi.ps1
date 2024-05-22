<#---
title: Process KPI
tag: kpi
api: post
---

#>

if ($env:WORKDIR -eq $null) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    $x = New-Item -Path $workdir -ItemType Directory 
}

$workdir = Resolve-Path $workdir
<#
## Extract
### Download Excel

Download the Excel file from the [link] Handled my Power Automate flow that will store the file at the following location: [link]
#>
write-host "Workdir: $workdir"

write-host "Getting Data"
magic-devices download excel
magic-devices extract fullkpi
magic-devices extract intune
magic-devices extract tabelle
magic-devices extract import-sql


<#
### Convert to SQL
Run the following command to import the Excel file into the database:

The Excel file will be imported into the database.

## Transform
### Transform the data
4. Transform the data into KPIs.
5. The KPIs will be available in the database.

## Load
### Load the data
6. Load the data into the blob store
7. The KPIs will be available in the blob store.


#>