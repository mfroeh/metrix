CREATE TABLE IF NOT EXISTS leagues (
    id BIGSERIAL PRIMARY KEY,
    puuid CHAR(78) NOT NULL REFERENCES summoners(puuid) ON DELETE CASCADE,
    summoner_id VARCHAR(63) NOT NULL REFERENCES summoners(summoner_id) ON DELETE CASCADE,
    queue_type TEXT NOT NULL,
    tier TEXT NOT NULL,
    rank INTEGER NOT NULL,
    wins INTEGER NOT NULL,
    losses INTEGER NOT NULL,
    hot_streak BOOLEAN NOT NULL,
    veteran BOOLEAN NOT NULL,
    fresh_blood BOOLEAN NOT NULL,
    inactive BOOLEAN NOT NULL,
    league_points INTEGER NOT NULL,
    rated_rating INTEGER NOT NULL,
    mini_series_wins INTEGER NOT NULL,
    mini_series_losses INTEGER NOT NULL,
    mini_series_target INTEGER NOT NULL,
    mini_series_progress TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    CONSTRAINT unique_puuid_queue_type UNIQUE (puuid, queue_type)
);

ALTER TABLE leagues
ADD CONSTRAINT check_rank CHECK (rank >= 0),
ADD CONSTRAINT check_wins CHECK (wins >= 0),
ADD CONSTRAINT check_losses CHECK (losses >= 0);