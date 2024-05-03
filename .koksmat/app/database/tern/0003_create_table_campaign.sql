/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   




CREATE TABLE public.campaign
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,url character varying COLLATE pg_catalog."default"  NOT NULL
    ,key character varying COLLATE pg_catalog."default"  NOT NULL
    ,displayname character varying COLLATE pg_catalog."default"  NOT NULL
    ,owner_id int  


);

                ALTER TABLE IF EXISTS public.campaign
                ADD FOREIGN KEY (owner_id)
                REFERENCES public.person (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.campaign;

