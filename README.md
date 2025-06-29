# 📝 go-blog

A lightweight blogging platform with Markdown support, clean RESTful API, and htmx frontend.

This project demonstrates a clear Go backend architecture suitable for learning and further extension.

---

## 🎯 Project Goals

✅ Clean and transparent backend structure  
✅ Separation of concerns (Clean Architecture principles)  
✅ RESTful API for posts, tags, and users  
✅ Search by tags and full-text search  
✅ Simple frontend with htmx  
✅ Easy to configure and run

---

## ⚙️ Technical Stack

| Component           | Technology                                         |
|---------------------|-----------------------------------------------------|
| **Language**        | Pure Go + htmx                                      |
| **HTTP Router**     | [Gin](https://github.com/gin-gonic/gin)             |
| **Configuration**   | [cleanenv](https://github.com/ilyakaznacheev/cleanenv) |
| **Database**        | PostgreSQL                                          |
| **DB Driver**       | [pgx](https://github.com/jackc/pgx) (with stdlib + sqlx) |
| **Authentication**  | [JWT](https://github.com/golang-jwt/jwt/v5)         |
| **Frontend**        | [htmx](https://htmx.org/)                           |
| **Markdown**        | [gomarkdown/markdown](https://github.com/gomarkdown/markdown) |
| **Migrations**      | [golang-migrate](https://github.com/golang-migrate/migrate) or [goose](https://github.com/pressly/goose) |

---