module git.darknebu.la/chaosdorf/freitagsfoo

go 1.13

require (
	git.darknebu.la/chaosdorf/freitagsfoo/src/db v0.0.0-00010101000000-000000000000
<<<<<<< HEAD
	git.darknebu.la/chaosdorf/freitagsfoo/src/structs v0.0.0-20200718235609-fc63adf55849
=======
	git.darknebu.la/chaosdorf/freitagsfoo/src/structs v0.0.0-00010101000000-000000000000
>>>>>>> refs/remotes/origin/master
	github.com/go-pg/pg/v9 v9.1.6
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.4
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
)

replace git.darknebu.la/chaosdorf/freitagsfoo/src/structs => ./src/structs

replace git.darknebu.la/chaosdorf/freitagsfoo/src/db => ./src/db
