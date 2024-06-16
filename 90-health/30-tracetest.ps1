$env:NATS_SUBJECT = "hello"
write-host $env:NATS_SUBJECT
koksmat trace log "Step 1"
Start-Sleep -Seconds 1

koksmat trace log "Step 2"
Start-Sleep -Seconds 1

koksmat trace log "Step 3"
Start-Sleep -Seconds 1

koksmat trace log "Step 4"
Start-Sleep -Seconds 1
