package persistant

import "github.com/web"

func Get(guid string) (web.User, error) {

}


func Create(usr web.RequestUser) (web.User, error) {
	return web.User{},nil
}

func Update(guid string,usr web.RequestUser) (web.User, error) {
	return web.User{},nil
}

func Delete(guid string) error {
	return nil
}