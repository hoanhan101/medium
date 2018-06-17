# medium

**medium** is a social media web application inspired by
[GopherFace](https://github.com/EngineerKamesh/gofullstack). The goal is to
build a simple, powerful, production-ready distributed system that follows good
practices of isomorphic application architecture.

**References:**
- [Web Programming with Go by Kamesh
  Balasubramanian](https://github.com/EngineerKamesh/gofullstack)
- [Isomorphic Go](https://www.packtpub.com/web-development/isomorphic-go)

## Project Status

It is an on-going project. Below are the main tasks and goals.

**Back-End Phase:**

- [x] Build a Web Sever instance.
- [x] Introduce custom Templates.
- [x] Follow Model View Controller architecture.
- [x] Chain logging middleware and panic recovery middleware.
- [x] Make RESTful API for CRUD operations.
- [x] Build a Web Form with validation and security kit.
- [x] Introduce File Uploads.
- [x] Implement datastores: MySQL, MongoDB, and Redis.
- [x] Use secure cookie and persist user session.
- [x] Introduce OpenSSL and add authentication middleware.
- [x] Implement an Asynchronous Task Queue for resizing images task.

**Front-End Phase:**
- [x] Introduce GopherJS to implement Client-Side routing and templates
  rendering.
- [x] Adapt Single Page Architecture.
- [ ] Implement core features and functionalities: user's profile, image 
  upload, friends list, news, feed,...

**Production Phase:**
- [ ] Introduce Nginx.
- [ ] Install a Let’s Encrypt SSL Certificate.
- [ ] Cross compile application.
- [ ] Deploy as a standalone server.
- [ ] Run the application as system service.
- [ ] Deploy the application as a multi-container Docker application.
