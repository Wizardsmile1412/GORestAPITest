Welcome to my goofy Go playground where Gin routers, SQLite tables, and JWTs throw a tiny party together. 🎉

## 🚀 Quick Start
- Make sure Go 1.24+ is on your PATH.
- From the repo root run go run . and the server pops up on http://localhost:8080.
- A local pi.db file is created on the fly with tables for users, events, and registrations.

## 🧩 What You Can Poke
- Public GET /events (all) and GET /events/:id (single) endpoints.
- JWT-protected create/update/delete on /events plus /events/:id/register for RSVPs.
- /signup to stash a user (bcrypt-salted), /login to grab a short-lived token.

## 🕹️ Playground Tips
- The /middlewares/Auth check expects the raw token inside the Authorization header (no "Bearer" prefix needed).
- Sample REST requests live under pi-test/*.http; copy-paste them into an HTTP client to blast the API.
- Forget cleanup? DELETE /events/:id and DELETE /events/:id/register are your digital broom. 🧹

Have fun breaking (and fixing) things! 😄
