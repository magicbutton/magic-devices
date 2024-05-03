<#---
title: Get Hub Site Spokes Pages
tag: hubsite-spokes-pages
connection: sharepoint
api: post
output: hubsite-spokes-pages.json
---#>
param (

    [string]$HubSiteId = "b80f09f2-c5e5-4f69-9944-33e8fe18a96c"
)

if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"


<#

## Supporting functions
#>
function EnsurePath($path) {
    if (-not (Test-Path $path)) {
        write-host "Creating directory  $path"
        New-Item -Path  $path -ItemType Directory | Out-Null
    }

}

function WriteTextFile($filepath, $content) {
    if (Test-Path $filepath) {
        $existingcontent = Get-Content $filepath
        

        # find the string "keep: true" in the $content read from Get-Content
        $keep = $existingcontent | Select-String -Pattern "keep: true"
        if ($keep) {
            # if the string is found, remove it from the $content
            write-host "Keeping $filepath" -ForegroundColor Yellow
            return
        }
    }
    $content | Out-File -FilePath $filepath -Encoding utf8NoBOM
    if ($verbose) {
        write-host "Writing $filepath" -ForegroundColor Yellow
    
    }
}

$hubsite = Get-PnPHubSite -Identity $HubSiteId
$hubsiteName = $hubsite.Title

$reportRoot = Join-Path $workdir "hubsite-spokes-pages" 
EnsurePath $reportRoot
$reportHub = Join-Path $reportRoot "$hubsiteName ($HubSiteId)"
ensurePath $reportHub



$childSites = Get-PnPHubSiteChild -Identity $HubSiteId
$sites = @()
foreach ($childSite in $childSites) {
    Connect-PnPOnline -Url $childSite  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

    $site = Get-PnPSite -Includes RootWeb, ServerRelativeUrl
    
    $web = $site.RootWeb
    Write-Host "Site $($web.Title) " -NoNewline
    $reportFilePath = join-path $reportHub "$($web.ServerRelativeUrl -replace "/","-").json"
    $sites += $reportFilePath 
    # if (Test-Path $reportFilePath) {
    #     write-host " - Skipping" -ForegroundColor Yellow
    #     continue
    # }
 
    try {
        $SitePages = Get-PnPListItem -List "Site Pages" -Fields "Title", "FileRef", "Created_x0020_Date", "Last_x0020_Modified", "Editor" -PageSize 5000
        
    }
    catch {
        $SitePages = Get-PnPListItem -List "SitePages" -Fields "Title", "FileRef", "Created_x0020_Date", "Last_x0020_Modified", "Editor" -PageSize 5000
        <#Do this if a terminating exception happens#>
    }
    Write-Output "has $($SitePages.Count) pages"
    $PagesDataColl = @()
    #Collect Site Pages data - Title, URL and other properties
    ForEach ($Page in $SitePages) {
        $Data = @{
            HubSiteId   = $HubSiteId
            PageName    = $Page.FieldValues.Title
            RelativeURL = $Page.FieldValues.FileRef     
            CreatedOn   = $Page.FieldValues.Created_x0020_Date
            ModifiedOn  = $Page.FieldValues.Last_x0020_Modified
            Editor      = $Page.FieldValues.Editor.Email
            ID          = $Page.ID
        }
        $PagesDataColl += $Data
    }

  
    $siteData = @{
        siteurl     = $childSite 
        title       = $web.Title
        HubSiteId   = $HubSiteId
        WelcomePage = $web.WelcomePage
        pages       = $PagesDataColl
    }

    ConvertTo-Json  -InputObject $siteData -Depth 10 | Out-File -FilePath $reportFilePath -Encoding:utf8NoBOM


    
    
    
    
}


ConvertTo-Json  -InputObject $sites -Depth 10
| Out-File -FilePath (join-path $workdir "hubsite-spokes-pages.json") -Encoding:utf8NoBOM