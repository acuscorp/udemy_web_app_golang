### Templates

### structure used ion this project
use command to run the app:
 go run cmd/web/main.go

-cmd/
    -web/
        -main.go
-pkg/
    -handlers/
        -handlers.go
    -render/
        -render.go
-templates/
    -home.page.tmpl
    -about.page.tmpl
-go.mod

### Layouts

is used to avoid duplication of work writing manually

key words used are **template**, **define**, **block**