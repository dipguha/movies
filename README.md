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
