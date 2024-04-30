package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type PageVariables struct {
    Title string
}

func main() {
    http.HandleFunc("/", HomePage)
    fmt.Println("Fut itt most http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    pageVariables := PageVariables{
        Title: "Istenem ember",
    }

    tmpl := template.Must(template.New("form").Parse(`
    <html>
    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        <form id="numberForm">
            <input type="number" id="input1" required="required" placeholder="Irj be egy szamot"/>
            <input type="number" id="input2" required="required" placeholder="Ide is"/>
            <button type="button" onclick="showMessage()">Kiszámolás itt most</button>
        </form>
        <p id="result"></p>

        <script>
            function showMessage() {
                var input1 = document.getElementById('input1').value;
                var input2 = document.getElementById('input2').value;
                if(input1 && input2) {
                    document.getElementById('result').innerText = "Hello World";
                } else {
                    document.getElementById('result').innerText = "Nincsenek szamok:c";
                }
            }
        </script>
    </body>
    </html>
    `))

    tmpl.Execute(w, pageVariables)
}
