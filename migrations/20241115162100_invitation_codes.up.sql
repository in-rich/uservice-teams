CREATE TABLE IF NOT EXISTS invitation_codes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    code TEXT NOT NULL,
    team_id UUID NOT NULL,

    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

--bun:split

CREATE FUNCTION ensure_unique_invitation_codes() RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM invitation_codes
        WHERE code = NEW.code
        AND expires_at > NOW()
    ) THEN
        RAISE SQLSTATE '23505' USING MESSAGE = 'Code already exists';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--bun:split

CREATE TRIGGER unique_invitation_codes
    BEFORE INSERT OR UPDATE ON invitation_codes
    FOR EACH ROW
    EXECUTE FUNCTION ensure_unique_invitation_codes();
