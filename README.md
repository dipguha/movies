# movies
Movies website using React and Go

# All software needed for this website
1. Install Node
    1. node --version (v16.15.0)
2. Install Go
    1. go version (go1.21.5 darwin/arm64)
3. Install VSC
    1. Extension - ES7+ React/Redux/React-Native snippets v4.4.3
    2. Extension - Go v0.40.1
    3. cmd + shift + p -> Go: Install/Update Tools
4. Install Docker
    1. Docker desktop

# Create React app
1. npx create-react-app my-app
2. cd my-app
3. npm start
4. index.html - remove all comments - public folder
5. src - index.js
6. Update npm version - npm install -g npm@10.2.5

# Create Go Movies apps
1. npx create-react-app go-movies-front-end (/Documents/GitHub/personal/movies)
2. Run the react app
    1. cd go-movies-front-end
    2. npm start
3. cd /Documents/GitHub/personal/movies/go-movies-front-end
    1. code .
4. Remove App.css, App.test.js, log.svg, reportWebVitals.js, setupTests.js
5. Clean up index.js
6. App.js clean up
7. cd go-movies-front-end
    1. npm start
    2. Local: http://localhost:3000, Network: http://192.168.1.226:3000
8. Add bootstrap
    1. https://getbootstrap.com/
    2. Include via CDN - copy Link part - <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
    3. index.html
9. Install React router
    1. npm i react-router-dom (go-movies-front-end)
10. React knowledge
    1. useState- hook to manage state in functional components
    2. useEffect - hook to manage side effect like fetching data
    3. Functinal component vs Class components 
    4. JSX - 
    5. Javascript expression and template
11. Go 
    1. Create new folder
    2. go mod init backend
12. Run go 
    1. go run ./cmd/api
    2. http://localhost:8080/
13. Install router
    1. chi router
    2. go get -u github.com/go-chi/chi/v5
    3. Port - sudo lsof -i :8080
14. Install DB using Docker
    1. docker-compose up -d
    2. docker-compose down
    3. Beekeeper studio to connect to the Postgres DB
    4. Connect using TablePlus (Name - movies, Host - localhost, User - postgres, pwd - postgres, Database - movies, SSL mode - PREFERRED)
15. Git commands
    1. git rm -r --cached go-movies-back-end/postgres-data
    2. git commit -m "Removed folder_name"    
16. Database driver, connecting to Postgres
    1. lib/pq
    2. pgx - used, 
    3. go get github.com/jackc/pgx/v4
    4. go get github.com/jackc/pgconn
17. JWT package for User authentication
    1. golang-jwt/jwt - https://github.com/golang-jwt/jwt
    2. go get github.com/golang-jwt/jwt/v4 
        1. go: downloading github.com/golang-jwt/jwt/v4 v4.5.0
        2. go: downloading github.com/golang-jwt/jwt v3.2.2+incompatible
        3. go: added github.com/golang-jwt/jwt/v4 v4.5.0
18. 
