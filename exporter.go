package election

type Exporter interface {
	Export(config Config) string
}
