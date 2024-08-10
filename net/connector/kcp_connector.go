package superConnector

import (
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
)

type (
	KCPConnector struct {
		cfacade.Component
		Connector
		Options
	}
)

func (*KCPConnector) Name() string {
	return "kcp_connector"
}

func (t *KCPConnector) OnAfterInit() {
}

func (t *KCPConnector) OnStop() {
	t.Stop()
}

func NewKCP(address string, opts ...Option) *KCPConnector {
	if address == "" {
		clog.Warn("Create kcp connector fail. Address is null.")
		return nil
	}

	tcp := &KCPConnector{
		Options: Options{
			address:  address,
			certFile: "",
			keyFile:  "",
			chanSize: 256,
		},
	}

	for _, opt := range opts {
		opt(&tcp.Options)
	}

	tcp.Connector = NewConnector(tcp.chanSize)

	return tcp
}

func (t *KCPConnector) Start() {
	listener, err := t.GetListener(t.certFile, t.keyFile, "kcp"+t.address)
	if err != nil {
		clog.Fatalf("failed to listen: %s", err)
	}

	clog.Infof("Kcp connector listening at Address %s", t.address)
	if t.certFile != "" || t.keyFile != "" {
		clog.Infof("certFile = %s, keyFile = %s", t.certFile, t.keyFile)
	}

	t.Connector.Start()

	for t.Running() {
		conn, err := listener.Accept()
		if err != nil {
			clog.Errorf("Failed to accept KCP connection: %s", err.Error())
			continue
		}

		t.InChan(conn)
	}
}

func (t *KCPConnector) Stop() {
	t.Connector.Stop()
}
