package tenant

import (
	"fmt"
	"strconv"
	"strings"
)

// Tenant contains settings for request processing
type Tenant struct {
	AccountID uint32
	ProjectID uint32
}

// New returns new Tenant for the given authToken.
func New(token string) (*Tenant, error) {
	var t Tenant
	if err := t.Init(token); err != nil {
		return nil, err
	}
	return &t, nil
}

// String returns string representation of t.
func (t *Tenant) String() string {
	if t.ProjectID == 0 {
		return fmt.Sprintf("%d", t.AccountID)
	}
	return fmt.Sprintf("%d:%d", t.AccountID, t.ProjectID)
}

// Init initializes tenant from token.
func (t *Tenant) Init(token string) error {
	accountID, projectID, err := ParseTenant(token)
	if err != nil {
		return fmt.Errorf("cannot parse authToken %q: %w", token, err)
	}

	t.Set(accountID, projectID)
	return nil
}

// ParseTenant parses token and returns accountID and projectID from it.
func ParseTenant(tenant string) (uint32, uint32, error) {
	tmp := strings.Split(tenant, ":")
	if len(tmp) > 2 {
		return 0, 0, fmt.Errorf("unexpected number of items in tenant %q; got %d; want 1 or 2", tenant, len(tmp))
	}
	n, err := strconv.ParseUint(tmp[0], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse accountID from %q: %w", tmp[0], err)
	}
	accountID := uint32(n)
	projectID := uint32(0)
	if len(tmp) > 1 {
		n, err := strconv.ParseUint(tmp[1], 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("cannot parse projectID from %q: %w", tmp[1], err)
		}
		projectID = uint32(n)
	}
	return accountID, projectID, nil
}

// Set sets accountID and projectID for the tenant.
func (t *Tenant) Set(accountID, projectID uint32) {
	t.AccountID = accountID
	t.ProjectID = projectID
}
