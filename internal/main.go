package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const mimeDBVersion = "1.50.0"

var mimeTypesExtensionsTpl = `package mime

var mimeTypesExtensions = map[string][]string{
{{range $k, $v := .MimeTypesExtensions}}"{{$k}}": { {{range $i := $v}}"{{$i}}",{{end}} },
{{end}}
}

var extensionMimeTypes = map[string][]string{
{{range $k, $v := .ExtensionMimeTypes}}"{{$k}}": { {{range $i := $v}}"{{$i}}",{{end}} },
{{end}}
}
`

func main() {
	var (
		err             error
		mimeDBData      map[string][]string
		freeDesktopData map[string][]string
		data            struct {
			MimeTypesExtensions map[string][]string
			ExtensionMimeTypes  map[string][]string
		}
	)

	if freeDesktopData, err = fetchFreeDesktopData(); err != nil {
		panic(err)
	}

	if mimeDBData, err = fetchMimiDBData(); err != nil {
		panic(err)
	}

	data.MimeTypesExtensions = map[string][]string{}
	data.ExtensionMimeTypes = map[string][]string{}

	for k, v := range freeDesktopData {
		data.MimeTypesExtensions[k] = append(data.MimeTypesExtensions[k], v...)
	}

	for k, v := range mimeDBData {
		data.MimeTypesExtensions[k] = unique(append(data.MimeTypesExtensions[k], v...))
	}

	for k, v := range data.MimeTypesExtensions {
		for _, j := range v {
			data.ExtensionMimeTypes[j] = unique(append(data.ExtensionMimeTypes[j], k))
		}
	}

	t := template.Must(template.New("tpl").Parse(mimeTypesExtensionsTpl))
	t.Execute(os.Stdout, data)
}

type mimiDbData struct {
	Source       string   `json:"source"`
	Compressible bool     `json:"compressible,omitempty"`
	Extensions   []string `json:"extensions,omitempty"`
}

func fetchMimiDBData() (map[string][]string, error) {
	var (
		err  error
		res  *http.Response
		buf  []byte
		body map[string]mimiDbData
		data map[string][]string
	)

	if res, err = http.Get(fmt.Sprintf("https://cdn.jsdelivr.net/gh/jshttp/mime-db@v%s/db.json", mimeDBVersion)); err != nil {
		return nil, err
	}

	if buf, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf, &body); err != nil {
		return nil, err
	}

	data = make(map[string][]string)

	for k, v := range body {
		if len(v.Extensions) > 0 {
			data[k] = v.Extensions
		}
	}

	return data, nil
}

type freeDesktopData struct {
	XMLName  xml.Name `xml:"mime-info"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	MimeType []struct {
		Text        string `xml:",chardata"`
		Type        string `xml:"type,attr"`
		Comment     string `xml:"comment"`
		GenericIcon struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"generic-icon"`
		Glob []struct {
			Text    string `xml:",chardata"`
			Pattern string `xml:"pattern,attr"`
		} `xml:"glob"`
		Magic struct {
			Text     string `xml:",chardata"`
			Priority string `xml:"priority,attr"`
			Match    []struct {
				Text   string `xml:",chardata"`
				Type   string `xml:"type,attr"`
				Value  string `xml:"value,attr"`
				Offset string `xml:"offset,attr"`
				Match  struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Value  string `xml:"value,attr"`
					Offset string `xml:"offset,attr"`
					Match  []struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Value  string `xml:"value,attr"`
						Offset string `xml:"offset,attr"`
					} `xml:"match"`
				} `xml:"match"`
			} `xml:"match"`
		} `xml:"magic"`
		Acronym         string `xml:"acronym"`
		ExpandedAcronym string `xml:"expanded-acronym"`
		SubClassOf      struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"sub-class-of"`
		Alias []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"alias"`
		RootXML struct {
			Text         string `xml:",chardata"`
			NamespaceURI string `xml:"namespaceURI,attr"`
			LocalName    string `xml:"localName,attr"`
		} `xml:"root-XML"`
	} `xml:"mime-type"`
}

func fetchFreeDesktopData() (map[string][]string, error) {
	var (
		err  error
		res  *http.Response
		buf  []byte
		body freeDesktopData
		data map[string][]string
	)

	if res, err = http.Get("https://gitlab.freedesktop.org/xdg/shared-mime-info/-/raw/master/data/freedesktop.org.xml.in"); err != nil {
		return nil, err
	}

	if buf, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(buf, &body); err != nil {
		return nil, err
	}

	data = make(map[string][]string)

	for _, i := range body.MimeType {
		if len(i.Glob) == 0 {
			continue
		}

		exts := make([]string, 0)

		for _, v := range i.Glob {
			if strings.HasPrefix(v.Pattern, "*.") {
				if !strings.HasPrefix(v.Pattern[2:], "[") && !strings.HasSuffix(v.Pattern[2:], "]") {
					exts = append(exts, v.Pattern[2:])
				}
			}
		}

		if len(exts) == 0 {
			continue
		}

		data[i.Type] = append(data[i.Type], exts...)

		for _, j := range i.Alias {
			data[j.Type] = append(data[j.Type], exts...)
		}
	}

	return data, nil
}

func unique(s []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
