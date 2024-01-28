package models

type object struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

type entity struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

type domain struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

type tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type role struct {
	ID     string   `json:"id"`
	Grants []string `json:"grants"`
}

type group struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

type assignment struct {
	Resource string `json:"resource"`
	Role     string `json:"role"`
}

type user struct {
	ID          string       `json:"id"`
	Assignments []assignment `json:"assignments"`
}

type BundleData struct {
	Objects  []object        `json:"objects"`
	Entities []entity        `json:"entities"`
	Domains  []domain        `json:"domains"`
	Tenants  []tenant        `json:"tenants"`
	Roles    map[string]role `json:"roles"`
	Groups   []group         `json:"groups"`
	Users    []user          `json:"users"`
}
