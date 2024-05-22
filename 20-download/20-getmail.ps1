<#---
title: Get mails
connection: graph
api: post
tag: getmails
---#>


if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"

$to = "niels.johansen@nexigroup.com"
$from = "valerio.moles@external.nexigroup.com"

$mailfolders = GraphAPI $env:GRAPH_ACCESSTOKEN "GET" "https://graph.microsoft.com/v1.0/users/$to/mailFolders"

$inboxFolderId = $mailfolders.value | Where-Object { $_.displayName -eq "Inbox" } | Select-Object -ExpandProperty id
if ($null -eq $inboxFolderId) {
    throw "Inbox Mailfolder not found"
}
$archiveFolderId = $mailfolders.value | Where-Object { $_.displayName -eq "Archive" } | Select-Object -ExpandProperty id
if ($null -eq $archiveFolderId) {
    throw "Archive Mailfolder not found"
}





$url = @"
https://graph.microsoft.com/v1.0/users/$to/messages?$('$')filter=parentFolderId eq '$inboxFolderId' and from/emailAddress/address eq '$from' and hasAttachments eq true&$('$')expand = attachments
"@

#write-host $url
#$url | Set-Clipboard
$mails = GraphAPI $env:GRAPH_ACCESSTOKEN "GET" $url 
# $mails = Get-Content (join-path $workdir "mails.json") | ConvertFrom-Json

if ($null -eq $mails) {
    throw "Failed to get mails"
}
write-host $mails.value.Count  -ForegroundColor Yellow
if ($mails.value.Count -eq 0) {
    write-host "No mails found"
    return
}


$mails.value | ForEach-Object {
   
    $mail = $_
    write-host $mail.subject ($inboxFolderId -eq $mail.parentFolderId) -ForegroundColor Yellow
 
    $mailFolder = join-path $workdir "mails" ("$($mail.receivedDateTime)-$($mail.subject)").Replace("/", "")
    if (-not (Test-Path $mailFolder)) {
        New-Item -Path $mailFolder -ItemType Directory | Out-Null
    }
    $mail.attachments | ForEach-Object {
  
        $attachment = $_
        write-host $attachment.name
        $filename = $attachment.name
 
        # Your base64 encoded string
        $base64String = $attachment.contentBytes

        # Decode the base64 string
        $fileBytes = [System.Convert]::FromBase64String($base64String)

        # Specify the file path where you want to save the Excel file
        $filePath = join-path $mailFolder $filename

        # Write the decoded bytes to a file
        [System.IO.File]::WriteAllBytes($filePath, $fileBytes)

        Write-Host "Excel file has been successfully written to $filePath"
    
        #$attachment.path = $path
    }

    $moveMessageUrl = "https://graph.microsoft.com/v1.0/users/$to/messages/$($mail.id)/move"

    $moveData = convertto-json @{
        destinationId = $archiveFolderId
    }

    $moveMessageResponse = GraphAPI $env:GRAPH_ACCESSTOKEN "POST"  $moveMessageUrl  $moveData

    if ($null -eq $moveMessageResponse) {
        throw "Failed to move message"
    }
    write-host "Was message moved" ($moveMessageResponse.parentFolderId -eq $archiveFolderId) -ForegroundColor Green



}
#$mails | ConvertTo-Json -Depth 10 | Out-File -FilePath (join-path $workdir "mails.json") -Encoding utf8NoBOM
