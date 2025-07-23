package discogs

import pb "github.com/brotherlogic/discogs/proto"

func ReleaseIsDigital(r *pb.Release) bool {
	// Is this a digital release
	for _, format := range r.GetFormats() {
		if format.GetName() == "CD" || format.GetName() == "CDr" || format.GetName() == "File" {
			return true
		}

		for _, desc := range format.GetDescriptions() {
			if desc == "CD" || desc == "CDr" || desc == "File" {
				return true
			}
		}
	}
	return false
}
