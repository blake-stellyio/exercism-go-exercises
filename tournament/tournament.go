package tournament

import "io"

func Tally(reader io.Reader, writer io.Writer) error {
	io.WriteString(writer, "Team                           | MP |  W |  D |  L |  P")
}
