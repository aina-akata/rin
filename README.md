# 🚀 RIN – REST Interface Nexus

**RIN** est une bibliothèque Go modulaire conçue pour simplifier et centraliser l’interaction avec des APIs REST dans des contextes DevOps et automatisation.  
Elle permet de décrire des endpoints dynamiques, d'ajouter facilement de l'authentification, et de router intelligemment les réponses selon leur code HTTP.

---

## 📦 Fonctionnalités principales

- ✅ Client HTTP réutilisable avec support de l’authentification (Token, Basic)
- ✅ Système de **ressources REST dynamiques** avec templating (`/users/{{.id}}`)
- ✅ Gestion simplifiée des réponses HTTP via un **router par code HTTP**
- ✅ Séparation claire des composants : `Client`, `API`, `Router`, `Ressource`, `Auth`

---

## 🧩 Structure des composants

| Composant | Description |
|----------|-------------|
| `Client` | Gère les requêtes HTTP, l’authentification, et l’envoi |
| `API` | Wrapper logique autour du client + enregistrement des ressources |
| `RestRessources` | Endpoint avec méthode HTTP et templating des paramètres |
| `CBRouter` | Router basé sur le code de réponse HTTP |
| `Authentication` | Interface d’auth (Token, Basic…) facilement extensible |

---

## ⚙️ Exemple d’utilisation

```go
package main

import (
    "fmt"
    "github.com/Rakotoarilala51/rin"
    "net/http"
)

func main() {
    api := rin.NewApi("https://api.github.com")
    api.SetAuth(rin.NewAuthToken("your_token_here"))

    router := rin.NewRouter()
    router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
        fmt.Println("Success:", resp.Status)
        return nil
    })

    res := rin.NewRessource("repos/{{.owner}}/{{.repo}}", http.MethodGet, router)
    api.AddRessource("GetRepo", res)

    err := api.Call("GetRepo", map[string]string{
        "owner": "RaMaitre",
        "repo": "rin",
    })
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```
## Installation

Utilisez `go get` pour ajouter la bibliothèque à votre projet :

```bash
go get github.com/Rakotoarilala51/rin

