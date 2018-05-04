package corelib

type Doc2018 struct {
	Doi         string   `json:"doi"`
	CoreID      string   `json:"coreId"`
	Oai         string   `json:"oai"`
	Identifiers []string `json:"identifiers"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Enrichments struct {
		References []struct {
			Date    string   `json:"date"`
			Raw     string   `json:"raw"`
			Title   string   `json:"title"`
			Authors []string `json:"authors"`
		} `json:"references"`
		DocumentType struct {
			Type       string  `json:"type"`
			Confidence float32 `json:"confidence"`
		} `json:"documentType"`
	} `json:"enrichments"`
	Contributors       []string `json:"contributors"`
	DatePublished      string   `json:"datePublished"`
	Abstract           string   `json:"abstract"`
	DownloadURL        string   `json:"downloadUrl"`
	FullTextIdentifier string   `json:"fullTextIdentifier"`
	PdfHashValue       string   `json:"pdfHashValue"`
	Publisher          string   `json:"publisher"`
	RawRecordXML       string   `json:"rawRecordXml"`
	//Journals           []string `json:"journals"`
	Journals []struct {
		Identifiers []string `json:"identifiers"`
		Title       string   `json:"title"`
	} `json:"journals"`
	Language struct {
		Code string `json:"code"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"language"`
	Relations []string `json:"relations"`
	Year      int      `json:"year"`
	Topics    []string `json:"topics"`
	Subjects  []string `json:"subjects"`
	FullText  string   `json:"fullText"`
}
