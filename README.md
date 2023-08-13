## eFilm-API Docummentation with Swagger OpenAPI

- Swagger API Specification <br>
  <a href="https://efilm-restfulapi.dimasfebriyant1.repl.co" target="_blank">https://efilm-restfulapi.dimasfebriyant1.repl.co</a>

- API End Point <br>
  <a href="https://efilm-api-project.fly.dev/" target="_blank">https://efilm-api-project.fly.dev</a>

# How to using with your application

<h3>samples</h3>

- with Go using <a href="https://pkg.go.dev/net/http" target="_blank">net/http</a>

  ```go
  resp, err := http.Get("https://efilm-api-project.fly.dev/api/movies/1")
  if err != nil {
      return err
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  // ...
  ```

- with javascript using <a href="https://www.npmjs.com/package/axios" target="_blank">axios</a>

  ```javascript
  axios
    .get("https://efilm-api-project.fly.dev/api/movies/1")
    .then(function (response) {
      // handle success
      console.log(response);
    })
    .catch(function (error) {
      // handle error
      console.log(error);
    })
    .finally(function () {
      // always executed
    });
  ```

- and others languages

  please check your library documentation about http request
