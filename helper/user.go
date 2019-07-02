package helper

import "os/user"

// HomeDirectory return user home dir path
func HomeDirectory() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}
