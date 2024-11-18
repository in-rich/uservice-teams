DROP FUNCTION IF EXISTS ensure_unique_invitation_codes;

--bun:split

DROP TRIGGER IF EXISTS unique_invitation_codes ON invitation_codes;

--bun:split

DROP TABLE IF EXISTS invitation_codes;
