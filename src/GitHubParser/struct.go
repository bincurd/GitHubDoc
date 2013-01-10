package GitHubParser

type Parser struct {
	TagHead string
	TagTail string
}

type GHError struct {
	message string
}

func (g *GHError) Error() string {
	return g.message
}
