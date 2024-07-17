CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    owner_id VARCHAR(255) NOT NULL,
    name VARCHAR(64) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
