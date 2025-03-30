CREATE TABLE IF NOT EXISTS comments (
  id bigserial PRIMARY KEY,
  post_id bigserial NOT NULL,
  user_id bigserial NOT NULL,
  content TEXT NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

-- INSERT INTO comments (post_id, user_id, content) VALUES(
--   1,1, 'first comment on your post'
-- );
-- INSERT INTO comments (post_id, user_id, content) VALUES(
--   2,1, 'second comment on your post'
-- );