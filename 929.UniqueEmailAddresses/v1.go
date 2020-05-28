import "strings"

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, v := range emails {
		ai := strings.Index(v, "@")
		local, domain := v[:ai], v[ai:]
		if i := strings.Index(local, "+"); i > 0 {
			local = local[:i]
		}
		local = strings.Replace(local, ".", "", -1)
		ue := local + "@" + domain
		if !m[ue] {
			m[ue] = true
		}
	}
	return len(m)
}
