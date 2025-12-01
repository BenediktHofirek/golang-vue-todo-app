-- User A ID: a1b2c3d4-e5f6-7890-1234-567890abcdef
-- User B ID: f0e9d8c7-b6a5-4321-fedc-ba9876543210

--
-- INSERT 10 SAMPLE TODO RECORDS
--
INSERT INTO todos (user_id, title, description, completed, due_date)
SELECT *
FROM (
    VALUES
    -- USER A TASKS
    ('a1b2c3d4-e5f6-7890-1234-567890abcdef',
     'Review Q4 Budget Forecast',
     'Go through the forecast spreadsheet and ensure all departmental budgets are accounted for.',
     FALSE,
     CURRENT_TIMESTAMP - INTERVAL '5 days 8 hours'),
     
    ('a1b2c3d4-e5f6-7890-1234-567890abcdef',
     'Send Standup Notes to Team',
     'Compile action items and decisions from the morning standup meeting.',
     FALSE,
     CURRENT_TIMESTAMP + INTERVAL '3 hours'),

    ('a1b2c3d4-e5f6-7890-1234-567890abcdef',
     'Research PostgreSQL Indexing Strategies',
     'Read up on B-tree vs. Hash vs. GiST indices for optimizing the "products" table.',
     TRUE,
     CURRENT_TIMESTAMP - INTERVAL '1 day'),

    ('a1b2c3d4-e5f6-7890-1234-567890abcdef',
     'Deploy Critical Hotfix 1.2.1',
     'The server is experiencing high latency. Must deploy the hotfix immediately.',
     FALSE,
     CURRENT_TIMESTAMP - INTERVAL '12 hours'),

    ('a1b2c3d4-e5f6-7890-1234-567890abcdef',
     'Write Draft Blog Post on New API',
     'The initial draft of the technical blog post is complete.',
     TRUE,
     CURRENT_TIMESTAMP + INTERVAL '2 days'),

    -- USER B TASKS
    ('f0e9d8c7-b6a5-4321-fedc-ba9876543210',
     'File 2024 Tax Return',
     'Upload all necessary documents and submit the final tax filing before the deadline.',
     TRUE,
     CURRENT_TIMESTAMP - INTERVAL '3 weeks'),

    ('f0e9d8c7-b6a5-4321-fedc-ba9876543210',
     'Plan Project Kickoff Agenda',
     'Outline the structure and key discussion points for the new marketing campaign kickoff.',
     FALSE,
     CURRENT_TIMESTAMP + INTERVAL '7 days'),

    ('f0e9d8c7-b6a5-4321-fedc-ba9876543210',
     'Book Summer Vacation Flights',
     'Search and book round-trip tickets to Kyoto for the family trip in July.',
     FALSE,
     CURRENT_TIMESTAMP + INTERVAL '90 days'),

    ('f0e9d8c7-b6a5-4321-fedc-ba9876543210',
     'Buy Groceries',
     'Milk, eggs, bread, and chicken for dinner.',
     FALSE,
     CURRENT_TIMESTAMP + INTERVAL '1 day 4 hours'),

    ('f0e9d8c7-b6a5-4321-fedc-ba9876543210',
     'Conduct Mid-Project Review',
     'Review progress with the dev team and adjust timeline if necessary.',
     FALSE,
     CURRENT_TIMESTAMP + INTERVAL '14 days')
) AS t(user_id, title, description, completed, due_date)
WHERE NOT EXISTS (SELECT 1 FROM todos);
