/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   




CREATE TABLE public.survey
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
    ,campaign_id int  
    ,question1 character varying COLLATE pg_catalog."default" 
    ,question2 character varying COLLATE pg_catalog."default" 
    ,question3 character varying COLLATE pg_catalog."default" 
    ,question4 character varying COLLATE pg_catalog."default" 
    ,question5 character varying COLLATE pg_catalog."default" 
    ,question6 character varying COLLATE pg_catalog."default" 
    ,question7 character varying COLLATE pg_catalog."default" 
    ,question8 character varying COLLATE pg_catalog."default" 
    ,question9 character varying COLLATE pg_catalog."default" 
    ,truefalse1 character varying COLLATE pg_catalog."default" 
    ,truefalse2 character varying COLLATE pg_catalog."default" 
    ,truefalse3 character varying COLLATE pg_catalog."default" 
    ,datetime1 character varying COLLATE pg_catalog."default" 
    ,datetime2 character varying COLLATE pg_catalog."default" 
    ,datetime3 character varying COLLATE pg_catalog."default" 
    ,number1 character varying COLLATE pg_catalog."default" 
    ,number2 character varying COLLATE pg_catalog."default" 
    ,number3 character varying COLLATE pg_catalog."default" 
    ,questions JSONB  


);

                ALTER TABLE IF EXISTS public.survey
                ADD FOREIGN KEY (owner_id)
                REFERENCES public.person (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                ALTER TABLE IF EXISTS public.survey
                ADD FOREIGN KEY (campaign_id)
                REFERENCES public.campaign (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.survey;

