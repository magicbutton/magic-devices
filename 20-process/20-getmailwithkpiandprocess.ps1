<#---
title: Get KPI report data
connection: graph, postgres
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
        <#Do this if a terminating exception happens#>
    }
   
    return $integer
}


<#

## Convert Excel to SQL
#>
function ConvertExcelToSQL($sheetname, $tablename) {
    $filename = join-path $workdir "devicekpi.xlsx"

    
    write-host "Workdir: $workdir"
    
    $exportdir = join-path $workdir "sqlimport" "devicekpi"
    if (-not (Test-Path $exportdir)) {
        New-Item -Path $exportdir -ItemType Directory | Out-Null
    }
    else {
        Remove-Item -Path "$($exportdir)/$($tablename)*" -Recurse -Force
    }
    
    Push-Location
    Set-Location $exportdir
    write-host $filename 
    try {

        magic-mix from excel to sql $filename $sheetname $tablename 
    }
    catch {
        Write-Host "Error: $_" -ForegroundColor Red
    }
    finally {
        Pop-Location
    }
}

<#

## Process SQL
#>
function RunSQL() {
    $sqldir = join-path $workdir "sqlimport" "devicekpi"

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
    Pop-Location
}

<#

## Extract KPIs using SQL
#>


function ExtractKPIs() {
    $w5SQL = @"
    select
    'w5' as KPI,
    " (column 4)" AS Numerator,
    " (column 5)" AS Denominator
from
    "excelimport"."devicekpi_tables"
WHERE
    "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'Antivirus installed and monitored'
ORDER BY
    id
LIMIT
    1
"@

    $w6SQL = @"
    select
    'w6' as KPI,
    " (column 4)" AS Numerator,
    " (column 5)" AS Denominator
from
    "excelimport"."devicekpi_tables"
WHERE
    "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'DLP installed and monitored'
ORDER BY
    id
LIMIT
    1    
"@
    $w8SQL = @"
    select
    'w8' as KPI,
    " (column 4)" AS Numerator,
    " (column 5)" AS Denominator
    from
    "excelimport"."devicekpi_tables"
    WHERE
    "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'System Software patched'
    ORDER BY
    id
    LIMIT
    1     
"@

    $w4SQL = @"
select
       'w4' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator
from
       "excelimport"."devicekpi_tables"
WHERE
       "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'Endpoint Detection & Response installed and monitored'
ORDER BY
       id
LIMIT
       1
"@
    $w3SQL = @"
select
       'w3' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator
from
       "excelimport"."devicekpi_tables"
WHERE
       "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'Disk-encryption (Bitlocker)'
ORDER BY
       id
LIMIT
       1
"@

    $w7aSQL = @"
select
       'w7a' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator
from
       "excelimport"."devicekpi_tables"
WHERE
       "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'Users with USB enablement authorization (with encryption)'
ORDER BY
       id
LIMIT
       1
"@

    $w2SQL = @"
select
       'w2' as KPI,
       " (column 2)" AS Numerator,
       " (column 4)" AS Denominator
FROM
       "excelimport"."devicekpi_intune"
where
       "INTUNE (column 1)" = 'Totale complessivo'
ORDER BY
       id
LIMIT
       1
"@

    $w7bSQL = @"
select
       'w7b' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator
from
       "excelimport"."devicekpi_tables"
WHERE
       "All PC / CORP+CARD+CONCARDIS+CMG (column 1)" = 'Users with USB enablement authorization (without encryption)'
ORDER BY
       id
LIMIT
       1
"@
    $w2 = magic-devices sql select $w2SQL | convertfrom-json
    $w3 = magic-devices sql select $w3SQL | convertfrom-json
    $w4 = magic-devices sql select $w4SQL | convertfrom-json
    $w5 = magic-devices sql select $w5SQL | convertfrom-json
    $w6 = magic-devices sql select $w6SQL | convertfrom-json
    $w7a = magic-devices sql select $w7aSQL | convertfrom-json
    $w7b = magic-devices sql select $w7bSQL | convertfrom-json
    $w8 = magic-devices sql select $w8SQL | convertfrom-json
    
    $kpis = @{
        "created" = (get-date).ToString("yyyy-MM-ddTHH:mm:ssZ") 
        w5        = @{
            num = IsIntegerOrMinusOne $w5.Numerator
            den = IsIntegerOrMinusOne $w5.Denominator
        }
        w6        = @{
            num = IsIntegerOrMinusOne $w6.Numerator
            den = IsIntegerOrMinusOne $w6.Denominator
        }
        w8        = @{
            num = IsIntegerOrMinusOne $w8.Numerator
            den = IsIntegerOrMinusOne $w8.Denominator
        }
        w4        = @{
            num = IsIntegerOrMinusOne $w4.Numerator
            den = IsIntegerOrMinusOne $w4.Denominator
        }
        w3        = @{
            num = IsIntegerOrMinusOne $w3.Numerator
            den = IsIntegerOrMinusOne $w3.Denominator
        }
        w2        = @{
            num = IsIntegerOrMinusOne $w2.Numerator
            den = IsIntegerOrMinusOne $w2.Denominator
        }
        w7a       = @{
            num = IsIntegerOrMinusOne $w7a.Numerator
        }
        w7b       = @{
            num = IsIntegerOrMinusOne $w7b.Numerator
        }

        cs11      = @{
            num = -1
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

write-host "Workdir: $workdir"

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

            koksmat trace log "Converting Excel to SQL"
            ConvertExcelToSQL  "tables"  "devicekpi_tables"
            ConvertExcelToSQL  "intune"  "devicekpi_intune"

            koksmat trace log "Executing SQL"
            RunSQL

            koksmat trace log "Quering SQL"

            ExtractKPIs

            koksmat trace log "Quering SQL"

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
