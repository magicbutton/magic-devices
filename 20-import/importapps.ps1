
<#---
title: Import CSV
tag: import-csv
api: post
output: apps.json
---#>
param (

    [string]$file = "Estrazione Catalogo NEAR_20240404.csv"
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
$inputfile = join-path $workdir $file

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
<#

Here we start working
#>


$apps = Import-Csv -Path $inputfile -Delimiter ";"

$outputfilename = (join-path $workdir "apps.json")
$outputcodefilename = (join-path $workdir "apps.go")

ConvertTo-Json  -InputObject $apps  -Depth 10
| Out-File -FilePath $outputfilename  -Encoding:utf8NoBOM

$apps2 = @()
<#
 "Application Name": "1app8",
    "Unique Code": "AC_583",
    "MVF Code": "UAP",
    "Alias": "Mobile Wallet, Wallet 1.0, SDK",
    "Application Type": "BA",
    "Application Category (Sterotype)": "APP",
    "Application Status": "Operation",
    "Application Status Notes": "Attivo",
    "WBS": "Non Definito",
    "Sub Component": "FALSE",
    "Description": "Applicazione che eroga servizi di payment acceptance alle piattaforme mobile SDK, sviluppate nativamente per Android e iOS. Tra i servizi erogati: user on-boarding, card enrollment and tokenization, payment acceptance, digital couponing, digital loyalty, transaction lifecycle management.",
    "Notes": "L'applicazione sarà affiancata nel 2023 da nuova soluzione \"Wallet 2.0\".\nServer di riferimento:\npptserverpeprd01\npptserverseprd01\n\nI merchant utilizzatori sono clienti Nexi Payments e sono: Conad, Evoca SRL, Cremonini SpA (Roadhouse),  Moneymanagement (per Vodafone SpA), Esselunga SpA, Eataly SpA",
    "User Category": "Esterni",
    "Internet Exposure": "TRUE",
    "URL": "https://merchant.sparkling18.com/",
    "Public IP": "185.198.117.173",
    "Data Location": "Europeo",
    "Building Type": "Built-in",
    "SW Repository": "Bitbucket",
    "URL SW Repository": "https://bitbucket.nexicloud.it/projects/LB/repos/ppt-server/browse",
    "Capital Stock Presence": "TRUE",
    "Capital Stock Notes": "Indicato nell'applicativo",
    "Reference Model Platform": "Channel & Acceptance",
    "Reference Model Level 1": "E-Commerce",
    "Reference Model Level 2": "Wallet",
    "Company": "Nexi Payments",
    "Division": "IT",
    "Director": "Giuseppe Dallona",
    "Competence Center": "Digital",
    "Competence Centre Head": "Alexandre Francesco Bove",
    "Office": "Digital Key Initiatives",
    "Office Head": "Marina Fisco",
    "Application Owner": "Marina Fisco",
    "Operational Contact": "Stefano Fiorino",
    "Business Office": "Online Connect",
    "Business Application Owner": "Francesco Vincenzo Perrini",
    "Application Ownership": "NEXI PAYMENTS",
    "User Group Companies": "NEXI PAYMENTS",
    "ID Service Offering": "SO00738",
    "Code Intellectual Property": "NEXI PAYMENTS",
    "Maintenance Contract Ownership": "NEXI PAYMENTS",
    "SW Maintenance Company": "NEXI PAYMENTS",
    "Service Model": "In House",
    "Platform Type": "Distributed",
    "Infrastructure Outsourcer (HW)": "NEXI PAYMENTS",
    "Production Server": "18",
    "NON Production Server": "18",
    "Systemic Relevance": "FALSE",
    "GDPR Flag": "TRUE",
    "GDPR Scope": "GDPR Dati anagrafici",
    "PCI Flag": "TRUE",
    "PCI Notes": "DSS",
    "PSD2 Flag": "TRUE",
    "PSD2 Notes": "Flussi progettati per aderire a SCA e gestione del soft decline",
    "ISO27001 Certification": "Vero",
    "ISAE3402 Certification": "Falso",
    "Tables Number": "212",
    "Programs Number": "2",
    "LoC Number": "NA",
    "Users Number": "200000",
    "Volumes": "Non Definito",
    "BCM Index": "0",
    "BIA IT": "1",
    "Declared RPO": "Realtime",
    "Declared RTO": "2h",
    "ITRA Flag": "TRUE",
    "Inventory Company": "NEXI PAYMENTS",
    "Verified By": "Luca Carini",
    "APP Creation Date": "21/06/2019",
    "Obsolescence Index": "2,236067977",
    "App Total Cost (beta)": "",
    "% Completeness of Mandatory Info": "0,952380952",
    "% Completeness of Optional Info": "1",
    "% Informations Completeness": "0,957142857"
#>
foreach ($app in $apps) {
    $app2 = @{
        "Application Name"                 = $app."Application Name"
        "Unique Code"                      = $app."Unique Code"
        "MVF Code"                         = $app."MVF Code"
        "Alias"                            = $app."Alias"
        "Application Type"                 = $app."Application Type"
        "Application Category (Sterotype)" = $app."Application Category (Sterotype)"
        "Application Status"               = $app."Application Status"
        "Application Status Notes"         = $app."Application Status Notes"
        "WBS"                              = $app."WBS"
        "Sub Component"                    = $app."Sub Component"
        "Description"                      = $app."Description"
        "Notes"                            = $app."Notes"
        "User Category"                    = $app."User Category"
        "Internet Exposure"                = $app."Internet Exposure"
        "URL"                              = $app."URL"
        "Public IP"                        = $app."Public IP"
        "Data Location"                    = $app."Data Location"
        "Building Type"                    = $app."Building Type"
        "SW Repository"                    = $app."SW Repository"
        "URL SW Repository"                = $app."URL SW Repository"
        "Capital Stock Presence"           = $app."Capital Stock Presence"
        "Capital Stock Notes"              = $app."Capital Stock Notes"
        "Reference Model Platform"         = $app."Reference Model Platform"
        "Reference Model Level 1"          = $app."Reference Model Level 1"
        "Reference Model Level 2"          = $app."Reference Model Level 2"
        "Company"                          = $app."Company"
        "Division"                         = $app."Division"
        "Director"                         = $app."Director"
        "Competence Center"                = $app."Competence Center"
        "Competence Centre Head"           = $app."Competence Centre Head"
        "Office"                           = $app."Office"
        "Office Head"                      = $app."Office Head"
        "Application Owner"                = $app."Application Owner"
        "Operational Contact"              = $app."Operational Contact"
        "Business Office"                  = $app."Business Office"
        "Business Application Owner"       = $app."Business Application Owner"
        "Application Ownership"            = $app."Application Ownership"
        "User Group Companies"             = $app."User Group Companies"
        "ID Service Offering"              = $app."ID Service Offering"
        "Code Intellectual Property"       = $app."Code Intellectual Property"
        "Maintenance Contract Ownership"   = $app."Maintenance Contract Ownership"
        "SW Maintenance Company"           = $app."SW Maintenance Company"
        "Service Model"                    = $app."Service Model"
        "Platform Type"                    = $app."Platform Type"
        "Infrastructure Outsourcer (HW)"   = $app."Infrastructure Outsourcer (HW)"
        "Production Server"                = $app."Production Server"
        "NON Production Server"            = $app."NON Production Server"
        "Systemic Relevance"               = $app."Systemic Relevance"
        "GDPR Flag"                        = $app."GDPR Flag"

        # "GDPR Scope"                       = $app."GDPR Scope"
        # "PCI Flag"                         = $app."PCI Flag"
        #"PCI Notes"                        = $app."PCI Notes"
        #"PSD2 Flag"                        = $app."PSD2 Flag"
        #"PSD2 Notes"                       = $app."PSD2 Notes"
        # "ISO27001 Certification"             = $app."ISO27001 Certification"
        # "ISAE3402 Certification"             = $app."ISAE3402 Certification"
        # "Tables Number"                      = $app."Tables Number"
        # "Programs Number"                    = $app."Programs Number"
        # "LoC Number"                         = $app."LoC Number"
        # "Users Number"                       = $app."Users Number"
        # "Volumes"                            = $app."Volumes"
        # "BCM Index"                          = $app."BCM Index"
        # "BIA IT"                             = $app."BIA IT"
        # "Declared RPO"                       = $app."Declared RPO"
        # "Declared RTO"                       = $app."Declared RTO"
        # "ITRA Flag"                          = $app."ITRA Flag"
        # "Inventory Company"                  = $app."Inventory Company"
        # "Verified By"                        = $app."Verified By"
        # "APP Creation Date"                  = $app."APP Creation Date"
        # "Obsolescence Index"                 = $app."Obsolescence Index"
        # "App Total Cost Beta"                = $app."App Total Cost (beta)"
        # "Pct Completeness of Mandatory Info" = $app."% Completeness of Mandatory Info"
        # "Pct Completeness of Optional Info"  = $app."% Completeness of Optional Info"
        # "Pct Informations Completeness"      = $app."% Informations Completeness"



    }
    $apps2 += $app2
}


$outputfilename2 = (join-path $workdir "apps2.json")

ConvertTo-Json  -InputObject $apps2  -Depth 10
| Out-File -FilePath $outputfilename2  -Encoding:utf8NoBOM


quicktype $outputfilename2 --out $outputcodefilename  --no-enums --debug all  #| Out-Null
 