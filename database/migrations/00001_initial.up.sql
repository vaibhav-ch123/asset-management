BEGIN;

CREATE TYPE employee_types AS ENUM (
    'full_time',
    'intern',
    'freelancer'
);

CREATE TYPE employee_roles AS ENUM (
    'admin',
    'manager',
    'hr',
    'developer'
);

CREATE TABLE IF NOT EXISTS employees(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR (30) NOT NULL,
    email VARCHAR (30) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    phone VARCHAR (10) UNIQUE NOT NULL,
    joining_date DATE NOT NULL,
    employee_type employee_types NOT NULL,
    employee_role employee_roles NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    archived_at TIMESTAMP 
);

CREATE TYPE asset_types AS ENUM (
    'laptop', 
    'monitor',
    'mouse',
    'keyboard',
    'phone'
);

CREATE TYPE asset_brands AS ENUM (
    'dell',
    'hp',
    'lenovo',
    'apple',
    'logitech',
    'samsung'
);

CREATE TYPE asset_statuses AS ENUM (
    'available',
    'assigned',
    'waiting__repair',
    'in_service',
    'damaged',
    'archived'
);

CREATE TABLE IF NOT EXISTS assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_name VARCHAR (30) NOT NULL,
    asset_type asset_types NOT NULL,
    asset_brand asset_brands NOT NULL,
    serial_number VARCHAR(30) UNIQUE NOT NULL,
    purchase_date DATE,
    warranty_expiry DATE,
    asset_status asset_statuses NOT NULL,
    charger_available BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    archived_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laptop_specs (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID UNIQUE NOT NULL REFERENCES assets(id),
    ram_gb INTEGER,
    storage_gb INTEGER,
    operating_system VARCHAR,
    screen_resolution VARCHAR,
    processor VARCHAR
);

CREATE TABLE IF NOT EXISTS monitor_specs (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID UNIQUE NOT NULL REFERENCES assets(id),
    screen_size VARCHAR,
    screen_resolution VARCHAR
);

CREATE TABLE IF NOT EXISTS mouse_specs (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID UNIQUE NOT NULL REFERENCES assets(id),
    wireless BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS keyboard_specs (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID UNIQUE NOT NULL REFERENCES assets(id),
    wireless BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS phone_specs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID UNIQUE NOT NULL REFERENCES assets(id),
    ram_gb INTEGER,
    storage_gb INTEGER,
    operating_system VARCHAR
);

CREATE TABLE IF NOT EXISTS asset_assignments (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID REFERENCES employees(id) NOT NULL,
    asset_id UUID REFERENCES assets(id) NOT NULL,
    assigned_date DATE NOT NULL,
    returned_date DATE,
    remark VARCHAR,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    archived_at TIMESTAMP
);


COMMIT;