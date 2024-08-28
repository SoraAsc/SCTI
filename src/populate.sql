-- Clear existing data
DROP TABLE IF EXISTS activities
TRUNCATE TABLE registrations, passwd, users, activities, day RESTART IDENTITY CASCADE;

-- Insert days
INSERT INTO day (day) VALUES (1), (2), (3), (4), (5);

-- Insert activities
INSERT INTO activities (spots, activity_type, room, speaker, topic, description, time, day) 
VALUES
  (2, 'MC', 'INF-1', 'Annabel', 'Atividade exclusiva de duas vagas', 'Atividade para testar erro de atividade cheia', '08:00:00', 1),
  (20, 'MC', 'INF-2', 'Zoey', 'Elixir em ação', 'Como o elixir é usado em situações reais', '08:00:00', 1),
  (20, 'PL', 'Auditório 1', 'Guto', 'Trabalhando fora', 'Como conseguir trabalho fora e os passos que você deve tomar', '16:00:00', 1),
  (20, 'MC', 'INF-1', 'Emanuel', 'Nix e Sistemas Reproduziveis', 'Garantindo a funcionalidade de um projeto em diversos sistemas', '08:00:00', 2),
  (20, 'MC', 'INF-2', 'Sophia', 'Godot e Game Development', 'Introdução para game development usando a ferramenta Godot', '08:00:00', 2),
  (20, 'PL', 'Auditório 1', 'Vasilis', 'Trabalhando como contribuidor Open Source', 'Os altos e baixos de ser um contribuidor', '16:00:00', 2),
  (20, 'MC', 'INF-1', 'Filipe', 'Java para web devs', 'Usando java para desenvolver uma aplicação web', '08:00:00', 3),
  (20, 'PL', 'Auditório 1', 'May', 'Linguagens velhas em um mundo novo', 'Por que linguagens antigas commo PHP e Cobol ainda são usadas', '16:00:00', 3),
  (20, 'MC', 'INF-2', 'Luis Ramirez', 'Treinando IAs', 'Como usar Pytorch e Numpy prara treinar sua própria IA', '08:00:00', 4),
  (20, 'MC', 'INF-2', 'Riveira', 'Introdução a Engines', 'Como uma engine funciona e como criar a sua', '08:00:00', 5);


-- Verify activities
-- SELECT topic, spots
-- FROM activities;

-- Insert users
INSERT INTO users (email, name, uuid, verificationCode, isVerified, isAdmin, isPaid)
VALUES
  ('usuario1@teste.com', 'Usuário Um', '123e4567-e89b-12d3-a456-426614174000', '123e4', TRUE, FALSE, FALSE),
  ('usuario2@teste.com', 'Usuário Dois', '223e4567-e89b-12d3-a456-426614174000', '223e4', TRUE, FALSE, TRUE),
  ('usuario3@teste.com', 'Usuário Três', '323e4567-e89b-12d3-a456-426614174000', '323e4', FALSE, FALSE, FALSE),
  ('admin@teste.com', 'Administrador', '423e4567-e89b-12d3-a456-426614174000', '423e4', TRUE, TRUE, TRUE),
  ('usuario5@teste.com', 'Usuário Cinco', '523e4567-e89b-12d3-a456-426614174000', '523e4', TRUE, FALSE, TRUE);

-- Insert passwords
INSERT INTO passwd (id, passwd)
SELECT id, CONCAT('senha_segura_', id)
FROM users;

-- Insert registrations for the 2-spot activity
INSERT INTO registrations (user_id, activity_id)
VALUES 
  ('123e4567-e89b-12d3-a456-426614174000', 1),  -- Usuário Um
  ('223e4567-e89b-12d3-a456-426614174000', 1);  -- Usuário Dois

-- Manually adjust spots for the exclusive activity
UPDATE activities
SET spots = spots - (SELECT COUNT(*) FROM registrations WHERE activity_id = 1)
WHERE id = 1;

-- Insert other registrations
INSERT INTO registrations (user_id, activity_id)
VALUES 
  ('323e4567-e89b-12d3-a456-426614174000', 2),  -- Usuário Três - Elixir em ação
  ('423e4567-e89b-12d3-a456-426614174000', 3),  -- Administrador - Trabalhando fora
  ('523e4567-e89b-12d3-a456-426614174000', 4);  -- Usuário Cinco - Nix e Sistemas Reproduziveis

UPDATE activities
SET spots = spots - 1
WHERE id = 2;

UPDATE activities
SET spots = spots - 1
WHERE id = 3;

UPDATE activities
SET spots = spots - 1
WHERE id = 4;

-- SELECT topic, spots
-- FROM activities;

