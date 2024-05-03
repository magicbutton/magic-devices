do $$
    declare
    arow record;
    key varchar(500);
    name varchar(500);
    owner varchar(500);
BEGIN
    FOR arow IN SELECT "Application Name (column 0)" as name,
    "Unique Code (column 1)" as key,
    "Application Owner (column 32)" AS owner FROM "excel"."applications"
    LOOP
        key = arow.key;
        name = arow.name;
        owner = arow.owner;
        INSERT INTO public.application(
	id, created_at, updated_at, deleted_at, tenant, name, description, key, displayname, imported_ownername, owner_id)
        VALUES (DEFAULT, DEFAULT, DEFAULT, DEFAULT, '','displayname:' || name || ' key:' || key, '',key,name,owner , (SELECT id FROM public.person WHERE displayname = owner LIMIT 1));
    END LOOP;
END;
$$;