BEGIN;

INSERT INTO users (username, password, createdAt, email, firstName, lastName) VALUES
    ('Vladik123', '$c23adw.vGA1O9wmRjrwAVXD98HNOgsNpDczlqm3Jq7KnEd1rVAGv3Fykk1a', '2010-01-19 07:14:07', 'vlad@ua.com', 'vlad', 'vrek'),
    ('NastyGRL', '$c23adw.vGA1O9wmRjrwAVXD98HNOgsNpDczlqm3Jq7KnEd1rVAGv3Fykk1a', '2010-01-19 07:14:07', 'aefdqwd@sw.com', 'Lina', 'Q'),
    ('charby', '$c23adw.vGA1O9wmRjrwAVXD98HNOgsNpDczlqm3Jq7KnEd1rVAGv3Fykk1a', '2010-01-19 07:14:07', 'char777@com.com', 'char', 'sunny');

INSERT INTO tweets (id, username, createdAt, text) VALUES
    (1, 'Vladik123', '2012-01-19 07:14:07', 'What a wonderful day!'),
    (2, 'Vladik123', '2012-01-19 07:15:00', 'Oh no! Someone is incorrect on the internet!'),
    (3, 'Vladik123', '2012-01-10 12:10:00', 'What a wonderful day!'),
    (4, 'Vladik123', '2013-11-30 11:11:00', 'What a wonderful day!'),
    (5, 'Vladik123', '2011-07-10 12:10:00', 'What a wonderful day!'),
    (6, 'Vladik123', '2014-05-15 12:10:00', 'What a wonderful day!'),
    (7, 'Vladik123', '2014-05-16 12:10:00', 'What a wonderful day!'),
    (8, 'Vladik123', '2014-05-17 12:10:00', 'What a wonderful day!'),
    (9, 'Vladik123', '2014-05-18 12:10:00', 'What a wonderful day!'),
    (10, 'Vladik123', '2014-05-19 12:15:00', 'What a wonderful day!'),
    (11, 'charby', '2016-04-26 07:10:00', 'Just sold bitcoin for 225$, don\'t see any investment perspectives :( #bitcoin');

INSERT INTO followers (follower, publisher) VALUES
   ('Vladik123', 'charby'),
   ('NastyGRL', 'charby');

COMMIT;