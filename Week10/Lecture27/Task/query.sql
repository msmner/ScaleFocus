-- name: SelectRecords :many
SELECT * FROM stories;

-- name: DeleteStory :exec
DELETE FROM stories;

-- name: CreateStory :execresult
INSERT INTO stories (story_id, story_title, story_score, created_at) VALUES (?, ?, ?, ?);

