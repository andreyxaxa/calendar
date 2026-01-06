package date

import (
	"fmt"
	"strings"
	"time"
)

// Date -.
type Date struct {
	time.Time
}

// UnmarshalJSON -.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected: YYYY-MM-DD")
	}

	d.Time = t

	return nil
}

// MarshalJSON -.
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}
