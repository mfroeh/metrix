CREATE TABLE IF NOT EXISTS matches (
    id BIGSERIAL PRIMARY KEY,
    match_id TEXT NOT NULL UNIQUE,
    data_version TEXT NOT NULL,
    end_of_game_result TEXT NOT NULL,
    game_creation TIMESTAMP WITH TIME ZONE NOT NULL,
    game_datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    game_id BIGINT NOT NULL,
    game_length FLOAT NOT NULL,
    game_version TEXT NOT NULL,
    map_id INTEGER NOT NULL,
    queue_id INTEGER NOT NULL,
    tft_game_type TEXT NOT NULL,
    tft_set_number INTEGER NOT NULL,
    tft_set_name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS matches_participants (
    id BIGSERIAL PRIMARY KEY,
    match_id BIGINT NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
    puuid CHAR(78) NOT NULL,
    gold_left INTEGER NOT NULL,
    last_round INTEGER NOT NULL,
    level INTEGER NOT NULL,
    placement INTEGER NOT NULL,
    player_eliminated INTEGER NOT NULL,
    total_damage_to_players INTEGER NOT NULL,
    time_eliminated FLOAT NOT NULL,

    companion_content_id TEXT NOT NULL,
    companion_item_id INTEGER NOT NULL,
    companion_skin_id INTEGER NOT NULL,
    companion_species TEXT NOT NULL,

    augments TEXT[] NOT NULL,

    CONSTRAINT matches_participants_match_id_puuid_key UNIQUE (match_id, puuid)
);

CREATE TABLE IF NOT EXISTS matches_participants_traits (
    id BIGSERIAL PRIMARY KEY,
    match_id BIGINT NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
    participant_id BIGINT NOT NULL REFERENCES matches_participants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    num_units INTEGER NOT NULL,
    style INTEGER NOT NULL,
    tier_current INTEGER NOT NULL,
    tier_total INTEGER NOT NULL,

    CONSTRAINT traits_match_id_participant_id_name_key UNIQUE (match_id, participant_id, name)
);

CREATE TABLE IF NOT EXISTS matches_participants_units (
    id BIGSERIAL PRIMARY KEY,
    match_id BIGINT NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
    participant_id BIGINT NOT NULL REFERENCES matches_participants(id) ON DELETE CASCADE,
    character_id TEXT NOT NULL,
    items TEXT[] NOT NULL,
    name TEXT NOT NULL,
    rarity INTEGER NOT NULL,
    tier INTEGER NOT NULL
);
