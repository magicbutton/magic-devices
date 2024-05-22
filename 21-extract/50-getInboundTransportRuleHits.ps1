<#---
title: Get Exchange 
connection: exchange    
output: InboundTransportRuleHits.json
tag:  get-exchangestats
api: post
---
#>


param(
   $startDate = "2024-04-01",
   $endDate = "2024-04-30"
)
if ($null -eq $env:WORKDIR ) {
   $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
   New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir




$data = Get-MailTrafficSummaryReport -StartDate $startDate -EndDate $endDate -Category "InboundTransportRuleHits"  
#| ConvertTo-Json | Out-File InboundTransportRuleHits.json 
$outputfile = join-path $workdir "InboundTransportRuleHits.json"
$data  | ConvertTo-Json -Depth 10 | Out-File -FilePath (join-path $outputfile  "InboundTransportRuleHits.json") -Encoding utf8NoBOM
# -EventType Spam
#$trafficReport = Get-MailTrafficReport -StartDate $startDate -EndDate $endDate -EventType Spam

$env:POSTGRES_DB = "postgres://ps_user:SecurePassword@localhost:5432/files?sslmode=disable"
magic-devices import $outputfile 
magic-devices sql select "select * from InboundTransportRuleHits"

