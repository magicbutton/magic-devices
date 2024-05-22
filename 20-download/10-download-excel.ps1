<#---
title: Download excel
tag: excel
api: post
connection: sharepoint
---

#>

param ($siteUrl = "https://christianiabpos.sharepoint.com/sites/devices", $href = "https://christianiabpos.sharepoint.com/sites/devices/Shared%20Documents/kpi/KPISICUREZZA_v1.xlsx")
if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = (Resolve-Path $workdir).Path
Connect-PnPOnline -Url $siteUrl   -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"


Push-Location
Set-Location $workdir
write-host "Downloading file to $workdir"
$destDir = $workdir
$path = split-path $href -Parent
$filename = split-path $href -Leaf
$filename = [System.Web.HttpUtility]::UrlDecode($filename )
$oldPref = $ErrorActionPreference
$ErrorActionPreference = "Continue"
Get-PnPFile -Url $href -Path $destDir -filename $filename -AsFile -Force -ErrorAction SilentlyContinue
$ErrorActionPreference = $oldPref

Pop-Location