


do $$
    declare
    arow record;

BEGIN
    FOR arow IN SELECT public.person.id, "uniqueid (column 7)" as uniqueid, "email (column 5)" as email from "people"."person" 
        LEFT JOIN public.person ON public.person.displayname ILIKE "people"."person"."displayname (column 6)"
        WHERE public.person.displayname IS NOT NULL
        ORDER BY public.person.displayname
       
    LOOP
  
       UPDATE public.person SET  unique_person_id=arow.uniqueid,  email=arow.email WHERE id = arow.id;

    END LOOP;
END;
$$;