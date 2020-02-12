package cmd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/olekukonko/tablewriter"
	"github.com/sfuruya0612/aie-emu/internal"
	"github.com/urfave/cli"
)

type replace struct {
	Name          string
	ManagedPolicy string
	InlinePolicy  string
	Group         string
	AccessKey     string
	PWLastUsed    string
	CreateDate    string
}

var header = []string{
	"Name",
	"ManagedPolicy",
	"InlinePolicy",
	"Group",
	"AccessKey",
	"PWLastUsed",
	"CreateDate",
}

var mdExHeader string = `<table>
  <tr>
    <th>Name</th>
    <th>ManagedPolicy</th>
    <th>InlinePolicy</th>
    <th>Group</th>
    <th>AccessKey</th>
    <th>PWLastUsed</th>
    <th>CreateDate</th>
  </tr>`

func GetIamList(c *cli.Context) error {
	profile := c.GlobalString("profile")
	output := c.GlobalString("output")

	client := internal.NewIamSess(profile, "ap-northeast-1")
	list, err := client.ListUsers(&iam.ListUsersInput{})
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	switch output {
	case "csv":
		if err := outputCsv(list); err != nil {
			return fmt.Errorf("%v", err)
		}
	case "md":
		if err := outputMarkdown(list); err != nil {
			return fmt.Errorf("%v", err)
		}
	case "ex":
		if err := outputMarkdownExtra(list); err != nil {
			return fmt.Errorf("%v", err)
		}
	case "stdout":
		if err := outputStdout(list); err != nil {
			return fmt.Errorf("%v", err)
		}
	default:
		return fmt.Errorf("No match output option pattern")
	}

	return nil
}

func outputCsv(list internal.Users) error {
	s := []string{"iamuser-list", time.Now().Format("2006-01-02") + ".csv"}
	filename := strings.Join(s, "-")

	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Get current dir: %v", err)
	}

	filepath := path.Join(pwd, filename)

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Create file: %v", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	if err := w.Write(header); err != nil {
		return fmt.Errorf("%v", err)
	}

	for _, i := range list {
		err := w.Write([]string{
			i.Name,
			i.ManagedPolicy,
			i.InlinePolicy,
			i.Group,
			i.AccessKey,
			i.PWLastUsed,
			i.CreateDate,
		})
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	}
	w.Flush()

	fmt.Printf("Output csv file: %v", filepath)
	return nil
}

func outputMarkdown(list internal.Users) error {
	w := tablewriter.NewWriter(os.Stdout)

	w.SetHeader(header)
	w.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	w.SetCenterSeparator("|")

	for _, i := range list {
		w.Append([]string{
			i.Name,
			i.ManagedPolicy,
			i.InlinePolicy,
			i.Group,
			i.AccessKey,
			i.PWLastUsed,
			i.CreateDate,
		})
	}
	w.Render()

	return nil
}

func outputMarkdownExtra(list internal.Users) error {
	var body string
	for _, i := range list {
		b, err := mdExBody(i)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		body = body + b
	}

	fmt.Printf("%v\n%v%v", mdExHeader, body, "</table>")
	return nil
}

func mdExBody(i internal.User) (string, error) {
	body := `  <tr>
    <td>{{.Name}}</td>
    <td>{{.ManagedPolicy}}</td>
    <td>{{.InlinePolicy}}</td>
    <td>{{.Group}}</td>
    <td>{{.AccessKey}}</td>
    <td>{{.PWLastUsed}}</td>
    <td>{{.CreateDate}}</td>
  </tr>
`

	var result bytes.Buffer

	tmp, err := template.New("body").Parse(body)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	r := replace{
		Name:          i.Name,
		ManagedPolicy: i.ManagedPolicy,
		InlinePolicy:  i.InlinePolicy,
		Group:         i.Group,
		AccessKey:     i.AccessKey,
		PWLastUsed:    i.PWLastUsed,
		CreateDate:    i.CreateDate,
	}

	if err = tmp.Execute(&result, r); err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return result.String(), nil
}

func outputStdout(list internal.Users) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)

	if _, err := fmt.Fprintln(w, strings.Join(header, "\t")); err != nil {
		return fmt.Errorf("%v", err)
	}

	for _, i := range list {
		if _, err := fmt.Fprintln(w, tabString(i)); err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func tabString(i internal.User) string {
	fields := []string{
		i.Name,
		i.ManagedPolicy,
		i.InlinePolicy,
		i.Group,
		i.AccessKey,
		i.PWLastUsed,
		i.CreateDate,
	}

	return strings.Join(fields, "\t")
}
