package behavioral

type Collection interface {
	Iterator() Iterator
}

type User struct {
	name string
	age  int
}

type UserCollection struct {
	users []*User
}

func (uc *UserCollection) Iterator() Iterator {
	return &UserIterator{
		users: uc.users,
	}
}
