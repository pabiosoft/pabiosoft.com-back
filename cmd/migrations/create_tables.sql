-- Table authors
CREATE TABLE IF NOT EXISTS authors (
                                       id UUID PRIMARY KEY,
                                       name VARCHAR(255) NOT NULL,
    country VARCHAR(255),
    profile_image_url TEXT
    );

-- Table statuses
CREATE TABLE IF NOT EXISTS statuses (
                                        id UUID PRIMARY KEY,
                                        name VARCHAR(255) NOT NULL
    );

-- Table visibilities
CREATE TABLE IF NOT EXISTS visibilities (
                                            id UUID PRIMARY KEY,
                                            name VARCHAR(255) NOT NULL
    );

-- Table articles
CREATE TABLE IF NOT EXISTS articles (
                                        id UUID PRIMARY KEY,
                                        type VARCHAR(255),
    cover_image_url TEXT,
    cover_text TEXT,
    date TIMESTAMP,
    url TEXT,
    author_id UUID REFERENCES authors (id) ON DELETE SET NULL,
    profile_image_url TEXT,
    estimate_time INT,
    meta_title VARCHAR(255),
    meta_description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_id UUID REFERENCES statuses (id) ON DELETE SET NULL,
    visibility_id UUID REFERENCES visibilities (id) ON DELETE SET NULL
    );

-- Table technologies
CREATE TABLE IF NOT EXISTS technologies (
                                            id UUID PRIMARY KEY,
                                            name VARCHAR(255) NOT NULL,
    logo_url TEXT,
    category VARCHAR(255)
    );

-- Table article_technologies
CREATE TABLE IF NOT EXISTS article_technologies (
                                                    article_id UUID REFERENCES articles (id) ON DELETE CASCADE,
    technology_id UUID REFERENCES technologies (id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, technology_id)
    );

-- Table chapters
CREATE TABLE IF NOT EXISTS chapters (
                                        id UUID PRIMARY KEY,
                                        article_id UUID REFERENCES articles (id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL
    );

-- Table contents
CREATE TABLE IF NOT EXISTS contents (
                                        id UUID PRIMARY KEY,
                                        chapter_id UUID REFERENCES chapters (id) ON DELETE CASCADE,
    type VARCHAR(255),
    value TEXT,
    language VARCHAR(255),
    media_type VARCHAR(255),
    src TEXT,
    alt_text TEXT
    );

-- Table related_articles
CREATE TABLE IF NOT EXISTS related_articles (
                                                article_id UUID REFERENCES articles (id) ON DELETE CASCADE,
    related_article_id UUID REFERENCES articles (id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, related_article_id)
    );

