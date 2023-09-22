package structtags

type clientConn struct{}

func newClientConn(_ ClientConfig) *clientConn {
	return &clientConn{}
}

type Client struct {
	conn *clientConn
}

type options struct {
	clientConfig ClientConfig
}

type Option func(*options)

func Config(clientConfig ClientConfig) Option {
	return func(o *options) {
		o.clientConfig = clientConfig
	}
}

func ConfigFromEnv() Option {
	return func(o *options) {
		o.clientConfig = FromEnv[ClientConfig]()
	}
}

func DefaultConfig() ClientConfig {
	return FromDefaultTags[ClientConfig]()
}

func NewClient(opts ...Option) (*Client, error) {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}

	conn := newClientConn(opt.clientConfig)

	return &Client{
		conn: conn,
	}, nil
}
