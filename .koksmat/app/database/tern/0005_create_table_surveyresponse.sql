/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   




CREATE TABLE public.surveyresponse
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,url character varying COLLATE pg_catalog."default"  NOT NULL
    ,responsedate character varying COLLATE pg_catalog."default"  
    ,key character varying COLLATE pg_catalog."default"  NOT NULL
    ,displayname character varying COLLATE pg_catalog."default"  NOT NULL
    ,respondent_id int  
    ,survey_id int  
    ,application_id int  
    ,questions JSONB  
    ,answers JSONB  
    ,answer1 character varying COLLATE pg_catalog."default" 
    ,answer2 character varying COLLATE pg_catalog."default" 
    ,answer3 character varying COLLATE pg_catalog."default" 
    ,answer4 character varying COLLATE pg_catalog."default" 
    ,answer5 character varying COLLATE pg_catalog."default" 
    ,answer6 character varying COLLATE pg_catalog."default" 
    ,answer7 character varying COLLATE pg_catalog."default" 
    ,answer8 character varying COLLATE pg_catalog."default" 
    ,answer9 character varying COLLATE pg_catalog."default" 
    ,truefalse1 boolean  
    ,truefalse2 boolean  
    ,truefalse3 boolean  
    ,datetime1 character varying COLLATE pg_catalog."default"  
    ,datetime2 character varying COLLATE pg_catalog."default"  
    ,datetime3 character varying COLLATE pg_catalog."default"  
    ,number1 character varying COLLATE pg_catalog."default"  
    ,number2 character varying COLLATE pg_catalog."default"  
    ,number3 character varying COLLATE pg_catalog."default"  


);

                ALTER TABLE IF EXISTS public.surveyresponse
                ADD FOREIGN KEY (respondent_id)
                REFERENCES public.person (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                ALTER TABLE IF EXISTS public.surveyresponse
                ADD FOREIGN KEY (survey_id)
                REFERENCES public.survey (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                ALTER TABLE IF EXISTS public.surveyresponse
                ADD FOREIGN KEY (application_id)
                REFERENCES public.application (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.surveyresponse;

