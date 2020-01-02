package types

import "encoding/xml"

type Book struct {
	XMLName   xml.Name `xml:"TEI"`
	Chardata  string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	TeiHeader struct {
		Text     string `xml:",chardata"`
		Lang     string `xml:"lang,attr"`
		FileDesc struct {
			Text      string `xml:",chardata"`
			TitleStmt struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"title"`
				Editor    string   `xml:"editor"`
				Author    string   `xml:"author"`
				Sponsor   string   `xml:"sponsor"`
				Funder    string   `xml:"funder"`
				Principal []string `xml:"principal"`
				Meeting   string   `xml:"meeting"`
				RespStmt  []struct {
					Text     string `xml:",chardata"`
					PersName struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"persName"`
					Resp string `xml:"resp"`
				} `xml:"respStmt"`
			} `xml:"titleStmt"`
			PublicationStmt struct {
				Text      string `xml:",chardata"`
				Authority string `xml:"authority"`
				Idno      struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"idno"`
				Availability struct {
					Text    string `xml:",chardata"`
					Licence struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
					} `xml:"licence"`
				} `xml:"availability"`
				Date      string `xml:"date"`
				Publisher string `xml:"publisher"`
				PubPlace  string `xml:"pubPlace"`
			} `xml:"publicationStmt"`
			SourceDesc struct {
				Text     string `xml:",chardata"`
				ListBibl struct {
					Text       string `xml:",chardata"`
					Lang       string `xml:"lang,attr"`
					BiblStruct struct {
						Text   string `xml:",chardata"`
						Monogr struct {
							Text   string `xml:",chardata"`
							Title  string `xml:"title"`
							Editor struct {
								Text     string `xml:",chardata"`
								PersName struct {
									Text string `xml:",chardata"`
									Name struct {
										Text string `xml:",chardata"`
										Lang string `xml:"lang,attr"`
									} `xml:"name"`
								} `xml:"persName"`
							} `xml:"editor"`
							Author struct {
								Text string `xml:",chardata"`
								Ref  string `xml:"ref,attr"`
							} `xml:"author"`
							Imprint struct {
								Text      string `xml:",chardata"`
								Publisher string `xml:"publisher"`
								PubPlace  string `xml:"pubPlace"`
								Date      string `xml:"date"`
							} `xml:"imprint"`
							BiblScope struct {
								Text string `xml:",chardata"`
								Unit string `xml:"unit,attr"`
							} `xml:"biblScope"`
						} `xml:"monogr"`
						Ref struct {
							Text   string `xml:",chardata"`
							Target string `xml:"target,attr"`
						} `xml:"ref"`
					} `xml:"biblStruct"`
				} `xml:"listBibl"`
			} `xml:"sourceDesc"`
		} `xml:"fileDesc"`
		EncodingDesc struct {
			Text     string `xml:",chardata"`
			P        string `xml:"p"`
			RefsDecl struct {
				Text        string `xml:",chardata"`
				N           string `xml:"n,attr"`
				CRefPattern struct {
					Text               string `xml:",chardata"`
					N                  string `xml:"n,attr"`
					MatchPattern       string `xml:"matchPattern,attr"`
					ReplacementPattern string `xml:"replacementPattern,attr"`
				} `xml:"cRefPattern"`
			} `xml:"refsDecl"`
		} `xml:"encodingDesc"`
		ProfileDesc struct {
			Text      string `xml:",chardata"`
			LangUsage struct {
				Text     string `xml:",chardata"`
				Language []struct {
					Text  string `xml:",chardata"`
					Ident string `xml:"ident,attr"`
				} `xml:"language"`
			} `xml:"langUsage"`
		} `xml:"profileDesc"`
		RevisionDesc struct {
			Text   string `xml:",chardata"`
			Change []struct {
				Text string `xml:",chardata"`
				Who  string `xml:"who,attr"`
				When string `xml:"when,attr"`
			} `xml:"change"`
		} `xml:"revisionDesc"`
	} `xml:"teiHeader"`
	Text struct {
		Text string `xml:",chardata"`
		Body struct {
			Text string `xml:",chardata"`
			Div  struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				N    string `xml:"n,attr"`
				Div  []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Subtype string `xml:"subtype,attr"`
					N       string `xml:"n,attr"`
					Pb      struct {
						Text string `xml:",chardata"`
						N    string `xml:"n,attr"`
					} `xml:"pb"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Pb   []struct {
							Text string `xml:",chardata"`
							N    string `xml:"n,attr"`
						} `xml:"pb"`
					} `xml:"p"`
				} `xml:"div"`
			} `xml:"div"`
		} `xml:"body"`
	} `xml:"text"`
} 
