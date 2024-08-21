package dashboard

import (
  DB "SCTI/database"
  "fmt"
  "strings"
  "strconv"
  "net/http"
)

func PostCadastros(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("accessToken")
  if err != nil {
    // fmt.Println("Error Getting cookie:", err)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  if cookie.Value == "-1" {
    // fmt.Println("Invalid accessToken")
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }
  activityID, err := strconv.Atoi(r.FormValue("id"))
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid activity ID: %v", err), http.StatusBadRequest)
    return
  }

  fmt.Println(activityID)
  w.WriteHeader(http.StatusOK)
}

func ActivitiesList() (string, error) {
    activities, err := DB.GetAllActivities()
    if err != nil {
        return "", fmt.Errorf("could not get activities: %v", err)
    }
    var html strings.Builder
    html.WriteString(`
        <div class="atividades">
    `)
    for _, a := range activities {
        html.WriteString(fmt.Sprintf(`
            <div class="atividade" id="activity-%v">
                <div class="info">
                  <div class="id">ID: %v</div>
                  <div class="tipo">%v</div>
                </div>
                <div class="content">
                  <div class="upper_content">
                    <div class="a_title">
                      <div class="topico">%v</div>
                      <div class="palestrante">%v</div>
                    </div>
                    <div class="sub_header">
                      <div class="dia">%v</div>
                      <div class="hora">%v</div>
                      <div class="sala">%v</div>
                      <div class="vagas">%v</div>
                    </div>
                  </div>
                  <div class="descricao">%v</div>
                </div>
                <button class="cadastrar" hx-post="/cadastrar" hx-trigger="click" hx-target="#activity-%v" hx-swap="delete" hx-vals='{"id": "%v"}'>
                    Cadastrar
                </button>
            </div>
        `, a.Activity_id, a.Activity_id, a.Activity_type, a.Topic, a.Speaker, a.Day, a.Time, a.Room, a.Spots, a.Description, a.Activity_id, a.Activity_id))
    }
    html.WriteString(`
        </div>
    `)
    return html.String(), nil
}

