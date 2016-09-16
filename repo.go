package main

import "fmt"

var currentId int

var providers Providers

// Give us some seed data
func init() {
	RepoCreateProvider(Provider{Name: "Write presentation"})
	RepoCreateProvider(Provider{Name: "Host meetup"})
}

func RepoFindProvider(id int) Provider {
	for _, t := range providers {
		if t.Id == id {
			return t
		}
	}
	// return empty Provider if not found
	return Provider{}
}

func RepoCreateProvider(p Provider) Provider {
	currentId += 1
	p.Id = currentId
	providers = append(providers, p)
	return p
}

func RepoDestroyProvider(id int) error {
	for i, p := range providers {
		if p.Id == id {
			providers = append(providers[:i], providers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Provider with id of %d to delete", id)
}
