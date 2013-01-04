package GitHubParser

type GHError struct {
	message string
}

func (g *GHError) Error() string {
	return g.message
}
