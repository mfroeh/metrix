CREATE TABLE IF NOT EXISTS summoners (
    id BIGSERIAL PRIMARY KEY,
    puuid CHAR(78) NOT NULL UNIQUE,
    account_id VARCHAR(56) NOT NULL UNIQUE,
    profile_icon_id INTEGER NOT NULL,
    summoner_id VARCHAR(63) NOT NULL UNIQUE,
    summoner_level INTEGER NOT NULL,
    revision_date timestamp with time zone NOT NULL,
    name CITEXT NOT NULL,
    tag CITEXT NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW()
);

ALTER TABLE summoners
ADD CONSTRAINT check_name_length CHECK (LENGTH(name) >= 3 AND LENGTH(name) <= 16),
ADD CONSTRAINT check_tag_length CHECK (LENGTH(tag) >= 3 AND LENGTH(tag) <= 5);
ALTER TABLE summoners
ADD CONSTRAINT check_summoner_level CHECK (summoner_level >= 1);