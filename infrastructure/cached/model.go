package cached

type StatusChat string

type GenerateText struct {
	Status string
	Text   string
	Error  error
}
