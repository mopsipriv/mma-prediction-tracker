-- Reset tables and auto-increment counters before seeding
TRUNCATE TABLE predictions, fights, events, fighters, predictors RESTART IDENTITY CASCADE;

-- 1. Insert Predictors
INSERT INTO predictors (name, type) VALUES
('me', 'human'),
('GPT-4o', 'ai'),
('DeepSeek-V3', 'ai'),
('Claude-3.5-Sonnet', 'ai');

-- 2. Insert Fighters
INSERT INTO fighters (first_name, last_name, nickname, weight_class) VALUES
('Rafael', 'Fiziev', 'Ataman', 'Lightweight'),
('Justin', 'Gaethje', 'The Highlight', 'Lightweight'),
('Islam', 'Makhachev', NULL, 'Lightweight'),
('Arman', 'Tsarukyan', 'Ahalkalakets', 'Lightweight'),
('Alexander', 'Volkanovski', 'The Great', 'Featherweight'),
('Diego', 'Lopes', NULL, 'Featherweight'),
('Shavkat', 'Rakhmonov', 'Nomad', 'Welterweight'),
('Belal', 'Muhammad', 'Remember the Name', 'Welterweight');

-- 3. Insert Events
INSERT INTO events (title, event_date, location) VALUES
('UFC 300', '2024-04-13 22:00:00+00', 'Las Vegas, NV, USA'),
('UFC 311', '2025-01-18 22:00:00+00', 'Los Angeles, CA, USA');

-- 4. Insert Fights
-- Fight 1: Gaethje vs Fiziev (Completed)
INSERT INTO fights (
    event_id, fighter_a_id, fighter_b_id, weight_class, 
    is_title_fight, is_main_event, odds_fighter_a, odds_fighter_b, 
    winner_id, win_method, win_round, win_time, status
) VALUES (
    1, 2, 1, 'Lightweight', 
    FALSE, FALSE, 2.10, 1.75, 
    2, 'Decision', 3, '5:00', 'completed'
);

-- Fight 2: Makhachev vs Tsarukyan (Scheduled)
INSERT INTO fights (
    event_id, fighter_a_id, fighter_b_id, weight_class, 
    is_title_fight, is_main_event, odds_fighter_a, odds_fighter_b, status
) VALUES (
    2, 3, 4, 'Lightweight', 
    TRUE, TRUE, 1.35, 3.25, 'scheduled'
);

-- Fight 3: Muhammad vs Rakhmonov (Scheduled)
INSERT INTO fights (
    event_id, fighter_a_id, fighter_b_id, weight_class, 
    is_title_fight, is_main_event, odds_fighter_a, odds_fighter_b, status
) VALUES (
    2, 8, 7, 'Welterweight', 
    TRUE, FALSE, 2.90, 1.42, 'scheduled'
);

-- 5. Insert Predictions
-- Predictions for Fight 1 (Gaethje vs Fiziev)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability, predicted_method, reasoning) VALUES
(1, 1, 1, 0.60, 'Decision', 'Fiziev will outstrike Gaethje speedwise over 3 rounds.'),
(1, 2, 2, 0.55, 'KO/TKO', 'Gaethje has superior leg kicks and attrition power in late rounds.'),
(1, 3, 2, 0.65, 'Decision', 'Gaethje`s takedown defense will force a standup war where he edges volume.'),
(1, 4, 1, 0.52, 'Decision', 'Fiziev technique in Muay Thai range gives him early round leverage.');

-- Predictions for Fight 2 (Makhachev vs Tsarukyan)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability, predicted_method, reasoning) VALUES
(2, 1, 3, 0.70, 'Submission', 'Makhachev controls position and will find a choke in championship rounds.'),
(2, 2, 3, 0.65, 'Decision', 'Tsarukyan will scramble well but Makhachev accumulates top control time.'),
(2, 3, 4, 0.55, 'Decision', 'Tsarukyan`s wrestling pace can disrupt Islam`s rhythm over 5 rounds.'),
(2, 4, 3, 0.60, 'Submission', 'Islam chain-wrestling advantage will yield a back-take submission.');