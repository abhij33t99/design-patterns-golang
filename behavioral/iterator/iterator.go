package behavioral

type Iterator interface {
	hasNext() bool
	next() *User
}

type UserIterator struct {
	users []*User
	index int
}

func (u *UserIterator) hasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) next() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
