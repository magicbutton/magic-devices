<#---
title: Get KPI report data
connection: graph
api: post
tag: devicekpi
---

Synopsis

1. Extracts mails from a specific mailfolder and downloads attachments to a specific folder.
2. If the attachment is an Excel file, it will be imported into a SQL database.
3. The imported data will be transformed into a JSON file
4. The file will be uploaded to a specific folder on a blob site
5. The mail will be moved to an archive folder.
6. An mail will be send with the result of the process.

#>



# Tested with
# -----------------------------------------
# IsIntegerOrMinusOne "nosad" # returns -1
# IsIntegerOrMinusOne "    1 " # returns 1
# IsIntegerOrMinusOne " 2 " # returns 2

function IsIntegerOrMinusOne($text) {
    # Parse the $text to see if it is an integer or -1

    $integer = -1
    try {
        $parsedInteger = [int]$text
        $integer = $parsedInteger
    }   
    catch {
        # Do nothing,  -1 is the value in case of error
    }
   
    return $integer
}


function Get-KPIValue {
    param (
        [Parameter(Mandatory = $true)]
        [array]$Data,
        [Parameter(Mandatory = $true)]
        [string]$KPI,
   
        [Parameter(Mandatory = $true)]
        [array]$Field
    )

    $result = $Data | Where-Object { $_.kpi -eq $KPI }
    if ($null -eq $result) {
        return -1
    }
    $value = $result.$Field
    if ($null -eq $value) {
        return -1
    }
    return $value
}
<#

## Extract KPIs using SQL
#>


function ExtractKPIs() {

    $data = magic-mix sql select "select * from devicekpi.kpi" | convertfrom-json
    
    $kpis = @{
        "created" = (get-date).ToString("yyyy-MM-ddTHH:mm:ssZ") 
        w5        = @{
            num = Get-KPIValue $data w5 numerator
            den = Get-KPIValue $data w5 denominator
        }
        w6        = @{
            num = Get-KPIValue $data w6 numerator
            den = Get-KPIValue $data w6 denominator
        }
        w8        = @{
            num = Get-KPIValue $data w8 numerator
            den = Get-KPIValue $data w8 denominator
            
        }
        w4        = @{
            num = Get-KPIValue $data w4 numerator
            den = Get-KPIValue $data w4 denominator
            
        }
        w3        = @{
            num = Get-KPIValue $data w3 numerator
            den = Get-KPIValue $data w3 denominator
        }
        w2        = @{
            num = Get-KPIValue $data w2 numerator
            den = Get-KPIValue $data w2 denominator
            
        }
        w7a       = @{
            num = Get-KPIValue $data w7a numerator
           
        }
        w7b       = @{
            num = Get-KPIValue $data w7b numerator
        }

        cs11      = @{
            num = Get-KPIValue $data cs11 numerator
        }
      
    
    }
    $kpis | ConvertTo-Json | Out-File -FilePath (join-path $workdir "kpis.json") -Encoding utf8NoBOM
 

}
<#

## Upload result to Blob
#>

function UploadBlob() {
    # Define the SAS URL and the local file path
    $sasUrl = $env:DEVICEKPIBLOB
    if ($null -eq $sasUrl) {
        throw "DEVICEKPIBLOB environment variable not set"
    }
    $localFilePath = (join-path $workdir "kpis.json")

    # Read the file content
    $fileContent = Get-Content -Path $localFilePath -Raw

    # Create a web request
    $webRequest = [System.Net.HttpWebRequest]::Create($sasUrl)
    $webRequest.Method = "PUT"

    # Set the content type
    $webRequest.ContentType = "application/octet-stream"
    $webRequest.ContentLength = $fileContent.Length
    $webRequest.Headers.Add("x-ms-blob-type", "BlockBlob")

    # Write the file content to the request stream
    $requestStream = $webRequest.GetRequestStream()
    $requestStream.Write([System.Text.Encoding]::UTF8.GetBytes($fileContent), 0, $fileContent.Length)
    $requestStream.Close()

    # Get the response
    $response = $webRequest.GetResponse()

    # Display the response status code
    Write-Output "Response Status Code: $($response.StatusCode)"

}

<#

## Process Mail

Prepare mail routine which will look for mails with attachments and process them

#>
if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}



$workdir = Resolve-Path $workdir
Set-Location $workdir
write-host "Workdir: $workdir"

ExtractKPIs
UploadBlob
return
$to = "niels.johansen@nexigroup.com"
$from = "valerio.moles@external.nexigroup.com"


# $mailfolders = GraphAPI $env:GRAPH_ACCESSTOKEN "GET" "https://graph.microsoft.com/v1.0/users/$to/mailFolders"

$inboxFolderId = "inbox" # $mailfolders.value | Where-Object { $_.displayName -eq "Inbox" } | Select-Object -ExpandProperty id
$archiveFolderId = "archive" # $mailfolders.value | Where-Object { $_.displayName -eq "Archive" } | Select-Object -ExpandProperty id

$url = @"
https://graph.microsoft.com/v1.0/users/$to/messages?$('$')filter=parentFolderId eq '$inboxFolderId' and from/emailAddress/address eq '$from' and hasAttachments eq true&$('$')expand = attachments
"@


koksmat trace log "Calling Graph to get mails to $to with attachments from $from" 
$mails = GraphAPI $env:GRAPH_ACCESSTOKEN "GET" $url 
# $mails = Get-Content (join-path $workdir "mails.json") | ConvertFrom-Json

if ($null -eq $mails) {
    throw "Failed to get mails"
}

if ($mails.value.Count -eq 0) {
    koksmat trace log  "No mails found" 
    return
}

if ($mails.value.Count -gt 1) {
    throw "More than 1 mail found"
  
}
$foundExcelFile = $false
$mails.value | ForEach-Object {
   
    $mail = $_
    koksmat trace log  "Got mail with subject $($mail.subject)" 

    $mail.attachments | ForEach-Object {
        
        $attachment = $_
        if ($attachment.contentType -ne "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") {
            write-host "Attachment is not an Excel file" -ForegroundColor Yellow
            
        }
        else {
            koksmat trace log  "Found Excel file $($attachment.name)" 

            
            $filename = $attachment.name
 
            $base64String = $attachment.contentBytes
            $fileBytes = [System.Convert]::FromBase64String($base64String)

            $excelfilename = join-path $workdir "devicekpi.xlsx"
            [System.IO.File]::WriteAllBytes($excelfilename, $fileBytes)

            Write-Host "Excel file has been successfully written to $excelfilename"
            $foundExcelFile = $true

            koksmat trace log "Converting Excel to JSON"
            magic-mix from excel to json $excelfilename devicekpi

            koksmat trace log "Uploading JSON"
            magic-mix upload devicekpi

            koksmat trace log "Reading KPIs from database"
            ExtractKPIs

            koksmat trace log "Uploading KPIs to Blob"
            UploadBlob
      

          
        }    
    }
    if ($foundExcelFile) {
        koksmat trace log "Moving mail to archive"
        $moveMessageUrl = "https://graph.microsoft.com/v1.0/users/$to/messages/$($mail.id)/move"

        $moveData = convertto-json @{
            destinationId = $archiveFolderId
        }

        $moveMessageResponse = GraphAPI $env:GRAPH_ACCESSTOKEN "POST"  $moveMessageUrl  $moveData

        if ($null -eq $moveMessageResponse) {
            throw "Failed to move message"
        }
       
    }
    koksmat trace log "Done"

}
