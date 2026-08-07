-- Очищаем таблицы перед заполнением (опционально, если перезапускаем сид)
TRUNCATE TABLE predictions, fights, events, fighters, predictors RESTART IDENTITY CASCADE;

-- 1. Добавляем предикторов (источники прогнозов)
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

-- 3. Добавляем турниры
INSERT INTO events (title, event_date, location) VALUES
('UFC 300', '2024-04-13 22:00:00+00', 'Las Vegas, NV, USA'),
('UFC 311', '2025-01-18 22:00:00+00', 'Los Angeles, CA, USA');

-- 4. Добавляем бои
-- Бой 1 (Завершён): Justin Gaethje vs Rafael Fiziev (Победитель: Gaethje)
INSERT INTO fights (event_id, fighter_a_id, fighter_b_id, weight_class, winner_id, win_method, win_round, status) VALUES
(1, 2, 1, 'Lightweight', 2, 'Decision', 3, 'completed');

-- Бой 2 (Предстоящий): Islam Makhachev vs Arman Tsarukyan
INSERT INTO fights (event_id, fighter_a_id, fighter_b_id, weight_class, status) VALUES
(2, 3, 4, 'Lightweight', 'scheduled');

-- Бой 3 (Предстоящий): Belal Muhammad vs Shavkat Rakhmonov
INSERT INTO fights (event_id, fighter_a_id, fighter_b_id, weight_class, status) VALUES
(2, 8, 7, 'Welterweight', 'scheduled');


-- 5. Добавляем прогнозы
-- Прогнозы на Бой 1 (Gaethje vs Fiziev)
-- Предиктор 1 (me) поставил на Fiziev (ID 1) с вероятностью 0.60 (проиграл)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 1, 1, 0.60);

-- GPT-4o поставил на Gaethje (ID 2) с вероятностью 0.55 (выиграл)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 2, 2, 0.55);

-- DeepSeek-V3 поставил на Gaethje (ID 2) с вероятностью 0.65 (выиграл)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 3, 2, 0.65);

-- Claude-3.5 поставил на Fiziev (ID 1) с вероятностью 0.52 (проиграл)
INSERT INTO predictions (fight_id, predictor_id, predicted_winner_id, probability) VALUES
(1, 4, 1, 0.52);

