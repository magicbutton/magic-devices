do $$
    declare
    arow record;
   
BEGIN
    FOR arow IN select id as application_id,displayname,owner_id, 1 as survey_id from "public"."application"
    LOOP
      
        INSERT INTO "public"."surveyresponse"(
	id, created_at, updated_at, deleted_at, tenant, name, description, url, "key", displayname, respondent_id, survey_id, application_id)
        VALUES (DEFAULT, DEFAULT, DEFAULT, DEFAULT, '', arow.displayname  , '', '', '', arow.displayname, arow.owner_id, arow.survey_id, arow.application_id);
    END LOOP;
END;
$$;