package postgresql

//type CLient interface {
//	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
//	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
//	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
//
//	Begin(ctx context.Context) (pgx.Tx, error)
//}

// Нужно для того, чтобы у нас не было проблем с контейнеризацией postgresql иначе могут быть непредвиденные ошибки
//func NewCLient(ctx context.Context, maxAttempts int, username, password, host, port, database string) {
//	dsn := fmt.Sprintf("postgresql: %s:%s", username, password)
//	utils.DoWithTries(func() error {
//		ctx, cancel := context.WithTimeout(ctx, dsn)
//		defer cancel()
//
//		pool, err := pgxpool.Connect(ctx, dsn)
//		if err != nil {
//			return err
//		}
//
//		return nil
//
//	}, maxAttempts, 5*time.Second)
//}
