package notfound

import (
  "SCTI/fileserver"
  "net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/notfound.gohtml")
  t.Execute(w, nil)
}
