package strand

import "strings"

func ToRNA(dna string) string {

	var output strings.Builder

	for _, x := range dna {

		switch {
		case x == 'G':
			output.WriteRune('C')
		case x == 'C':
			output.WriteRune('G')
		case x == 'T':
			output.WriteRune('A')
		case x == 'A':
			output.WriteRune('U')

		}

	}

	return output.String()

}
