
select  
       'w5' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator,* from exceldevice.tabelle 
WHERE "Nexi + SIA IT (column 1)" = 'Antivirus installed and monitored'
ORDER BY id
LIMIT 1;


select  
       'w6' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator,* from exceldevice.tabelle 
WHERE "Nexi + SIA IT (column 1)" = 'DLP installed and monitored'
ORDER BY id
LIMIT 1;

select  
       'w4' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator,* from exceldevice.tabelle 
WHERE "Nexi + SIA IT (column 1)" = 'Endpoint Detection & Response installed and monitored'
ORDER BY id
LIMIT 1;

select  
       'w3' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator,* from exceldevice.tabelle 
WHERE "Nexi + SIA IT (column 1)" = 'Disk-encryption (Bitlocker)'
ORDER BY id
LIMIT 1;



select  
       'w7a' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator,* from exceldevice.tabelle 
WHERE "Nexi + SIA IT (column 1)" = 'Users with USB enablement authorization (with encryption)'
ORDER BY id
LIMIT 1;

select  
       'w2' as KPI,
       " (column 2)" AS Numerator,
       " (column 4)" AS Denominator,
       *

FROM exceldevice.intune
where "INTUNE (column 1)" = 'Totale complessivo'

ORDER BY id









select  
       'w7b' as KPI,
       " (column 4)" AS Numerator,
       " (column 5)" AS Denominator,* from exceldevice.tabelle 
WHERE "Nexi + SIA IT (column 1)" = 'Users with USB enablement authorization (without encryption)'
ORDER BY id
LIMIT 1;




