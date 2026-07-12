BEGIN;

DROP TABLE IF EXISTS asset_assignments;

DROP TABLE IF EXISTS phone_specs;
DROP TABLE IF EXISTS keyboard_specs;
DROP TABLE IF EXISTS mouse_specs;
DROP TABLE IF EXISTS monitor_specs;
DROP TABLE IF EXISTS laptop_specs;

DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS employees;

DROP TYPE IF EXISTS asset_statuses;
DROP TYPE IF EXISTS asset_brands;
DROP TYPE IF EXISTS asset_types;

DROP TYPE IF EXISTS employee_roles;
DROP TYPE IF EXISTS employee_types;

COMMIT;