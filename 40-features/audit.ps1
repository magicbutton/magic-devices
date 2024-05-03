<#---
title: Audit Log
aigenerated: true
aigenerator: ChatGPT 3.5
tag: auditlog
description: 'Designing a version control system for PostgreSQL to track insert, update, and delete operations on specific tables while utilizing audit columns like `id` and `who` for tracking changes is an interesting challenge. Here''s a basic approach to achieve this.'

---

Designing a version control system for PostgreSQL to track insert, update, and delete operations on specific tables while utilizing audit columns like `id` and `who` for tracking changes is an interesting challenge. Here's a basic approach to achieve this:

### Data Model:

1. **Audit Table:**
   - Create an audit table for each of your main tables to store the historical changes.
   - This table will store metadata about each change (like operation type, timestamp, and user).
   - Here's a basic structure:

```sql
CREATE TABLE table_name_audit (
    audit_id SERIAL PRIMARY KEY,
    operation_type CHAR(1) NOT NULL, -- 'I' for insert, 'U' for update, 'D' for delete
    operation_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    who VARCHAR(255) NOT NULL,
    id INT NOT NULL, -- Reference to the corresponding row in the main table
    -- other_audit_columns -- If needed
);
```

2. * * Main Tables:**
- Your existing tables with the `id` and `who` columns.

### SQL Triggers:

Now, you need to set up triggers on the main tables to automatically record changes into the corresponding audit tables.

1. * * Insert Trigger:**

```sql
CREATE OR REPLACE FUNCTION table_name_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
INSERT INTO table_name_audit (operation_type, who, id)
VALUES ('I', NEW.who, NEW.id);
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_name_insert
AFTER INSERT ON table_name
FOR EACH ROW
EXECUTE FUNCTION table_name_insert_trigger();
```

2. * * Update Trigger:**

```sql
CREATE OR REPLACE FUNCTION table_name_update_trigger()
RETURNS TRIGGER AS $$
BEGIN
INSERT INTO table_name_audit (operation_type, who, id)
VALUES ('U', NEW.who, NEW.id);
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_name_update
AFTER UPDATE ON table_name
FOR EACH ROW
EXECUTE FUNCTION table_name_update_trigger();
```

3. * * Delete Trigger:**

```sql
CREATE OR REPLACE FUNCTION table_name_delete_trigger()
RETURNS TRIGGER AS $$
BEGIN
INSERT INTO table_name_audit (operation_type, who, id)
VALUES ('D', OLD.who, OLD.id);
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_name_delete
BEFORE DELETE ON table_name
FOR EACH ROW
EXECUTE FUNCTION table_name_delete_trigger();
```

### Usage:

- Whenever an insert, update, or delete operation occurs on the main table, the corresponding trigger will automatically record the change in the audit table.
- You can query the audit table to retrieve historical changes and track who initiated each change.

This is a basic setup to get you started. Depending on your specific requirements, you might need to add more functionality or customize the solution further. Additionally, ensure that proper permissions are set to control access to the audit tables and triggers.

#>