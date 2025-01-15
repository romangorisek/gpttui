package dbservice

type DbService interface {
	FetchData() (string, error)
}
