-- Clear tables before populating (optional, if restarting the seed)
TRUNCATE TABLE predictions, fights, events, fighters, predictors RESTART IDENTITY CASCADE;

-- 1. Adding predictors (sources of forecasts)
INSERT INTO predictors (name, type) VALUES
('me', 'human'),
('GPT-4o', 'ai'),
('DeepSeek-V3', 'ai'),
('Claude-3.5-Sonnet', 'ai');

-- 2. Добавляем бойцов
INSERT INTO fighters (first_name, last_name, nickname, weight_class) VALUES
('Rafael', 'Fiziev', 'Ataman', 'Lightweight'),
('Justin', 'Gaethje', 'The Highlight', 'Lightweight'),
('Islam', 'Makhachev', NULL, 'Lightweight'),
('Arman', 'Tsarukyan', 'Ahalkalakets', 'Lightweight'),
('Alexander', 'Volkanovski', 'The Great', 'Featherweight'),
('Diego', 'Lopes', NULL, 'Featherweight'),
('Shavkat', 'Rakhmonov', 'Nomad', 'Welterweight'),
('Belal', 'Muhammad', 'Remember the Name', 'Welterweight');

-- 3. Adding fighters
INSERT INTO events (title, event_date, location) VALUES
('UFC 300', '2024-04-13 22:00:00+00', 'Las Vegas, NV, USA'),
('UFC 311', '2025-01-18 22:00:00+00', 'Los Angeles, CA, USA');

-- 4. Adding fights
-- Fight 1 (Completed): Justin Gaethje vs. Rafael Fiziev (Winner: Gaethje)
INSERT INTO fights (event_id, fighter_a_id, fighter_b_id, weight_class, winner_id, win_method, win_round, status) VALUES
(1, 2, 1, 'Lightweight', 2, 'Decision', 3, 'completed');

-- Fight 2 (Upcoming): Islam Makhachev vs Arman Tsarukyan
INSERT INTO fights (event_id, fighter_a_id, fighter_b_id, weight_class, status) VALUES
(2, 3, 4, 'Lightweight', 'scheduled');

-- Fight 3 (Upcoming): Belal Muhammad vs. Shavkat Rakhmonov
INSERT INTO fights (event_id, fighter_a_id, fighter_b_id, weight_class, status) VALUES
(2, 8, 7, 'Welterweight', 'scheduled');


-- 5. Adding predictions
-- Predictions for Fight 1 (Gaethje vs Fiziev)
-- Predictor 1 (me) bet on Fiziev (ID 1) with a probability of 0.60 (lost)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 1, 1, 0.60);

-- GPT-4o bet on Gaethje (ID 2) with a probability of 0.55 (won)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 2, 2, 0.55);

-- DeepSeek-V3 bet on Gaethje (ID 2) with a probability of 0.65 (won)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 3, 2, 0.65);

-- Claude-3.5 bet on Fiziev (ID 1) with a probability of 0.52 (lost)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 4, 1, 0.52);

-- Predictions for Fight 2 (Makhachev vs. Tsarukyan)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(2, 1, 3, 0.70), -- me -> Makhachev (70%)
(2, 2, 3, 0.65), -- GPT-4o -> Makhachev (65%)
(2, 3, 4, 0.55), -- DeepSeek-V3 -> Tsarukyan (55%)
(2, 4, 3, 0.60); -- Claude-3.5 -> Makhachev (60%)