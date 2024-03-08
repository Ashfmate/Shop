package main

import (
	"fmt"
  // "os"
  "net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
  /* contents, err := os.ReadFile("View/index.html")
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(-1)
  } */
  cart := Cart{}
  cart.products = []Product{
    {10, "Phone", randTime(), randTime()},
    {5, "Video Game", randTime(), randTime()},
    {15, "Laptop", randTime(), randTime()},
  }
  contents := "<body>\n<ul>\n"
  for _, val := range cart.products {
    contents += fmt.Sprintln("<ul>", "<b>", val.name, "</b></ul>" , "<li>", val.price, "</li>\n<li>", val.rel_date, "</li>\n<li>", val.exp_date, "</li>")
  }
  contents += "</ul>\n</body>\n"
  fmt.Fprintln(w, string(contents));
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}
