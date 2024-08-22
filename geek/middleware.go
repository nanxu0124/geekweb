package geek

type Middleware func(next HandleFunc) HandleFunc
