-- Table authors
INSERT INTO authors (id, name, country, profile_image_url) VALUES
                                                               ('8e21a1ff-4cd7-4c7c-9394-2c35a7e7a1a2', 'Jane Doe', 'France', 'https://example.com/images/jane_doe.jpg'),
                                                               ('e12a3aff-5da4-4cd7-8895-2f9b6a9e8c1b', 'John Smith', 'USA', 'https://example.com/images/john_smith.jpg');

-- Table statuses
INSERT INTO statuses (id, name) VALUES
                                    ('c1b9b2d1-3f4b-4c5f-9a1c-3e9d6a8f2b78', 'Draft'),
                                    ('d2a9c3d4-6a3f-7c8d-9b1e-4e7f5a9b6e2c', 'Published');

-- Table visibilities
INSERT INTO visibilities (id, name) VALUES
                                        ('a1c3b2d4-3f6a-7c5f-9a1c-4e7f5a9b6e2c', 'Public'),
                                        ('b2d4c3a6-5f7a-9c1e-3f4b-7a8e6c9b2d1f', 'Private');

-- Table articles
INSERT INTO articles (id, type, cover_image_url, cover_text, date, url, author_id, profile_image_url, estimate_time, meta_title, meta_description, created_at, updated_at, status_id, visibility_id) VALUES
                                                                                                                                                                                                         ('3d2f1e9b-7c8d-4a6f-b1c3-5e9a7f6d8b78', 'Article', 'https://example.com/images/article1.jpg', 'Introduction to Go', '2024-12-01 10:00:00', 'https://example.com/article1', '8e21a1ff-4cd7-4c7c-9394-2c35a7e7a1a2', 'https://example.com/profiles/jane_doe.jpg', 5, 'Go Basics', 'Learn the basics of Go programming.', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'd2a9c3d4-6a3f-7c8d-9b1e-4e7f5a9b6e2c', 'a1c3b2d4-3f6a-7c5f-9a1c-4e7f5a9b6e2c'),
                                                                                                                                                                                                         ('5f7a6b4d-9c8e-3b2c-4a1e-6f7a9b5d3c8f', 'Article', 'https://example.com/images/article2.jpg', 'Advanced Go Patterns', '2024-12-02 12:00:00', 'https://example.com/article2', 'e12a3aff-5da4-4cd7-8895-2f9b6a9e8c1b', 'https://example.com/profiles/john_smith.jpg', 10, 'Go Patterns', 'Advanced patterns in Go programming.', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'd2a9c3d4-6a3f-7c8d-9b1e-4e7f5a9b6e2c', 'b2d4c3a6-5f7a-9c1e-3f4b-7a8e6c9b2d1f');

-- Table technologies
INSERT INTO technologies (id, name, logo_url, category) VALUES
                                                            ('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d', 'Go', 'https://example.com/images/go.png', 'Programming Language'),
                                                            ('2b3c4d5e-6f7a-8b9c-0d1e-3f4a5b6c7d8e', 'Docker', 'https://example.com/images/docker.png', 'DevOps');

-- Table article_technologies
INSERT INTO article_technologies (article_id, technology_id) VALUES
                                                                 ('3d2f1e9b-7c8d-4a6f-b1c3-5e9a7f6d8b78', '1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d'),
                                                                 ('5f7a6b4d-9c8e-3b2c-4a1e-6f7a9b5d3c8f', '2b3c4d5e-6f7a-8b9c-0d1e-3f4a5b6c7d8e');

-- Table chapters
INSERT INTO chapters (id, article_id, title) VALUES
                                                 ('6a7b8c9d-0e1f-2a3b-4c5d-6e7f8a9b0c1d', '3d2f1e9b-7c8d-4a6f-b1c3-5e9a7f6d8b78', 'Getting Started with Go'),
                                                 ('7b8c9d0e-1f2a-3b4c-5d6e-7f8a9b0c1d2e', '5f7a6b4d-9c8e-3b2c-4a1e-6f7a9b5d3c8f', 'Deep Dive into Go Patterns');

-- Table contents
INSERT INTO contents (id, chapter_id, type, value, language, media_type, src, alt_text) VALUES
                                                                                            ('8c9d0e1f-2a3b-4c5d-6e7f-8a9b0c1d2e3f', '6a7b8c9d-0e1f-2a3b-4c5d-6e7f8a9b0c1d', 'text', 'Welcome to Go programming!', 'English', NULL, NULL, NULL),
                                                                                            ('9d0e1f2a-3b4c-5d6e-7f8a-9b0c1d2e3f4a', '7b8c9d0e-1f2a-3b4c-5d6e-7f8a9b0c1d2e', 'code', 'package main\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}', 'English', 'code', NULL, NULL);

-- Table related_articles
INSERT INTO related_articles (article_id, related_article_id) VALUES
                                                                  ('3d2f1e9b-7c8d-4a6f-b1c3-5e9a7f6d8b78', '5f7a6b4d-9c8e-3b2c-4a1e-6f7a9b5d3c8f'),
                                                                  ('5f7a6b4d-9c8e-3b2c-4a1e-6f7a9b5d3c8f', '3d2f1e9b-7c8d-4a6f-b1c3-5e9a7f6d8b78');
