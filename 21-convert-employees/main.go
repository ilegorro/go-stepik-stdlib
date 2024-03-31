package main

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	XMLName xml.Name `xml:"employee"`
	Id      int      `xml:"id,attr,omitempty"`
	Name    string   `xml:"name,omitempty"`
	City    string   `xml:"city,omitempty"`
	Salary  int      `xml:"salary,omitempty"`
}

type Department struct {
	XMLName   xml.Name   `xml:"department"`
	Code      string     `xml:"code"`
	Employees []Employee `xml:"employees>employee,omitempty"`
}

type Organization struct {
	XMLName     xml.Name     `xml:"organization"`
	Departments []Department `xml:"department,omitempty"`
}

func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	var org Organization
	dec := xml.NewDecoder(inXML)
	err := dec.Decode(&org)
	if err != nil {
		return err
	}

	w := csv.NewWriter(outCSV)
	err = w.Write([]string{"id", "name", "city", "department", "salary"})
	if err != nil {
		return err
	}
	for _, dep := range org.Departments {
		for _, emp := range dep.Employees {
			err = w.Write([]string{strconv.Itoa(emp.Id), emp.Name, emp.City, dep.Code, strconv.Itoa(emp.Salary)})
			if err != nil {
				return err
			}
		}
	}
	w.Flush()
	if err = w.Error(); err != nil {
		return err
	}

	return nil
}

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
            <employee id="22">
                <name>Фаня</name>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	err := ConvertEmployees(out, in)
	if err != nil {
		panic(err)
	}
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
