DROP INDEX IF EXISTS teams_by_user_id;
DROP INDEX IF EXISTS users_by_team_id;

--bun:split

DROP TABLE IF EXISTS team_members;

--bun:split

DROP TYPE IF EXISTS team_roles;
