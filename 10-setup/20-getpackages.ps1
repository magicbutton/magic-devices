<#---
title: Download Mate
output: koksmat-mate.zip
---


#>

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions") {
    $path = join-path $PSScriptRoot ".." ".."
}
else {
    $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path

$releaseUrl = "https://github.com/magicbutton/magic-mix/archive/refs/tags/v0.0.3.3.zip"
$request = Invoke-RestMethod $releaseUrl -Method 'GET' -Headers $headers -OutFile (join-path $koksmatDir "packages" "magic-mix.zip")


