```shell
docker exec -it go_mariadb bash
apt install -y mariadb-client

mysql -u root -p

 ```

## POST CREATE ARTICLE
```shell
{
    "@type": "Book",
    "coverImageUrl": "https://example.com/images/article1.jpg",
    "coverText": "Hey drake to Go test insert 1",
    "profileImageUrl": "https://example.com/profiles/jane_doe.jpg",
    "date": "2024-12-01T10:00:00Z",
    "url": "https://example.com/article1",
    "author": {
        "@id": "8e21a1ff-4cd7-4c7c-9394-2c35a7e7a1a2"
    },
    "chapters": [
        {
            "title": "Bonsoir chapitre 1 Started with Go",
            "content": [
                {
                    "type": "text",
                    "value": "Welcome to Go programming!",
                    "language": "Go",
                    "mediaType": "text/plain",
                    "src": "",
                    "altText": ""
                }
            ]
        },
         {
            "title": "Bonsoir chapitre 2 Started with Go",
            "content": [
                {
                    "type": "code",
                    "value": "\n echo 'bonjour php '; ",
                    "language": "Go",
                    "mediaType": "text/plain",
                    "src": "",
                    "altText": ""
                }
            ]
        }
    ],
    "technologies": [
        {
            "id": "1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d"
        }
    ],
    "relatedArticles": [
        {
            "@id": "5f7a6b4d-9c8e-3b2c-4a1e-6f7a9b5d3c8f"
        }
    ],
    "estimateTime": 5,
    "metaTitle": "Go Basics",
    "metaDescription": "Learn the basics of Go programming.",
    "createdAt": "2024-12-28T19:03:15Z",
    "updatedAt": "2024-12-28T19:03:15Z",
    "status": {
        "id": "d2a9c3d4-6a3f-7c8d-9b1e-4e7f5a9b6e2c"
     },
    "visibility": {
        "id": "a1c3b2d4-3f6a-7c5f-9a1c-4e7f5a9b6e2c"
        }
}
```