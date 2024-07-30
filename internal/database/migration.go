package database

import (
	"bytes"
	atlas "github.com/Totus-Floreo/Atlas-SDK-Go"
	"github.com/open-streaming-solutions/user-service/internal/logging"
	"log/slog"
	"net/url"
)

func DoMigration(logger logging.Logger, dbURL, devDbURL *url.URL, DesiredURLs []*url.URL) error {
	log := logger.With("step", "migrations")
	log.Info("Starting database migration")

	var buf bytes.Buffer
	client := atlas.NewClient(&buf)
	opts := atlas.SchemaApplyOptions{
		CurrentURL:  dbURL,
		DesiredURLs: DesiredURLs,
		DevURL:      devDbURL,
		Approval:    true,
	}

	AttrURLs := make([]slog.Attr, 0, len(DesiredURLs))
	for _, desiredURL := range DesiredURLs {
		AttrURLs = append(AttrURLs, slog.String("url", desiredURL.String()))
	}
	log.Info("Applying schema", AttrURLs)

	if err := client.SchemaApply(opts); err != nil {
		slog.Error("Unable to apply schema", slog.String("error", err.Error()))
		return err
	}

	log.Info("Database migration completed successfully", slog.String("schema", buf.String()))

	return nil
}
