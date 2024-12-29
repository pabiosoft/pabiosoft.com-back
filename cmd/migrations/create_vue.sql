CREATE OR REPLACE VIEW article_details_view AS
SELECT
    a.id AS article_id,
    a.type AS article_type,
    a.cover_image_url,
    a.cover_text,
    a.date,
    a.url,
    a.profile_image_url,
    a.estimate_time,
    a.meta_title,
    a.meta_description,
    a.created_at,
    a.updated_at,
    au.id AS author_id,
    au.name AS author_name,
    au.country AS author_country,
    au.profile_image_url AS author_profile_image_url,
    s.id AS status_id,
    s.name AS status_name,
    v.id AS visibility_id,
    v.name AS visibility_name,
    -- Group chapters with their contents as JSON
    (
        SELECT JSON_ARRAYAGG(
                       JSON_OBJECT(
                               'id', c.id,
                               'title', c.title,
                               'content', (
                                   SELECT JSON_ARRAYAGG(
                                                  JSON_OBJECT(
                                                          'type', ct.type,
                                                          'value', ct.value,
                                                          'language', ct.language,
                                                          'mediaType', ct.media_type,
                                                          'src', ct.src,
                                                          'altText', ct.alt_text
                                                  )
                                          )
                                   FROM contents ct
                                   WHERE ct.chapter_id = c.id
                               )
                       )
               )
        FROM chapters c
        WHERE c.article_id = a.id
    ) AS chapters,
    -- Group technologies as JSON
    (
        SELECT JSON_ARRAYAGG(
                       JSON_OBJECT(
                               'id', t.id,
                               'name', t.name,
                               'logoUrl', t.logo_url,
                               'category', t.category
                       )
               )
        FROM article_technologies at
        INNER JOIN technologies t ON at.technology_id = t.id
WHERE at.article_id = a.id
    ) AS technologies,
-- Group related articles as JSON
    (
SELECT JSON_ARRAYAGG(
    JSON_OBJECT(
    'id', ra.related_article_id,
    'coverImageUrl', r.cover_image_url,
    'author', JSON_OBJECT(
    'id', rau.id,
    'name', rau.name,
    'country', rau.country,
    'profileImageUrl', rau.profile_image_url
    )
    )
    )
FROM related_articles ra
    INNER JOIN articles r ON ra.related_article_id = r.id
    LEFT JOIN authors rau ON r.author_id = rau.id
WHERE ra.article_id = a.id
    ) AS relatedArticles
FROM articles a
    LEFT JOIN authors au ON a.author_id = au.id
    LEFT JOIN statuses s ON a.status_id = s.id
    LEFT JOIN visibilities v ON a.visibility_id = v.id;


----- TEST
SELECT *
FROM article_details_view
WHERE article_id = '1e184480-f0ac-452c-a88e-86f18f53708e';
