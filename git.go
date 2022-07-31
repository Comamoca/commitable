package main
import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type InputResult struct {
  Prefix string
  Summary string
  BreakingChangeMsg string
  RelatedIssue string
  Scope string
}

func getStatus() ([][]string, error) {
  status, err := exec.Command("git", "status", "-s").Output()
  // fmt.Println(status)
  var statusFiles [][]string

  if err != nil {
    // fmt.Println(err)
    return nil, err
  }
  statusTxt := string(status)

  splitedLines := strings.Split(statusTxt, "\n")

  for _, line := range splitedLines {
    if len(line) == 0 {
      continue
    }

    splited := strings.Split(line, " ")
    // fmt.Println("splited")
    // fmt.Println(splited)
    statusFile := [] string{splited[0], splited[1]}
    statusFiles = append(statusFiles, statusFile)
  }

  // fmt.Println(statusFiles)
  return statusFiles, nil
}

func getFileString(statusFiles [][]string) string {
  var flist []string

  for _, statusFile := range statusFiles {
    // fmt.Println(statusFiles[0])
    // fmt.Println(statusFiles[1])
    flist = append(flist, statusFile[1])
  }

  // fmt.Println(flist)
  return strings.Join(flist, ", ")
}

func Commit(msg string) error {
  file, err := ioutil.TempFile("", "COMMIT_MSG")
  if err != nil {
    // fmt.Println(err)
	  return err
  }

  // fmt.Println(file.Name())
  defer os.Remove(file.Name())

  _, err = file.Write([]byte(msg))
  if err != nil{
    // fmt.Println(err)
	  return err
  }

  cmd := exec.Command("git", "commit", "-F")
  cmd.Args = append(cmd.Args, file.Name())

  err = cmd.Run()

  if err != nil {
    return err
  }

  return nil
}

func convertStruct(in PromptResult, files string) InputResult {
  // fmt.Println(in)
  prefix := in.Prefix 
  summary := in.Summary
  breakingChange := in.BreakingChangeMsg
  issue := in.RelatedIssue

  out := InputResult{
    Prefix: prefix,
    Summary: summary,
    BreakingChangeMsg: breakingChange,
    RelatedIssue: issue,
    Scope: files,
  }

  return out
}

func genCommitMsg(input PromptResult) (string, error) {
  writer := new(strings.Builder)
  // fmt.Println("input")
  // fmt.Println(input)

  tmpl := `
  {{.Prefix}}

  {{.Summary}}

  {{.BreakingChangeMsg}}

  {{.RelatedIssue}}`

  t, err := template.New("commit").Parse(tmpl)
  if err != nil {
    return "", err
  }

  if err = t.Execute(writer, input); err != nil {
    return "", err
  }

  return writer.String(), nil

}
