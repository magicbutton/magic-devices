do $$
    declare
    arow record;
    foo varchar(500);
BEGIN
    FOR arow IN SELECT DISTINCT "Application Owner (column 32)" AS name FROM "excel"."applications"
    LOOP
        foo = arow.name;
        INSERT INTO public.person(
	id, created_at, updated_at, deleted_at, tenant, name, description, unique_person_id, displayname, email)
        VALUES (DEFAULT, DEFAULT, DEFAULT, DEFAULT, '','displayname:' || foo , '',foo,foo,'');
    END LOOP;
END;
$$;