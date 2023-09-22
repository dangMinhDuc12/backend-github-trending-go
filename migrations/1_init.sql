
-- +migrate Up
CREATE TABLE public.users (
  user_id TEXT PRIMARY KEY,
  full_name TEXT,
  email TEXT UNIQUE,
  password TEXT,
  role TEXT,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE  public.repos (
  repo_name TEXT PRIMARY KEY,
  description TEXT,
  url TEXT,
  color TEXT,
  lang TEXT,
  fork TEXT,
  stars TEXT,
  stars_today TEXT,
  build_by TEXT,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE public.bookmarks (
  bid TEXT PRIMARY KEY,
  user_id TEXT,
  repo_name TEXT,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

ALTER TABLE public.bookmarks ADD FOREIGN KEY (user_id) REFERENCES public.users (user_id);
ALTER TABLE public.bookmarks ADD FOREIGN KEY (repo_name) REFERENCES public.repos (repo_name);


-- +migrate Down
DROP TABLE public.bookmarks;
DROP TABLE public.users;
DROP TABLE public.repos;
