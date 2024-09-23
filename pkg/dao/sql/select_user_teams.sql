SELECT teams.id, teams.owner_id, teams.name, teams.created_at
FROM team_members JOIN teams ON team_members.team_id = teams.id
WHERE team_members.user_id = ?
ORDER BY team_members.created_at DESC
LIMIT ? OFFSET ?;
