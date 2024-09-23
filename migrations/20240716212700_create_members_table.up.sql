CREATE TYPE team_roles AS ENUM ('admin', 'member');

--bun:split

CREATE TABLE team_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    team_id UUID NOT NULL,
    user_id VARCHAR(255) NOT NULL,

    role team_roles NOT NULL DEFAULT 'member',

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_team_members UNIQUE (team_id, user_id)
);

--bun:split

CREATE INDEX teams_by_user_id ON team_members (user_id);
CREATE INDEX users_by_team_id ON team_members (team_id);
