/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   




CREATE TABLE public.device
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,unique_device_id character varying COLLATE pg_catalog."default"  NOT NULL
    ,displayname character varying COLLATE pg_catalog."default"  NOT NULL
    ,macaddress character varying COLLATE pg_catalog."default"  NOT NULL
    ,devicetype_id int   NOT NULL
    ,antivirus boolean  
    ,dlp boolean  
    ,endpointdetection boolean  
    ,diskentryptionprimarydisk boolean  
    ,diskentryptionsecondarydisk boolean  


);

                ALTER TABLE IF EXISTS public.device
                ADD FOREIGN KEY (devicetype_id)
                REFERENCES public.devicetype (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.device;

