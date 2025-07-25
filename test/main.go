package main

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "strings"
    "sync"
)

// Relation represents the mapping of locations to dates for an artist
// as returned by the relations API.
type Relation struct {
    ID             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

// DateEntry represents the dates array for an artist.
type DateEntry struct {
    ID    int      `json:"id"`
    Dates []string `json:"dates"`
}

// LocationEntry represents the locations array for an artist.
type LocationEntry struct {
    ID        int      `json:"id"`
    Locations []string `json:"locations"`
}

// Artist is the base artist data from the artists API.
type Artist struct {
    ID            int      `json:"id"`
    Image         string   `json:"image"`
    Name          string   `json:"name"`
    Members       []string `json:"members"`
    CreationDate  int      `json:"creationDate"`
    FirstAlbum    string   `json:"firstAlbum"`
    LocationsURL  string   `json:"locations"`
    DatesURL      string   `json:"concertDates"`
    RelationsURL  string   `json:"relations"`
}

// ArtistFull combines the Artist info with fetched locations, dates, and relation map.
type ArtistFull struct {
    Artist
    Locations   []string
    Dates       []string
    RelationMap map[string][]string
}

// fetchJSON fetches JSON from a URL and decodes into target.
func fetchJSON(url string, target interface{}) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    log.Println("Server running on :8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}

// homeHandler fetches all artist-related data and renders the home template.
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" || r.Method != http.MethodGet {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    // Step 1: Fetch list of artists
    var artists []Artist
    if err := fetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
        http.Error(w, "Failed to load artists", http.StatusInternalServerError)
        return
    }

    // Step 2: Concurrently fetch locations, dates, and relations for each artist
    fullList := make([]ArtistFull, len(artists))
    var wg sync.WaitGroup

    for i, art := range artists {
        wg.Add(1)
        go func(i int, art Artist) {
            defer wg.Done()
            af := ArtistFull{Artist: art}

            // Fetch locations
            var locEntry LocationEntry
            if err := fetchJSON(art.LocationsURL, &locEntry); err == nil {
                af.Locations = locEntry.Locations
            }

            // Fetch dates
            var dateEntry DateEntry
            if err := fetchJSON(art.DatesURL, &dateEntry); err == nil {
                af.Dates = dateEntry.Dates
            }

            // Fetch relations
            var rel Relation
            if err := fetchJSON(art.RelationsURL, &rel); err == nil {
                af.RelationMap = rel.DatesLocations
            }

            fullList[i] = af
        }(i, art)
    }
    wg.Wait()

    // Step 3: Parse and execute template with join function
    tmpl := template.Must(
        template.New("home.html").Funcs(template.FuncMap{"join": strings.Join}).
            ParseFiles("templates/html/home.html"),
    )

    if err := tmpl.Execute(w, fullList); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}