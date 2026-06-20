-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    telegram_id BIGINT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE organizations(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    standard_hours INT DEFAULT 8,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TYPE ROLE AS ENUM ('owner', 'manager', 'worker');
CREATE TYPE SALARY_TYPE AS ENUM('per/hour', 'per/day', 'per/month');

CREATE TABLE user_organizations(
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    org_id UUID REFERENCES organizations(id),
    role ROLE NOT NULL,
    salary_rate NUMERIC DEFAULT 0,
    salary_type SALARY_TYPE NOT NULL,
    joined_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (user_id, org_id)
);

CREATE TYPE STATUS AS ENUM('PENDING', 'VERIFIED', 'REJECTED');

CREATE TABLE attendance_records(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    org_id UUID REFERENCES organizations(id),
    checked_in_at TIMESTAMP NOT NULL DEFAULT NOW(),
    checked_out_at TIMESTAMP,
    checkout_initiator UUID REFERENCES users(id) ON DELETE SET NULL,
    is_auto_closed BOOLEAN DEFAULT FALSE,
    status STATUS NOT NULL DEFAULT 'PENDING',
    verified_by UUID REFERENCES users(id) ON DELETE SET NULL,
    verified_at TIMESTAMP,
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE schedules(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    org_id UUID REFERENCES organizations(id),
    days_of_week INT[],
    check_in_time TIME NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE schedules;
DROP TABLE attendance_records;
DROP TYPE STATUS;
DROP TABLE user_organizations;
DROP TYPE SALARY_TYPE;
DROP TYPE ROLE;
DROP TABLE users;
DROP TABLE organizations;
-- +goose StatementEnd

