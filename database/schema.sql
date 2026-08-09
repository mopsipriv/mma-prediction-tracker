-- 1. Predictors (Sources of fight predictions: humans or AI models)
CREATE TABLE IF NOT EXISTS predictors (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    type VARCHAR(20) NOT NULL CHECK (type IN ('human', 'ai')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. Fighters
CREATE TABLE IF NOT EXISTS fighters (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    nickname VARCHAR(50),
    weight_class VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. Events / Tournaments
CREATE TABLE IF NOT EXISTS events (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    event_date TIMESTAMP WITH TIME ZONE NOT NULL,
    location VARCHAR(150),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 4. Fights
CREATE TABLE IF NOT EXISTS fights (
    id BIGSERIAL PRIMARY KEY,
    event_id BIGINT NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    fighter_a_id BIGINT NOT NULL REFERENCES fighters(id),
    fighter_b_id BIGINT NOT NULL REFERENCES fighters(id),
    weight_class VARCHAR(50) NOT NULL,
    
    -- Fight metadata
    is_title_fight BOOLEAN DEFAULT FALSE,
    is_main_event BOOLEAN DEFAULT FALSE,
    
    -- Closing betting odds
    odds_fighter_a DECIMAL(5, 2),
    odds_fighter_b DECIMAL(5, 2),
    
    -- Fight results (NULL until completed)
    winner_id BIGINT REFERENCES fighters(id),
    win_method VARCHAR(50),
    win_round INT CHECK (win_round BETWEEN 1 AND 5),
    win_time VARCHAR(10),                             -- e.g., '3:42'
    status VARCHAR(20) DEFAULT 'scheduled' CHECK (status IN ('scheduled', 'completed', 'cancelled')),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_different_fighters CHECK (fighter_a_id <> fighter_b_id)
);

-- 5. Predictions
CREATE TABLE IF NOT EXISTS predictions (
    id BIGSERIAL PRIMARY KEY,
    fight_id BIGINT NOT NULL REFERENCES fights(id) ON DELETE CASCADE,
    predictor_id BIGINT NOT NULL REFERENCES predictors(id) ON DELETE CASCADE,
    predicted_winner_id BIGINT NOT NULL REFERENCES fighters(id),
    
    -- Probability of victory (0.00 to 1.00)
    probability DECIMAL(3, 2) NOT NULL CHECK (probability >= 0.00 AND probability <= 1.00),
    
    -- Extended prediction details
    predicted_method VARCHAR(50) CHECK (predicted_method IN ('KO/TKO', 'Submission', 'Decision', 'Any')),
    predicted_round INT CHECK (predicted_round BETWEEN 1 AND 5),
    reasoning TEXT,                                   -- Detailed AI / human breakdown
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_predictor_per_fight UNIQUE (fight_id, predictor_id)
);