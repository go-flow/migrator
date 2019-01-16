package migrator

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-flow/migrator/db"
)

// FileMigrator is Migrator implementation for SQL
// files stored on file system
type FileMigrator struct {
	Migrator
	Path string
}

// NewFileMigrator - creates Migrations for files
func NewFileMigrator(path string, dialect string, db db.Store) FileMigrator {
	fm := FileMigrator{
		Migrator: newMigrator(dialect, db),
		Path:     path,
	}

	err := fm.loadMigrations()
	if err != nil {
		panic(err)
	}
	return fm
}

func (fm *FileMigrator) loadMigrations() error {
	dir := fm.Path
	fi, err := os.Stat(dir)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return errors.New("Provided path is not directory")
	}

	return filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			matches := migrationRegEx.FindAllStringSubmatch(info.Name(), -1)
			if matches == nil || len(matches) == 0 {
				return nil
			}
			f, err := os.Open(p)
			if err != nil {
				return err
			}

			content, err := fm.loadMigrationContent(f)
			if err != nil {
				return err
			}

			match := matches[0]
			dir := match[4]

			migration := Migration{
				Version:   match[1],
				Name:      match[2],
				Content:   content,
				Direction: dir,
			}

			fm.Migrations[dir] = append(fm.Migrations[dir], migration)
		}
		return nil
	})
}

func (fm *FileMigrator) loadMigrationContent(r io.Reader) (string, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	content := string(raw)
	temp := template.Must(template.New("sql").Parse(content))
	var buff bytes.Buffer
	err = temp.Execute(&buff, nil)
	if err != nil {
		return "", err
	}

	content = buff.String()
	return content, nil
}
