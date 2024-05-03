create OR REPLACE view application_view as
select application.key as key,  application.displayname as name, person.name as owner from application
 left join person on application.owner_id = person.id
