package go_repo_puppy

import go_repo_dog "github.com/And-code/go-repo-dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return go_repo_dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return go_repo_dog.WhenGrownUp(Barks())
}
